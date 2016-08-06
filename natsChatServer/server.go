package natsChatServer

import (
	"encoding/json"
	"errors"
	"time"

	"fmt"
	"git.oschina.net/dddailing/schoolapp/controllers"
	"github.com/dailing/levlog"
	"github.com/garyburd/redigo/redis"
	"github.com/nats-io/nats"
	"runtime"
	"strings"
)

type Message struct {
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
	From    string    `json:"from"`
	To      string    `json:"to"`
}

type Messages struct {
	Msgs []Message `json:"messages"`
}

type NatsServer struct {
	RedisDB         *redis.Pool
	MsgSentConn     *nats.Conn
	MsgListenConn   *nats.Conn
	MsgSendReply    *nats.Conn
	StartEvent      *nats.Conn
	StartEventReply *nats.Conn
	Timeout         time.Duration
	SendRequest     chan string
}

func ShowThings(v interface{}) {
	bytes, err := json.Marshal(v)
	levlog.E(err)
	levlog.Info(string(bytes))
}

func parseToken(token string) controllers.TypeTokenInfo {
	actualToken := controllers.BaseDecode(token)
	return controllers.ParseToken(actualToken)
}

func NewServer() *NatsServer {
	server := &NatsServer{
		RedisDB: &redis.Pool{
			MaxIdle:     100,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", "localhost:6379")
				if err != nil {
					return nil, err
				}
				return c, err
			},
		},
		SendRequest: make(chan string, 1024),
		Timeout:     time.Second * 5,
	}
	var err error
	server.MsgSentConn, err = nats.Connect("nats://localhost:4222")
	levlog.E(err)
	server.MsgListenConn, err = nats.Connect("nats://localhost:4222")
	levlog.E(err)
	server.MsgSendReply, err = nats.Connect("nats://localhost:4222")
	levlog.E(err)
	server.StartEvent, err = nats.Connect("nats://localhost:4222")
	levlog.E(err)
	server.StartEventReply, err = nats.Connect("nats://localhost:4222")
	levlog.E(err)
	if err != nil {
		return nil
	}
	return server
}

func (c *NatsServer) TriggerSentEvent(destination string) {
	select {
	case c.SendRequest <- destination:
	default:
		levlog.Warning("Delayed Message deliver")
	}
}

func (c *NatsServer) Start() {
	// listen to client sent message
	c.MsgListenConn.Subscribe("Sent.*.*", func(msg *nats.Msg) {
		strs := strings.Split(msg.Subject, ".")
		if len(strs) < 3 {
			levlog.Error("Not correct subject : ", msg.Subject)
		}
		token := strs[1]
		destination := strs[2]
		userInfo := parseToken(token)
		levlog.Info("Recving Sent Event : ",
			userInfo.UserID,
			" --> ",
			destination,
			":",
			string(msg.Data),
		)
		messages := Messages{}
		err := json.Unmarshal(msg.Data, &messages)
		levlog.E(err)
		if err != nil {
			return
		}
		for i, _ := range messages.Msgs {
			messages.Msgs[i].From = fmt.Sprint(userInfo.UserID)
			messages.Msgs[i].To = destination
			messages.Msgs[i].Time = time.Now()
		}
		c.Store(messages)
		// reply to client
		c.MsgSendReply.Publish(msg.Reply, []byte("OK"))
		c.TriggerSentEvent(destination)
	})
	// start sending routing
	sendingFunc := func() {
		// TODO make this function exit diligently
		for {
			destination := <-c.SendRequest
			levlog.Info("Sending Messages to ", destination)
			msgs, err := c.FetchMessage(destination)
			levlog.E(err)
			if err != nil {
				continue
			}
			if msgs.Msgs == nil || len(msgs.Msgs) <= 0 {
				continue
			}
			err = c.SentMessages(msgs)
			levlog.E(err)
			if err != nil {
				c.Store(msgs)
				continue
			}
		}
	}
	for i := 0; i < 1; i++ {
		go sendingFunc()
	}
	// Listen to start event
	c.StartEvent.Subscribe("StartUp.*", func(msg *nats.Msg) {
		strs := strings.Split(msg.Subject, ".")
		token := strs[1]
		userInfo := parseToken(token)
		levlog.Info("Start up Of client : ", userInfo.UserID)
		c.TriggerSentEvent(fmt.Sprint(userInfo.UserID))
		c.StartEventReply.Publish(msg.Reply, []byte("OK"))
	})
}

func (c *NatsServer) GenMsgKey(from string) string {
	return "msgStorage." + from
}

func (c *NatsServer) GetMsgSentSubject(to string) string {
	return "Recv." + controllers.BaseEncode(to)
}

func (c *NatsServer) SentMessages(msgs Messages) error {
	if len(msgs.Msgs) <= 0 {
		levlog.Warning("Empty messages")
		return errors.New("Empty message")
	}
	destnation := msgs.Msgs[0].To
	// perform check
	for _, i := range msgs.Msgs {
		if i.To != destnation {
			levlog.Error("Destination Not Match")
			return errors.New("Destination Not Match")
		}
	}
	token := controllers.GenerateTokenByUserID(destnation)
	subject := c.GetMsgSentSubject(token)
	payload, err := json.Marshal(msgs)
	levlog.E(err)
	if err != nil {
		return nil
	}
	resp, err := c.MsgSentConn.Request(subject, payload, c.Timeout)
	levlog.E(err)
	if err != nil {
		levlog.Error(resp)
		return err
	}
	return nil
}

func (c *NatsServer) FetchMessage(destination string) (msg Messages, err error) {
	key := c.GenMsgKey(destination)
	msg = Messages{
		Msgs: make([]Message, 0),
	}
	conn := c.RedisDB.Get()
	if conn == nil {
		levlog.Info("Error , No connection")
		err = errors.New("No connection")
		return
	} else {
		defer conn.Close()
	}
	listLen, err := redis.Int(conn.Do("LLEN", key))
	levlog.Info("get number of strings : ", listLen, " from ", key)
	levlog.E(err)
	if err != nil {
		return
	}
	rev, err := redis.Strings(conn.Do("LRANGE", key, 0, listLen))
	levlog.E(err)
	for _, mText := range rev {
		tmp := Message{}
		err = json.Unmarshal([]byte(mText), &tmp)
		levlog.E(err)
		if err != nil {
			levlog.Error("Msg is:", string(mText))
			return
		}
		msg.Msgs = append(msg.Msgs, tmp)
	}
	_, err = conn.Do("LTRIM", key, listLen, -1)
	levlog.E(err)
	return
}

func (c *NatsServer) Store(msgs Messages) error {
	conn := c.RedisDB.Get()
	if conn == nil {
		levlog.Error("Error, cannot connect to redis server")
		return errors.New("No connection")
	} else {
		defer conn.Close()
	}
	msgMap := make(map[string]([]interface{}))
	levlog.Info(msgMap)
	for _, msg := range msgs.Msgs {
		key := c.GenMsgKey(msg.To)
		if _, ok := msgMap[key]; ok == false {
			msgMap[key] = make([]interface{}, 0)
		}
		payload, err := json.Marshal(msg)
		levlog.E(err)
		if err != nil {
			return err
		}
		levlog.Info(string(payload))
		_, file, line, _ := runtime.Caller(1)
		levlog.Info("Caller is : ", file, ":", line)
		msgMap[key] = append(msgMap[key], string(payload))
	}
	// TODO change this to a transaction or pipe
	for key, val := range msgMap {
		/* since redis in aixinwu is old , no multi insert is avil
		change to single add this time
		*/
		//tmpList := make([]interface{}, 0)
		//tmpList = append(tmpList, key)
		//tmpList = append(tmpList, val...)
		//levlog.Println("push cmd is :: ", tmpList)
		//_, err := conn.Do("RPUSH", tmpList...)
		//levlog.E(err)
		//if err != nil {
		//	return err
		//}
		////////////////////////////////// single add implementation
		for _, v := range val {
			_, err := conn.Do("RPUSH", key, v)
			levlog.E(err)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

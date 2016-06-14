package natsChatServer

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/dailing/levlog"
)

func TestStorage(t *testing.T) {
	levlog.Info("start")
	server := NewServer()
	msg := Messages{
		Msgs: make([]Message, 0),
	}
	msg.Msgs = append(msg.Msgs, Message{
		Time:    time.Now(),
		Content: "sadf",
		From:    "1",
		To:      "0",
	})
	msg.Msgs = append(msg.Msgs, Message{
		Time:    time.Now(),
		Content: "yhsth",
		From:    "1",
		To:      "0",
	})
	server.Store(msg)
	s, _ := server.FetchMessage("0")
	server.Store(s)
	s, _ = server.FetchMessage("0")

	assert.Equal(t, s, msg)

}

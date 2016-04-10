echo "{\"mataData\":{\"TimeStamp\":1459409899384623648,\"Device\":\"test\"},\"userinfo\":{\"$1\":\"test\",\"password\":\"$2\",\"ID\":\"\",\"Phone\":\"\",\"Email\":\"\"}}"
curl localhost:12345/login --data "{\"mataData\":{\"TimeStamp\":1459409899384623648,\"Device\":\"test\"},\"userinfo\":{\"username\":\"$1\",\"password\":\"$2\",\"ID\":\"\",\"Phone\":\"\",\"Email\":\"\"}}" -v

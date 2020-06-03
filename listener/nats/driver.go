package nats

import (
	"log"
	"polygon/listener"
	"polygon/logger"
	logArt "polygon/logger/art"
	"github.com/nats-io/nats.go"
)

func a() {
	logger.Fatal(logArt.Msg{
		Id:     "1",
		Path:   "listener/nats/driver.go",
		Msg:    "失败",
		Format: "",
		Ext:    nil,
	})
}

var NatsConn *nats.Conn

type NatsSvc struct {}

func init() {
	listener.Register("nats", &NatsSvc{})
}

func (l *NatsSvc)IsUp() bool {
	if NatsCfgVal.IsUp {
		return true
	}
	return false
}

func (l *NatsSvc)Listen() {

	var err error
	// 连接服务
	NatsConn, err = nats.Connect(NatsCfgVal.Addr)
	if err != nil {
		log.Panicf("Nats Panic: %v\n", err)
	}

	logger.Info(logArt.Msg{
		Id:     NatsConn.ConnectedServerId(),
		Path:   "listenner/nats/driver.go",
		Msg:    "nats connect",
		Format: "",
		Ext:    nil,
	})
}

func (l *NatsSvc)UnListen() {
	NatsConn.Close()
}




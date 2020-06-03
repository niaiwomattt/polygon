package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/golang/protobuf/proto"
	"polygon/logger"
	logArt "polygon/logger/art"
)

var hs map[string]func()

func A(subject string, msg *nats.Msg)  {
	
}

func DecoratorProtobuf(path func(),pb proto.Message)  {
	
}

func b(msg *nats.Msg, pb proto.Message)  {
	pb = new(proto.Message)
	err := proto.Unmarshal(msg.Data, pb)
	if err != nil {
		logger.Error(logArt.Msg{
			Id:     "",
			Path:   "",
			Msg:    "",
			Format: "",
			Ext:    nil,
		})
	}

}

//svcNats.Nats.Subscribe(config.GetNats().Subject[n], func(msg *nats.Msg) {
//	if !vt.newType(msg) {
//		return
//	}
//	if !vt.request() {
//		return
//	}
//	vt.process()
//})
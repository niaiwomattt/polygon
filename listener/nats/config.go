package nats

import "polygon/configer"

type NatsCfg struct {
	IsUp    bool
	Addr    string
	Subject map[string]string
}
var NatsCfgVal *NatsCfg

func init()  {
	NatsCfgVal = &NatsCfg{}
	configer.Register("natsSvc", NatsCfgVal)
}

func (c *NatsCfg)GetConfig(file string) configer.Configer {
	return NatsCfgVal
}


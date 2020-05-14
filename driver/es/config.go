package es

import "polygon/configer"

type EsCfg struct {
	Addr string
}

var esVal *EsCfg

func init()  {
	esVal = &EsCfg{}
	configer.Register("es", esVal)
}

func (e *EsCfg)GetConfig(file string) configer.Configer {
	return esVal
}
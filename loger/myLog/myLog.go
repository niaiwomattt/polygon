package myLog

import (
	"github.com/sirupsen/logrus"
)

// 我想实现和接口分开

type Log struct {

}

// 接口和实现分开，但是又要符合这个消息结构
// 另外这个字段定义只是简单放置
type Fatal struct {
	Msg string
}

type Error struct {
	Fatal
	More string
}

type Warn struct {
	Error
	Data interface{}
}

type Info struct {
	Warn
	Info string
}

type Debug struct {
	Info
	Line string
}

func (l *Log)Fatal(f interface{})  {
	logrus.Fatal(f)
}

func (l *Log)Error(e interface{})  {

}

func (l *Log)Warn(e interface{})  {

}

func (l *Log)Info(e interface{})  {

}

func (l *Log)Debug(e interface{})  {

}
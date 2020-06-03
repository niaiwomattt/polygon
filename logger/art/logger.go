package art

import "github.com/sirupsen/logrus"

// 我想实现和接口分开
type Log struct {}

func (l *Log)Fatal(i ...interface{})  {
	logrus.Fatal(i)
}

func (l *Log)Error(i ...interface{})  {
	logrus.Error(i)
}

func (l *Log)Warn(i ...interface{})  {
	logrus.Warn(i)
}

func (l *Log)Info(i ...interface{})  {
	logrus.Info(i)
}

// 可以加上发生的行，等非常详细的地方，在调试时候可以存在的信息，比如文件名等，在特殊编译后就不存在了，这种东西就只有 debug 模式才能打出来
func (l *Log)Debug(i ...interface{})  {
	logrus.Debug(i)
}

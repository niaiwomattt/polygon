package loger

// 定义日志接口
// 此接口应当严格实现，因为要在应用中调用

type Loger interface {
	Fatal(interface{}) // 致命，此类错误会如果发生，当前应用将崩溃，不解决无法启动应用
	Error(interface{}) // 错误，不会崩溃，但是影响使用，必须马上解决。
	Warn(interface{}) // 警告，系统能继续运行, 但是必须引起关注，系统服务都可用，但是数据可能有问题
	Info(interface{}) // 信息，这个应该用来反馈系统的当前状态给最终用户的，所以，在这里输出的信息，应该对最终用户具有实际意义，也就是最终用户要能够看得明白是什么意思才行。 从某种角度上说，Info 输出的信息可以看作是软件产品的一部分（就像那些交互界面上的文字一样），所以需要谨慎对待，不可随便。
	Debug(interface{}) // 调试，比info 多更多相关信息，比如，模块，行，等等
}

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

func (l *Log)Fatal(f Fatal)  {

}

func (l *Log)Error(e Error)  {

}

func (l *Log)Warn(e Error)  {

}

func (l *Log)Info(e Error)  {

}

func (l *Log)Debug(e Error)  {

}

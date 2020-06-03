package logger

import "time"

// 定义日志接口
// 此接口应当严格实现，因为要在应用中调用

// 这些信息是每个日志里都必须包含的
// 整体说来就是，谁，什么时间，在哪儿，发生了什么事情，具体详情
// Who 谁，唯一标识
// When 什么时间，时间记录 Time time.Time // 时间默认都有，不需要另外加
// Where 在哪儿，文件路径，行，列
// What 发生了什么，日志消息，这个是外部传入的
// Context 详情，这个不是必须的，可以当做额外参数处理

//
// 格式定义： time path|URI id  msg ext...
// 格式解析： when where    who what context
type Logger interface {
	Fatal(...interface{}) // 致命，此类错误会如果发生，当前应用将崩溃，不解决无法启动应用
	Error(...interface{}) // 错误，不会崩溃，但是影响使用，必须马上解决。
	Warn(...interface{}) // 警告，系统能继续运行, 但是必须引起关注，系统服务都可用，但是数据可能有问题
	Info(...interface{}) // 信息，这个应该用来反馈系统的当前状态给最终用户的，所以，在这里输出的信息，应该对最终用户具有实际意义，也就是最终用户要能够看得明白是什么意思才行。 从某种角度上说，Info 输出的信息可以看作是软件产品的一部分（就像那些交互界面上的文字一样），所以需要谨慎对待，不可随便。
	Debug(...interface{}) // 调试，比info 多更多相关信息，比如，模块，行，等等
}

type Message interface {
	When() time.Time
	Where() string
	Who() string
	What() string
	Context() string
}

var drive Logger

// 注册日志驱动
func RegisterDrive(l Logger)  {
	drive = l
}

func Fatal(i ...interface{})  {
	format(drive.Fatal, i)
}

func Error(i ...interface{})  {
	format(drive.Error, i)
}

func Warn(i ...interface{})  {
	format(drive.Warn, i)
}

func Info(i ...interface{})  {
	format(drive.Info, i)
}

func Debug(i ...interface{})  {
	format(drive.Debug, i)
}

func format(f func(...interface{}), i ...interface{})  {
	m := isMessage(i)
	if m != nil {
		f(m.When(), m.Where(), m.Who(), m. What(), m.Context())
	}
	f(i)
}

func isMessage(i ...interface{}) Message {
	if len(i) > 0 {
		if v,ok := i[0].(Message); ok {
			return v
		}
	}
	return nil
}
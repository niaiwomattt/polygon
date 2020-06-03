package art

import (
	"fmt"
	"time"
)

// 这些信息是每个日志里都必须包含的
// 整体说来就是，谁，什么时间，在哪儿，发生了什么事情，具体详情
// Who 谁，唯一标识
// When 什么时间，时间记录 Time time.Time // 时间默认都有，不需要另外加
// Where 在哪儿，文件路径，行，列
// What 发生了什么，日志消息，这个是外部传入的
// Context 详情，这个不是必须的，可以当做额外参数处理

//
// 格式定义： time path|URI id msg ext...

// 接口和实现分开，但是又要符合这个消息结构
// 另外这个字段定义只是简单放置
type Msg struct {
	Id  string  // who
	Path string // where
	Msg string  // what
	Format string // context format
	Ext []interface{} // context
}

func (m *Msg)When() time.Time  {
	return time.Now()
}

func (m *Msg)Where() string  {
	return m.Path
}

func (m *Msg)Who() string {
	return m.Id
}

func (m *Msg)What() string  {
	return m.Msg
}

func (m *Msg)Context() string  {
	return fmt.Sprintf(m.Format, m.Ext)
}

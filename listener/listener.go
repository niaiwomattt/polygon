package listener

// 定义监听器接口
// 此接口应该严格实现
// 统一调用监听

type Listener interface {
	Listen()
	IsUp() bool
	UnListen()
}

var bind map[string]Listener

func Register(name string, listen Listener)  {
	bind[name] = listen
}

func Listen()  {
	for k, v := range bind {
		_ = k
		if !v.IsUp() {
			continue
		}
		v.Listen()
	}
}

func UnListen()  {
	for k, v := range bind {
		_ = k
		if !v.IsUp() {
			continue
		}
		v.UnListen()
	}
}
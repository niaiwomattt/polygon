package driver

// 定义驱动接口
// 此接口应当严格实现，不返回东西，但是要驱动成功
// 不在接口层统一调用，因为启动是谁用谁调

type Driver interface {
	Connect()
	Close()
}

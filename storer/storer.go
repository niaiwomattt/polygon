package storer

// 定义存储接口
// 此接口不用严格实现，只需要对象包含对应方法就行，传参返回值不用考虑，我们不是为了go语言上的接口实现
// 而是为了逻辑上的接口实现，统一一点而已
// 只包含数据操作
type Storer interface {
	Get()
	List()
	Create()
	Update()
}

package configer


// 配置接口
// 此接口应当严格实现，这样统一配置时才能成功
type Configer interface {
	GetConfig(file string) interface{}
}

type Cfg struct {

}

var fileList map[string]string

var cfgVal map[string]Configer

// 注册配置
func Register(name string, c Configer)  {
	cfgVal[name] = c
}

// 获取对应配置
func Get(name string) interface{}  {
	v := cfgVal[name]
	if v != nil {
		return v
	}
	return v.GetConfig(fileList[name])
}

// 获取基本配置
func InitConfig() {

}
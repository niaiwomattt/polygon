package configer

import (
	"github.com/BurntSushi/toml"
	"polygon/logger"
)

// 配置接口
// 此接口应当严格实现，这样统一配置时才能成功
type Configer interface {
	GetConfig(file string) Configer
}


var fileList map[string]string

var cfgVal map[string]Configer

// 注册配置
func Register(name string, c Configer) {
	if len(cfgVal) == 0 {
		cfgVal = make(map[string]Configer)
	}
	cfgVal[name] = c
}

// 获取对应配置
func Get(name string) Configer {
	v := cfgVal[name]
	if v != nil {
		return v
	}
	return v.GetConfig(fileList[name])
}


// 获取基本配置
func InitConfig() {
	_, err := toml.DecodeFile("config.toml", &fileList)
	if err != nil {
		loger.Fatal(err)
	}

	for k, _ := range cfgVal {
		toml.DecodeFile(fileList[k],cfgVal[k])
	}
}
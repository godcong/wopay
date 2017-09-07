package config

import (
	"flag"

	"io/ioutil"

	"log"

	toml "github.com/pelletier/go-toml"
)

var tomlConfig = flag.String("config", "../config.toml", "toml config")

type Wechat struct {
	AppId       string
	MchId       string
	AppSecret   string
	RedirectUri string
}

type Config struct {
	Wechat Wechat
}

var config *Config

func init() {
	ConfigInstance()
}
func loadConfig() *toml.Tree {

	tree, err := toml.LoadFile(*tomlConfig)
	if err != nil {
		return nil
	}
	return tree
}

func ConfigInstance() *Config {
	if config == nil {
		cfg := &Config{}
		//tree := loadConfig()
		data, _ := ioutil.ReadFile(*tomlConfig)
		log.Println(string(data))
		toml.Unmarshal(data, cfg)
		config = cfg
	}
	return config
}

func GetWechat() Wechat {
	//wechat := make(Wechat)
	//for k, v := range config. {
	//wechat[k] = InterfaceToString(v)
	//}

	return ConfigInstance().Wechat
}

func InterfaceToString(v interface{}) string {
	if v, b := v.(string); b {
		return v
	}
	return ""
}

//
//func (w Wechat) Key() string {
//	return w["key"]
//}
//
//func (w Wechat) AppId() string {
//	return w["appid"]
//}
//
//func (w Wechat) AppSecret() string {
//	return w["appsecret"]
//}
//
//func (w Wechat) MchId() string {
//	return w["mchid"]
//}
//
//func (w Wechat) RedirectUrl() string {
//	return w["redirect_url"]
//}

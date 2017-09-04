package wopay

import (
	"flag"

	"github.com/godcong/wopay/wx"
	toml "github.com/pelletier/go-toml"
)

var tomlConfig = flag.String("config", "config.toml", "toml config")

type Config struct {
	Wechat wx.PayConfig
}

func LoadConfig() *toml.Tree {
	tree, err := toml.LoadFile(*tomlConfig)
	if err != nil {
		return nil
	}
	return tree
}

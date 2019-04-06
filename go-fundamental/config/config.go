package config

import (
	"flag"
	"io/ioutil"
	"sync/atomic"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

var configFilePath string
var globalConfig atomic.Value

func init() {
	flag.StringVar(&configFilePath, "conf", "", "configuration file path")
}

// Initialize 会从configFilePath文件中读取yaml配置
func Initialize(config interface{}) interface{} {
	if !flag.Parsed() {
		flag.Parse()
	}

	if configFilePath != "" {
		dat, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			glog.Fatalf("read config file %v error. err = %v", configFilePath, err)
		}
		err = yaml.Unmarshal(dat, config)
		if err != nil {
			glog.Infof("parse config file %v error. err = %v", configFilePath, err)
		}

		glog.Infof("config initialize success. config = %v\n", config)

		globalConfig.Store(config)
	} else {
		glog.Infof("config file is null. no -c ??")
	}
	return globalConfig.Load()
}

// Get 返回config实例
func Get() interface{} {
	return globalConfig.Load()
}

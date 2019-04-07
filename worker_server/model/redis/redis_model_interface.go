package redismodel

import (
	"github.com/golang/glog"
	"github.com/palanceli/MVCSample/worker_server/config"
)

// RedisModel Redis在model层暴露的接口
type RedisModel interface {
	SetData(dataType int32, content string) error
	GetData(dataType int32) (content string, err error)
}

// CreateRedisModel 接口工厂
func CreateRedisModel(conf *config.WorkerConfig) RedisModel {
	m := &redisStore{}
	if err := m.init(conf); err != nil {
		glog.Fatalf("init MySQL failed. err = %v", err)
	}
	return m
}

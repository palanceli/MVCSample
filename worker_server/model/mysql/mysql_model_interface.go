package mysqlmodel

import (
	"github.com/golang/glog"
	"github.com/palanceli/MVCSample/worker_server/config"
)

// MySQLModel MySQL在model层暴露的接口
type MySQLModel interface {
	SaveData(dataType int32, content string) error
	QueryData(dataType int32) (string, error)
	SyncTables() error
}

// CreateMysqlModel 接口工厂
func CreateMysqlModel(conf *config.WorkerConfig) MySQLModel {
	m := &mysqlStore{}
	if err := m.init(conf); err != nil {
		glog.Fatalf("init MySQL failed. err = %v", err)
	}
	return m
}

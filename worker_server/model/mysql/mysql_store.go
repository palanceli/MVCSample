package mysqlmodel

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 如果不import该行，连接数据库时会提示找不到driverName
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/palanceli/MVCSample/worker_server/config"
	"github.com/palanceli/MVCSample/worker_server/model"
)

type mysqlStore struct {
	engine *xorm.Engine
}

func (m *mysqlStore) init(conf *config.WorkerConfig) error {
	e, err := xorm.NewEngine("mysql", conf.MySQLAddr)
	if err != nil {
		return fmt.Errorf("FAILED to Create Xorm Engine. err=%v", err)
	}
	e.ShowSQL(conf.ShowSQL)
	m.engine = e
	return err
}

func (m *mysqlStore) SaveData(dataType int32, content string) error {
	workerData := &model.WorkerData{
		DataType: model.GetDataType(dataType),
		Content:  content,
		Time:     time.Now(),
	}
	_, err := m.engine.Insert(workerData)
	if err != nil {
		glog.Fatalf("FAILED to save data to mysql. err=%v", err)
	}
	return nil
}

func (m *mysqlStore) QueryData(dataType int32) (string, error) {
	var workerData model.WorkerData
	getOK, err := m.engine.Where("data_type=?", model.GetDataType(dataType)).Desc("uptdate_time").Get(&workerData)
	if err != nil {
		glog.Fatalf("FAILED to query data from mysql. err=%v", err)
	}
	if !getOK {
		return "", fmt.Errorf("Data (dataType=%d) NOT exits", dataType)
	}
	return workerData.Content, nil
}

// SyncTables ..
func (m *mysqlStore) SyncTables() error {
	workData := &model.WorkerData{}
	err := m.engine.Sync2(workData)
	if err != nil {
		glog.Fatalf("FAILED to sync table. err=%v", err)
	}
	return nil
}

package mysqlmodel

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"

	"github.com/golang/glog"
	fundmentalconfig "github.com/palanceli/MVCSample/go-fundamental/config"
	"github.com/palanceli/MVCSample/worker_server/config"
)

var mysqlSQLModel MySQLModel

func TestMain(m *testing.M) {
	defer glog.Flush()

	flag.Set("logtostderr", "true")
	flag.Set("conf", "../../../config.yml")
	flag.Parse()

	fundmentalconfig.Initialize(&config.WorkerConfig{})
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)

	rand.Seed(time.Now().UnixNano())
	cfg.ShowSQL = true
	mysqlSQLModel = CreateMysqlModel(cfg)
	ret := m.Run()
	os.Exit(ret)
}

func TestSaveAndQueryData(t *testing.T) {
	dataType := rand.Int31()
	content := uuid.New().String()
	err := mysqlSQLModel.SaveData(dataType, content)
	assert.Nil(t, err, "FAILED to SaveData into mysql.")
	queryContent, err := mysqlSQLModel.QueryData(dataType)
	assert.Nil(t, err, "FAILED to queryData from mysql.")
	assert.Equal(t, content, queryContent, "SaveData != QueryData")
}

func TestSyncTables(t *testing.T) {
	err := mysqlSQLModel.SyncTables()
	assert.Nil(t, err, "FAILED to SyncTables")
}

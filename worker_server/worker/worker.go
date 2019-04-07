package worker

import (
	"sync"
	"time"

	redismodel "github.com/palanceli/MVCSample/worker_server/model/redis"

	"github.com/golang/glog"
	"github.com/palanceli/MVCSample/worker_server/config"
	mysqlmodel "github.com/palanceli/MVCSample/worker_server/model/mysql"
)

var workerOnce sync.Once
var worker *Worker

// Worker 业务逻辑的实现者
type Worker struct {
	cfg        *config.WorkerConfig
	redisModel redismodel.RedisModel
	mysqlModel mysqlmodel.MySQLModel
	notifier   *notifier
}

// CreateWorker 工厂方法
func CreateWorker(cfg *config.WorkerConfig) *Worker {
	workerOnce.Do(func() {
		worker = &Worker{
			cfg:        cfg,
			redisModel: redismodel.CreateRedisModel(cfg),
			mysqlModel: mysqlmodel.CreateMysqlModel(cfg),
			notifier: &notifier{cfg: cfg,
				timeout: 5 * time.Second,
			},
		}
	})
	return worker
}

// SetData ...
func (w *Worker) SetData(dataType int32, content string) error {
	err := w.redisModel.SetData(dataType, content)
	if err != nil {
		glog.Fatalf("FAILED to save data to redis. err=%v", err)
	}
	err = w.mysqlModel.SaveData(dataType, content)
	if err != nil {
		glog.Fatalf("FAILED to save data to mysql. err=%v", err)
	}

	err = w.notifier.Notify(dataType, content)
	if err != nil {
		glog.Fatalf("FAILED to notify receiver. err=%v", err)
	}
	return nil
}

// GetData ...
func (w *Worker) GetData(dataType int32) (string, error) {
	content, err := w.redisModel.GetData(dataType)
	if err != nil {
		glog.Fatalf("FAILED to GetData from redis. err=%v", err)
	}
	return content, nil
}

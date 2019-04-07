package worker

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

func TestMain(m *testing.M) {
	defer glog.Flush()

	flag.Set("logtostderr", "true")
	flag.Set("conf", "../../config.yml")
	flag.Parse()

	fundmentalconfig.Initialize(&config.WorkerConfig{})
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)

	rand.Seed(time.Now().UnixNano())
	cfg.ShowSQL = true
	ret := m.Run()
	os.Exit(ret)
}

func TestSetAndGetData(t *testing.T) {
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)
	worker := CreateWorker(cfg)
	dataType := rand.Int31()
	content := uuid.New().String()
	err := worker.SetData(dataType, content)
	assert.Nil(t, err, "FAILED to SetData")
	queryContent, err := worker.GetData(dataType)
	assert.Nil(t, err, "FAILED to GetData")
	assert.Equal(t, content, queryContent, "")
}

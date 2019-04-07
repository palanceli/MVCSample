package worker

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	fundmentalconfig "github.com/palanceli/MVCSample/go-fundamental/config"
	"github.com/palanceli/MVCSample/worker_server/config"
)

func TestNotify(t *testing.T) {
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)
	notifier := &notifier{cfg: cfg,
		timeout: 5 * time.Second,
	}
	dataType := rand.Int31()
	content := uuid.New().String()
	notifier.Notify(dataType, content)
}

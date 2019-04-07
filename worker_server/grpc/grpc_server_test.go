package grpc

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"git.windimg.com/giantthong/go-fundamental/credentials"
	"github.com/golang/glog"
	fundmentalconfig "github.com/palanceli/MVCSample/go-fundamental/config"
	serverhelper "github.com/palanceli/MVCSample/go-fundamental/server_helper"
	worker_api "github.com/palanceli/MVCSample/worker"
	"github.com/palanceli/MVCSample/worker_server/config"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func init() {
	glog.Info("Visiting...")
	flag.Visit(func(i *flag.Flag) {
		glog.Infof("(%s, %v)", i.Name, i.Value)
	})
	flag.Set("logtostderr", "true")
	flag.Set("conf", "../../config.yml")
	flag.Parse()
}

func TestMain(m *testing.M) {
	defer glog.Flush()

	fundmentalconfig.Initialize(&config.WorkerConfig{})
	rand.Seed(time.Now().UnixNano())

	ret := m.Run()
	os.Exit(ret)
}

func getGRPCClientConn(connCtx context.Context, grpcAddr string, grpcProxyTLSOn bool) (*grpc.ClientConn, error) {
	var options []grpc.DialOption
	options = append(options, grpc.WithBlock())

	creds, err := credentials.DialOption(options, grpcProxyTLSOn)
	if err != nil {
		return nil, fmt.Errorf("FAILED to create TLS credentials. err=%v", err)
	}
	return grpc.DialContext(connCtx, grpcAddr, creds...)
}

func TestSetDataAndGetData(t *testing.T) {
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)

	// 启动server
	grpcServer := NewServer(cfg)
	go func() {
		grpcServer.Run(serverhelper.SignalContext(context.Background()))
	}()

	// 调用grpc接口
	// 建立连接
	conn, err := getGRPCClientConn(context.Background(), cfg.GrpcAddr, false)
	assert.Nil(t, err)
	defer conn.Close()

	// 执行GRPC接口
	client := worker_api.NewWorkerServiceClient(conn)
	setDataRequest := worker_api.SetDataRequest{
		Type:    123,
		Content: "test_123",
	}
	setDataReply, err := client.SetData(context.Background(), &setDataRequest)
	assert.Nil(t, err, "FAILED to SetData")
	assert.Equal(t, setDataReply.Status, int32(200), "status != 0")
	t.Logf("Status=%d", setDataReply.Status)

	getDataRequest := worker_api.GetDataRequest{
		Type: 123,
	}
	getDataReply, err := client.GetData(context.Background(), &getDataRequest)
	assert.Nil(t, err, "FAILED to GetData")
	assert.Equal(t, setDataRequest.Content, getDataReply.Msg, "get data != set data")
}

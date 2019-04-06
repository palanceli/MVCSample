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
	receiver_api "github.com/palanceli/MVCSample/receiver"
	"github.com/palanceli/MVCSample/receiver_server/config"
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

	fundmentalconfig.Initialize(&config.ReceiverConfig{})
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

func TestReceiveData(t *testing.T) {
	cfg := fundmentalconfig.Get().(*config.ReceiverConfig)

	// 启动server
	grpcServer := NewServer(cfg)
	go func() {
		grpcServer.Run(serverhelper.SignalContext(context.Background()))
	}()

	// 调用grpc接口
	timeOut := 10 * time.Second
	connCtx, connCancel := context.WithTimeout(context.Background(), timeOut)
	defer connCancel()

	// 建立连接
	conn, err := getGRPCClientConn(connCtx, cfg.GrpcAddr, false)
	assert.Nil(t, err)
	defer conn.Close()

	// 执行GRPC接口
	client := receiver_api.NewReceiverServiceClient(conn)
	receiveDataRequest := receiver_api.ReceiveDataRequest{
		Type:    1,
		Content: "test",
	}
	callCtx, callCancel := context.WithTimeout(context.Background(), timeOut)
	defer callCancel()
	receiveDataReply, err := client.ReceiveData(callCtx, &receiveDataRequest)
	assert.Equal(t, receiveDataReply.Status, int32(200), "status != 0")
	t.Logf("Receive status=%d, msg=%s", receiveDataReply.Status, receiveDataReply.Msg)
}

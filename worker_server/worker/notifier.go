package worker

import (
	"context"
	"fmt"
	"time"

	"git.windimg.com/giantthong/go-fundamental/credentials"
	"github.com/golang/glog"
	receiver_api "github.com/palanceli/MVCSample/receiver"
	"github.com/palanceli/MVCSample/worker_server/config"
	"google.golang.org/grpc"
)

type notifier struct {
	cfg     *config.WorkerConfig
	timeout time.Duration
}

func (n *notifier) getGRPCClientConn(connCtx context.Context, grpcAddr string, grpcProxyTLSOn bool) (*grpc.ClientConn, error) {
	var options []grpc.DialOption
	options = append(options, grpc.WithBlock())

	creds, err := credentials.DialOption(options, grpcProxyTLSOn)
	if err != nil {
		return nil, fmt.Errorf("FAILED to create TLS credentials. err=%v", err)
	}
	return grpc.DialContext(connCtx, grpcAddr, creds...)
}

func (n *notifier) getReceiverClient() (*grpc.ClientConn, receiver_api.ReceiverServiceClient) {
	// 调用grpc接口
	connCtx, connCancel := context.WithTimeout(context.Background(), n.timeout)
	defer connCancel()

	// 建立连接
	conn, err := n.getGRPCClientConn(connCtx, n.cfg.ReceiverGrpcAddr, false)
	if err != nil {
		glog.Fatalf("FAILED to connect receiver. err=%v", err)
	}

	// 执行GRPC接口
	return conn, receiver_api.NewReceiverServiceClient(conn)
}

func (n *notifier) Notify(dataType int32, content string) error {
	conn, receiverClient := n.getReceiverClient()
	defer conn.Close()

	receiveDataRequest := receiver_api.ReceiveDataRequest{
		Type:    dataType,
		Content: content,
	}
	callCtx, callCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer callCancel()

	receiveDataReply, err := receiverClient.ReceiveData(callCtx, &receiveDataRequest)
	if err != nil {
		glog.Fatalf("FAILED to notify receiver. err=%v", err)
	}
	glog.Infof("Receive status=%d, msg=%s", receiveDataReply.Status, receiveDataReply.Msg)
	return nil
}

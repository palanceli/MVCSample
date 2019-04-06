package worker

import (
	"context"
	"fmt"
	"time"

	"git.windimg.com/giantthong/go-fundamental/credentials"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	receiver_api "github.com/palanceli/MVCSample/receiver"
	"google.golang.org/grpc"
)

// WorkerData ...
type WorkerData struct {
	ID       int       `xorm:"'id' pk autoincr"`
	DataType string    `xorm:"data_type" json:"data_type"`
	Content  []byte    `xorm:"data_content" json:"data_content"`
	Time     time.Time `xorm:"uptdate_time"`
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

func generateKey(dataType int32) string {
	return fmt.Sprintf("data_type_%04d", dataType)
}

// SetData ...
func SetData(dataType int32, content string, redisClient *redis.Client, mysqlEngine *xorm.Engine, receiverGrpcAddr string) error {
	key := generateKey(dataType)
	_, err := redisClient.Set(key, content, 0).Result()
	if err != nil {
		glog.Fatalf("FAILED to save data to redis. err=%v", err)
	}
	workerData := &WorkerData{
		DataType: key,
		Content:  []byte(content),
		Time:     time.Now(),
	}
	_, err = mysqlEngine.Insert(workerData)
	if err != nil {
		glog.Fatalf("FAILED to save data to mysql. err=%v", err)
	}

	// 调用grpc接口
	timeOut := 10 * time.Second
	connCtx, connCancel := context.WithTimeout(context.Background(), timeOut)
	defer connCancel()

	// 建立连接
	conn, err := getGRPCClientConn(connCtx, receiverGrpcAddr, false)
	if err != nil {
		glog.Fatalf("FAILED to connect receiver. err=%v", err)
	}
	defer conn.Close()

	// 执行GRPC接口
	client := receiver_api.NewReceiverServiceClient(conn)
	receiveDataRequest := receiver_api.ReceiveDataRequest{
		Type:    dataType,
		Content: content,
	}
	callCtx, callCancel := context.WithTimeout(context.Background(), timeOut)
	defer callCancel()
	receiveDataReply, err := client.ReceiveData(callCtx, &receiveDataRequest)
	if err != nil {
		glog.Fatalf("FAILED to notify receiver. err=%v", err)
	}
	glog.Infof("Receive status=%d, msg=%s", receiveDataReply.Status, receiveDataReply.Msg)
	return nil
}

// GetData ...
func GetData(dataType int32, redisClient *redis.Client) (string, error) {
	key := generateKey(dataType)
	content, err := redisClient.Get(key).Result()
	if err != nil {
		glog.Fatalf("FAILED to GetData from redis. err=%v", err)
	}
	return content, nil
}

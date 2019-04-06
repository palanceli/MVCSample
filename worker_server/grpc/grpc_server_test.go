package grpc

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"git.windimg.com/giantthong/go-fundamental/credentials"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	fundmentalconfig "github.com/palanceli/MVCSample/go-fundamental/config"
	serverhelper "github.com/palanceli/MVCSample/go-fundamental/server_helper"
	worker_api "github.com/palanceli/MVCSample/worker"
	"github.com/palanceli/MVCSample/worker_server/config"
	"github.com/palanceli/MVCSample/worker_server/worker"
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

func parseRedisAddr(addr string) (host string, password string, db int, err error) {
	db = 0
	u, err := url.Parse(addr)
	if err != nil {
		return "", "", 0, fmt.Errorf("FAILED to parse redis addr. err=%v", err)
	}
	host = u.Host
	db64, _ := strconv.ParseInt(u.User.Username(), 0, 32)

	db = int(db64)
	password, _ = u.User.Password()

	glog.V(8).Infof("parse redis URI. addr = %s, host = %s, password = %s, db = %d", addr, host, password, db)
	return host, password, db, nil
}

func initRedis(cfg *config.WorkerConfig) *redis.Client {
	host, passwd, db, err := parseRedisAddr(cfg.RedisAddr)
	if err != nil {
		glog.Fatalf("FAILED to init redis. err=%v", err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passwd,
		DB:       db,
	})
}

func initMySQL(cfg *config.WorkerConfig) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", cfg.MySQLAddr)
	if err != nil {
		glog.Fatalf("FAILED to init mysql. err=%v", err)
	}
	return engine
}

func TestSetDataAndGetData(t *testing.T) {
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)

	redisClient := initRedis(cfg)
	mysqlEngine := initMySQL(cfg)
	// 启动server
	grpcServer := NewServer(cfg, redisClient, mysqlEngine)
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

func TestSync(t *testing.T) {
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)
	mysqlEngine := initMySQL(cfg)
	mysqlEngine.ShowSQL(true)
	err := mysqlEngine.Sync2(&worker.WorkerData{})
	assert.Nil(t, err, "FAILED to sync mysql table")
}

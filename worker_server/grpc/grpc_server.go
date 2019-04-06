package grpc

import (
	"context"
	"net"
	"time"

	"github.com/go-xorm/xorm"

	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	worker_api "github.com/palanceli/MVCSample/worker"
	"github.com/palanceli/MVCSample/worker_server/config"
	"github.com/palanceli/MVCSample/worker_server/worker"
)

const maxsize = 1024 * 1024 * 10

// Server ...
type Server struct {
	cfg         *config.WorkerConfig
	redisClient *redis.Client
	mysqlEngine *xorm.Engine
}

// SetData ...
func (s *Server) SetData(ctx context.Context, request *worker_api.SetDataRequest) (*worker_api.SetDataReply, error) {
	glog.Infof("Set data (Type=%d, content=%s)", request.Type, request.Content)
	err := worker.SetData(request.Type, request.Content, s.redisClient, s.mysqlEngine, s.cfg.ReceiverGrpcAddr)
	if err != nil {
		glog.Fatalf("FAILED to SetData. err=%v", err)
	}
	return &worker_api.SetDataReply{Status: 200, Msg: "OK"}, nil
}

// GetData ...
func (s *Server) GetData(ctx context.Context, request *worker_api.GetDataRequest) (*worker_api.GetDataReply, error) {
	glog.Infof("Get data (Type=%d)", request.Type)
	content, err := worker.GetData(request.Type, s.redisClient)
	if err != nil {
		glog.Fatalf("FAILED to GetData. err=%v", err)
	}
	reply := worker_api.GetDataReply{Msg: content}
	return &reply, nil
}

// NewServer ...
func NewServer(cfg *config.WorkerConfig, redisClient *redis.Client, mysqlEngine *xorm.Engine) *Server {
	return &Server{
		cfg:         cfg,
		redisClient: redisClient,
		mysqlEngine: mysqlEngine,
	}
}

// Run ...
func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	glog.Infof("starting grpc server on %s", s.cfg.GrpcAddr)

	opts := []grpc.ServerOption{grpc.MaxRecvMsgSize(maxsize), grpc.MaxSendMsgSize(maxsize)}
	opts = append(opts,
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    20 * time.Second,
				Timeout: 10 * time.Second,
			},
		))
	srv := grpc.NewServer(opts...)
	//此处注册一下grpc proto定义的方法
	worker_api.RegisterWorkerServiceServer(srv, s)

	l, err := net.Listen("tcp", s.cfg.GrpcAddr)
	if err != nil {
		glog.Fatalf("FAILED to start grpc server. err=%v", err)
	}

	go func() {
		if err := srv.Serve(l); err != nil {
			glog.Infof("grpc serve:%v", err)
		}
		cancel()
	}()

	<-ctx.Done()

	glog.Infof("grpc server shutting down")

	srv.Stop()
	return nil
}

package grpc

import (
	"context"
	"net"
	"time"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	receiver_api "github.com/palanceli/MVCSample/receiver"
	"github.com/palanceli/MVCSample/receiver_server/config"
)

const maxsize = 1024 * 1024 * 10

// Server ...
type Server struct {
	cfg *config.ReceiverConfig
}

// ReceiveData 用于接收数据
func (s *Server) ReceiveData(ctx context.Context, request *receiver_api.ReceiveDataRequest) (*receiver_api.ReceiveDataReply, error) {
	glog.Infof("Received data (Type=%d, content=%s)", request.Type, request.Content)
	return &receiver_api.ReceiveDataReply{Status: 200, Msg: "OK"}, nil
}

// NewServer ...
func NewServer(cfg *config.ReceiverConfig) *Server {
	return &Server{
		cfg: cfg,
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
	receiver_api.RegisterReceiverServiceServer(srv, s)

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

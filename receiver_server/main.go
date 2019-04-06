package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	fundmentalconfig "github.com/palanceli/MVCSample/go-fundamental/config"
	serverhelper "github.com/palanceli/MVCSample/go-fundamental/server_helper"
	"github.com/palanceli/MVCSample/receiver_server/config"
	"github.com/palanceli/MVCSample/receiver_server/grpc"
)

func main() {
	fundmentalconfig.Initialize(&config.ReceiverConfig{})
	cfg := fundmentalconfig.Get().(*config.ReceiverConfig)

	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infof("config:%+v", cfg)
	glog.Flush()

	grpcServer := grpc.NewServer(cfg)
	grpcServer.Run(serverhelper.SignalContext(context.Background()))
}

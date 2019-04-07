package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"strconv"

	"github.com/go-xorm/xorm"

	"github.com/go-redis/redis"
	"github.com/golang/glog"
	fundmentalconfig "github.com/palanceli/MVCSample/go-fundamental/config"
	serverhelper "github.com/palanceli/MVCSample/go-fundamental/server_helper"
	"github.com/palanceli/MVCSample/worker_server/config"
	"github.com/palanceli/MVCSample/worker_server/grpc"
)

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

func main() {
	fundmentalconfig.Initialize(&config.WorkerConfig{})
	cfg := fundmentalconfig.Get().(*config.WorkerConfig)

	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infof("config:%+v", cfg)
	glog.Flush()

	grpcServer := grpc.NewServer(cfg)
	grpcServer.Run(serverhelper.SignalContext(context.Background()))
}

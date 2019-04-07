package redismodel

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"github.com/palanceli/MVCSample/worker_server/config"
	"github.com/palanceli/MVCSample/worker_server/model"
)

type redisStore struct {
	redisClient *redis.Client
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

func (m *redisStore) init(conf *config.WorkerConfig) error {
	host, passwd, db, err := parseRedisAddr(conf.RedisAddr)
	if err != nil {
		glog.Fatalf("FAILED to init redis. err=%v", err)
	}
	m.redisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passwd,
		DB:       db,
	})
	return nil
}

func (m *redisStore) SetData(dataType int32, content string) error {
	_, err := m.redisClient.Set(model.GetDataType(dataType), content, 0).Result()
	if err != nil {
		glog.Fatalf("FAILED to set data to redis. err=%v", err)
	}
	return err
}

func (m *redisStore) GetData(dataType int32) (content string, err error) {
	content, err = m.redisClient.Get(model.GetDataType(dataType)).Result()
	if err != nil {
		glog.Fatalf("FAILED to get data from redis. err=%v", err)
	}
	return content, nil
}

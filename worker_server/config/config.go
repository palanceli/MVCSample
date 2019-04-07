package config

// WorkerConfig ...
type WorkerConfig struct {
	GrpcAddr         string `yaml:"worker_grpc_addr"`
	RedisAddr        string `yaml:"worker_redis_addr"`
	MySQLAddr        string `yaml:"worker_mysql_addr"`
	ReceiverGrpcAddr string `yaml:"receiver_grpc_addr"`
	ShowSQL          bool   `yaml:"show_sql"`
}

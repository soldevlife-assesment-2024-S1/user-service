package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Redis      RedisConfig      `envconfig:"redis"`
	HttpServer HttpServerConfig `envconfig:"http_server"`
	HttpClient HttpClientConfig `envconfig:"http_client"`
	Logger     LoggerConfig     `envconfig:"logger"`
	Database   DatabaseConfig   `envconfig:"database"`
}

type DatabaseConfig struct {
	Host         string `envconfig:"database_host"`
	Port         int    `envconfig:"database_port"`
	Username     string `envconfig:"database_username"`
	Password     string `envconfig:"database_password"`
	DBName       string `envconfig:"database_db_name"`
	SSL          string `envconfig:"database_ssl"`
	SchemaName   string `envconfig:"database_schema_name"`
	MaxIdleConns int    `envconfig:"database_max_idle_conns"`
	MaxOpenConns int    `envconfig:"database_max_open_conns"`
	Timeout      int    `envconfig:"database_timeout"`
}

type RedisConfig struct {
	Host            string        `envconfig:"redis_host"`
	Port            string        `envconfig:"redis_port"`
	Username        string        `envconfig:"redis_username"`
	Password        string        `envconfig:"redis_password"`
	DB              int           `envconfig:"redis_db"`
	MaxRetries      int           `envconfig:"redis_max_retries"`
	PoolFIFO        bool          `envconfig:"redis_pool_fifo"`
	PoolSize        int           `envconfig:"redis_pool_size"`
	PoolTimeout     time.Duration `envconfig:"redis_pool_timeout"`
	MinIdleConns    int           `envconfig:"redis_min_idle_conns"`
	MaxIdleConns    int           `envconfig:"redis_max_idle_conns"`
	ConnMaxIdleTime time.Duration `envconfig:"redis_conn_max_idle_time"`
	ConnMaxLifetime time.Duration `envconfig:"redis_conn_max_lifetime"`
}

type HttpClientConfig struct {
	Host                string  `envconfig:"http_client_host"`
	Port                string  `envconfig:"http_client_port"`
	Timeout             int     `envconfig:"http_client_timeout"`
	ConsecutiveFailures int     `envconfig:"http_client_consecutive_failures"`
	ErrorRate           float64 `envconfig:"http_client_error_rate"` // 0.001 - 0.999
	Threshold           int     `envconfig:"http_client_threshold"`
	Type                string  `envconfig:"http_client_type"` // consecutive, error_rate
}

type HttpServerConfig struct {
	Host string `envconfig:"http_server_host"`
	Port string `envconfig:"http_server_port"`
}

type LoggerConfig struct {
	IsVerbose       bool   `envconfig:"logger_is_verbose"`
	LoggerCollector string `envconfig:"logger_logger_collector"`
}

func InitConfig() *Config {
	var Cfg Config

	err := envconfig.Process("", &Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Cfg
}

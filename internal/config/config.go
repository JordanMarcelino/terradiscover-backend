package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App        *AppConfig
	HttpServer *HttpServerConfig
	Database   *DatabaseConfig
	Jwt        *JwtConfig
	Logger     *LoggerConfig
}

type AppConfig struct {
	Environment string `mapstructure:"APP_ENVIRONMENT"`
	BCryptCost  int    `mapstructure:"APP_BCRYPT_COST"`
}

type HttpServerConfig struct {
	Host                 string `mapstructure:"HTTP_SERVER_HOST"`
	Port                 int    `mapstructure:"HTTP_SERVER_PORT"`
	GracePeriod          int    `mapstructure:"HTTP_SERVER_GRACE_PERIOD"`
	RequestTimeoutPeriod int    `mapstructure:"HTTP_SERVER_REQUEST_TIMEOUT_PERIOD"`
}

type DatabaseConfig struct {
	Host                  string `mapstructure:"DB_HOST"`
	DbName                string `mapstructure:"DB_NAME"`
	Username              string `mapstructure:"DB_USER"`
	Password              string `mapstructure:"DB_PASSWORD"`
	Sslmode               string `mapstructure:"DB_SSL_MODE"`
	Port                  int    `mapstructure:"DB_PORT"`
	MaxIdleConn           int    `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxOpenConn           int    `mapstructure:"DB_MAX_OPEN_CONN"`
	MaxConnLifetimeMinute int    `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

type JwtConfig struct {
	AllowedAlgs     []string `mapstructure:"JWT_ALLOWED_ALGS"`
	Issuer          string   `mapstructure:"JWT_ISSUER"`
	SecretKey       string   `mapstructure:"JWT_SECRET_KEY"`
	TokenDuration   int      `mapstructure:"JWT_TOKEN_DURATION"`
	RefreshDuration int      `mapstructure:"JWT_REFRESH_DURATION"`
}

type LoggerConfig struct {
	Level int `mapstructure:"LOGGER_LEVEL"`
}

func InitConfig() *Config {
	configPath := parseConfigPath()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	return &Config{
		App:        initAppConfig(),
		Database:   initDbConfig(),
		HttpServer: initHttpServerConfig(),
		Jwt:        initJwtConfig(),
		Logger:     initLoggerConfig(),
	}
}

func initLoggerConfig() *LoggerConfig {
	loggerConfig := &LoggerConfig{}

	if err := viper.Unmarshal(&loggerConfig); err != nil {
		log.Fatalf("error mapping logger config: %v", err)
	}

	return loggerConfig
}

func initJwtConfig() *JwtConfig {
	jwtConfig := &JwtConfig{}

	if err := viper.Unmarshal(&jwtConfig); err != nil {
		log.Fatalf("error mapping jwt config: %v", err)
	}

	return jwtConfig
}

func initDbConfig() *DatabaseConfig {
	dbConfig := &DatabaseConfig{}

	if err := viper.Unmarshal(&dbConfig); err != nil {
		log.Fatalf("error mapping database config: %v", err)
	}

	return dbConfig
}

func initHttpServerConfig() *HttpServerConfig {
	httpServerConfig := &HttpServerConfig{}

	if err := viper.Unmarshal(&httpServerConfig); err != nil {
		log.Fatalf("error mapping http server config: %v", err)
	}

	return httpServerConfig
}

func initAppConfig() *AppConfig {
	appConfig := &AppConfig{}

	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("error mapping app config: %v", err)
	}

	return appConfig
}

func parseConfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wd
}

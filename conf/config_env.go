package conf

import (
	"github.com/caarlos0/env/v8"
	"log"
)

var config Config

type Config struct {
	Port        string `envconfig:"PORT"`
	IsDebug     bool   `envconfig:"IS_DEBUG"`
	ServiceHost string `envconfig:"SERVICE_HOST"`

	MySQL struct {
		DBHost         string `env:"DB_HOST" envDefault:"localhost"`
		DBPort         string `env:"DB_PORT" envDefault:"5432"`
		DBUser         string `env:"DB_USER" envDefault:"mysql"`
		DBPass         string `env:"DB_PASS" envDefault:"1"`
		DBName         string `env:"DB_NAME" envDefault:"mysql"`
		EnableDB       string `env:"ENABLE_DB" envDefault:"true"`
		DBMaxIdleConns int    `env:"DB_MAX_IDLE_CONNS"`
		DBMaxOpenConns int    `env:"DB_MAX_OPEN_CONNS"`
		CountRetryTx   int    `env:"DB_TX_RETRY_COUNT"`

		DBTestHost string `env:"DB_TEST_HOST" envDefault:"localhost"`
		DBTestPort string `env:"DB_TEST_PORT" envDefault:"3306"`
		DBTestUser string `env:"DB_TEST_USER" envDefault:"test"`
		DBTestPass string `env:"DB_TEST_PASS" envDefault:"1"`
		DBTestName string `env:"DB_TEST_NAME" envDefault:"test"`
	}

	HealthCheck struct {
		CronJobFlag         bool   `envconfig:"CRON_JOB_FLAG"`
		HealthCheckEndPoint string `envconfig:"HEALTH_CHECK_ENDPOINT"`
	}

	AWSConfig struct {
		Region    string `envconfig:"AWS_REGION"`
		AccessKey string `envconfig:"AWS_ACCESS_KEY"`
		SecretKey string `envconfig:"AWS_SECRET_KEY"`
	}

	S3Config struct {
		KeyUUID    string `envconfig:"S3_KEY_UUID"`
		BucketName string `envconfig:"S3_BUCKET_NAME"`
		EndPoint   string `envconfig:"S3_ENDPOINT"`
		SiteURL    string `envconfig:"S3_SITE_URL"`
		DefaultDir string `envconfig:"S3_DEFAULT_DIR"`
	}

	SentryDSN      string `envconfig:"SENTRY_DSN"`
	TokenSecretKey string `envconfig:"TOKEN_SECRET_KEY"`
}

func GetConfig() *Config {
	return &config
}

func SetEnv() {
	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}
}

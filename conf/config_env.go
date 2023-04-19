package conf

var config *Config

type Config struct {
	Port        string `envconfig:"PORT"`
	IsDebug     bool   `envconfig:"IS_DEBUG"`
	ServiceHost string `envconfig:"SERVICE_HOST"`

	MySQL struct {
		Host           string `envconfig:"DB_HOST"`
		Port           string `envconfig:"DB_PORT"`
		User           string `envconfig:"DB_USER"`
		Pass           string `envconfig:"DB_PASS"`
		DBName         string `envconfig:"DB_NAME"`
		DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`
		DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
		CountRetryTx   int    `envconfig:"DB_TX_RETRY_COUNT"`
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
	return config
}

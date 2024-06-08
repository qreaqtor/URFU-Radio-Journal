package config

type ServerConfig struct {
	PostgresConfig
	MinioConfig
	AuthConfig

	Port        int      `env:"PORT" env-required:"true"`
	ApiVersion  int      `env:"API_VERSION" env-required:"true"`
	MaxFileSize int64    `env:"MAX_FILE_SIZE" env-required:"true"` // in bytes
	Env         string   `env:"ENV" env-required:"true"`
	Ssl         bool     `env:"SSL" env-default:"false"`
	Origins     []string `env:"ALLOW_ORIGINS" env-default:"*"`
}

type MinioConfig struct {
	User     string `env:"MINIO_USER" env-required:"true"`
	Password string `env:"MINIO_PASSWORD" env-required:"true"`
	Endpoint string `env:"MINIO_ENDPOINT" env-default:"localhost:9000"`
}

type PostgresConfig struct {
	User         string `env:"DB_USER" env-required:"true"`
	Password     string `env:"DB_PASSWORD" env-required:"true"`
	DbName       string `env:"DB_NAME" env-required:"true"`
	Host         string `env:"DB_HOST" env-default:"localhost"`
	Port         int    `env:"DB_PORT" env-default:"5432"`
	ConnAttempts int    `env:"CONNECTION_ATTEMPTS" env-default:"5"`
}

type AuthConfig struct {
	Password      string `env:"ADMIN_PASSWORD" env-required:"true"`
	Login         string `env:"ADMIN_LOGIN" env-required:"true"`
	Secret        string `env:"SECRET" env-required:"true"`
	TokenLifetime int    `env:"TOKEN_LIFETIME" env-required:"true"`
}

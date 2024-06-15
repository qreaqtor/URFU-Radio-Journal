package config

type ServerConfig struct {
	PostgresConfig `yaml:"postgres" env-required:"true"`
	MinioConfig    `yaml:"minio" env-required:"true"`
	AuthConfig     `yaml:"auth" env-required:"true"`
	Monitoring     `yaml:"monitoring" env-required:"true"`

	Env         string   `yaml:"env" env-required:"true"`
	Port        int      `yaml:"port" env-required:"true"`
	ApiVersion  int      `yaml:"api_version" env-required:"true"`
	MaxFileSize int64    `yaml:"max_file_size" env-default:"10485760"` // in bytes, 10 Mbyte
	Ssl         bool     `yaml:"ssl" env-default:"false"`
	Origins     []string `yaml:"allow_origins" env-default:"*"`
	Methods     []string `yaml:"allow_methods" env-default:"*"`
	Headers     []string `yaml:"allow_headers"`
}

type MinioConfig struct {
	User     string         `yaml:"user" env-required:"true"`
	Password string         `yaml:"password" env-required:"true"`
	Endpoint string         `yaml:"endpoint" env-default:"localhost:9000"`
	Buckets  []BucketConfig `yaml:"buckets" env-required:"true"`
}

type BucketConfig struct {
	Name         string   `yaml:"name" env-required:"true"`
	ContentTypes []string `yaml:"content_types" env-required:"true"`
}

type PostgresConfig struct {
	User         string `yaml:"user" env-required:"true"`
	Password     string `yaml:"password" env-required:"true"`
	Database     string `yaml:"database" env-required:"true"`
	Host         string `yaml:"host" env-default:"localhost"`
	Port         int    `yaml:"port" env-default:"5432"`
	ConnAttempts int    `yaml:"connection_attempts" env-default:"5"`
}

type AuthConfig struct {
	Password      string `yaml:"admin_password" env-required:"true"`
	Login         string `yaml:"admin_login" env-required:"true"`
	Secret        string `yaml:"secret" env-required:"true"`
	TokenLifetime int    `yaml:"token_lifetime" env-required:"true"`
}

type Monitoring struct {
	ContentTypes []string `yaml:"content_types" env-required:"true"`
}

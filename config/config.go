package config

// Config represents the minimal application settings required to boot the
// service. Each nested struct captures one concern (server, data stores,
// application behavior) so that future configuration values can extend these
// sections without breaking existing consumers.
type Config struct {
	Server   ServerConfig   `json:"server" yaml:"server"`
	Database DatabaseConfig `json:"database" yaml:"database"`
	Redis    RedisConfig    `json:"redis" yaml:"redis"`
	App      AppConfig      `json:"app" yaml:"app"`
}

// ServerConfig holds network-facing settings for the HTTP server.
type ServerConfig struct {
	Host         string `json:"host" yaml:"host"`
	Port         string `json:"port" yaml:"port"`
	ReadTimeout  int    `json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout" yaml:"writeTimeout"`
}

// DatabaseConfig captures the primary relational database connection details.
// Either MySQLDSN can be used directly or the discrete fields can be populated
// by env-specific loaders.
type DatabaseConfig struct {
	MySQLDSN string `json:"mysqlDSN" yaml:"mysqlDSN"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Name     string `json:"name" yaml:"name"`
}

// RedisConfig includes the minimum values needed to connect to Redis.
type RedisConfig struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}

// AppConfig defines per-environment application behavior such as log levels.
type AppConfig struct {
	Env      Environment `json:"env" yaml:"env"`
	LogLevel string      `json:"logLevel" yaml:"logLevel"`
}

// Environment enumerates supported runtime environments.
type Environment string

// Supported environment values.
const (
	EnvDevelopment Environment = "dev"
	EnvProduction  Environment = "prod"
)

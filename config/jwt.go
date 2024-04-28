package config

type Jwt struct {
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
}

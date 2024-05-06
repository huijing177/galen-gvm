package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// 验证码
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `json:"pgsql" yaml:"pgsql"`
}

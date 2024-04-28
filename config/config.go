package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `json:"pgsql" yaml:"pgsql"`
}

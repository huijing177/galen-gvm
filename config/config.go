package config

type Server struct {
	System System `json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `json:"pgsql" yaml:"pgsql"`
}

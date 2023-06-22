package config

type MQ struct {
	Server string `mapstructure:"server" json:"server" yaml:"server"` // 级别
	Port   string `mapstructure:"port" json:"port" yaml:"port"`       // 日志前缀
}

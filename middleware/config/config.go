package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	MQ    MQ    `mapstructure:"rocket-mq" json:"rocket-mq" yaml:"rocket-mq"`
}

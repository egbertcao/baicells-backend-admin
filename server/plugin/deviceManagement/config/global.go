package config

type Config struct {
	Mongo  Mongo    `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Rabbit RabbitMq `mapstructure:"rabbit" json:"rabbit" yaml:"rabbit"`
	Mqtt   Mqtt     `mapstructure:"mqtt" json:"mqtt" yaml:"mqtt"`
}

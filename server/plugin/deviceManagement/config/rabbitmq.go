package config

type RabbitMq struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Qos      string `mapstructure:"qos" json:"qos" yaml:"qos"`                // 数据库密码
}

func RabbitMqDns(r *RabbitMq) string {
	return "amqp://" + r.Username + ":" + r.Password + "@" + r.Path
}

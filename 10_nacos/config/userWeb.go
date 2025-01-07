package config

type ServerConfig struct {
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTConfig     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	ConsulConfig  ConsulConfig  `mapstructure:"consul" json:"consul"`
	Name          string        `mapstructure:"name" json:"name"`
	Port          int           `mapstructure:"port" json:"port"`
}

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"signing_key" json:"signing_key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

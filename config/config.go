package config

import (
	"github.com/spf13/viper"
)

type allConfig struct {
	MySQL  *MysqlConfig
	Redis  *RedisConfig
	Server *ServerConfig
	Log    *LogConfig
	Mail   *MailConfig
}

// MysqlConfig mysql数据库配置
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Auth string `mapstructure:"auth"`
	Db   int    `mapstructure:"db"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type LogConfig struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type MailConfig struct {
	Port     int64  `json:"port" mapstructure:"port" yaml:"port"`
	From     string `json:"from" mapstructure:"from" yaml:"from"`
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Secret   string `json:"secret" mapstructure:"secret" yaml:"secret"`
	Nickname string `json:"nickname" mapstructure:"nickname" yaml:"nickname"`
}

var Config = new(allConfig)

func InitConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置错误-1")
	}

	cErr := viper.Unmarshal(Config)
	if cErr != nil {
		panic(cErr.Error())
	}
}

package infrastructure

import (
	"github.com/golang-graphql/exception"
	"github.com/spf13/viper"
)

type Config struct {
	MongoPoolMin           int    `mapstructure:"MONGO_POOL_MIN"`
	MongoPoolMax           int    `mapstructure:"MONGO_POOL_MAX"`
	MongoMaxIdleTimeSecond int    `mapstructure:"MONGO_MAX_IDLE_TIME_SECOND"`
	MongoURI               string `mapstructure:"MONGO_URI"`
	MongoDatabae           string `mapstructure:"MONGO_DATABASE"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)
	err = viper.Unmarshal(&config)
	exception.PanicIfNeeded(err)
	return config, err
}

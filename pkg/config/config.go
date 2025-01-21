package config

import (
	"log"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigName("pocone")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}

	return v
}

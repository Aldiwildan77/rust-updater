package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/Aldiwildan77/rust-notifier-api/library/files"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Port int `mapstructure:"PORT"`
}

func init() {
	loadEnvVars()
}

func loadEnvVars() {
	os.Setenv("TZ", "Asia/Jakarta")

	cdir, _ := os.Getwd()
	ef := fmt.Sprintf("%s/%s", cdir, ".env")

	var err error
	if files.FileExists(ef) {
		err = loadConfigFile(ef)
	} else {
		err = loadConfigEnvVar()
	}

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err.Error()))
	}
}

func loadConfigFile(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Cfg)
	return
}

func loadConfigEnvVar() (err error) {
	viper.AutomaticEnv()

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &Cfg,
	}

	n := make(map[string]interface{})
	for _, e := range os.Environ() {
		data := strings.Split(e, "=")
		n[data[0]] = data[1]
	}

	decoder, _ := mapstructure.NewDecoder(config)
	decoder.Decode(n)

	return
}

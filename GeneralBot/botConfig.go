package GeneralBot

import (
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	Telegram_Token string
	Admin_Id       int64
	Bot_Debug      bool
	Bot_Helper     string
}

func LoadConfig(path string) (cfg Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	r := reflect.TypeOf(cfg)

	for j := 0; j < r.NumField(); j++ {
		f := r.Field(j)

		if err = viper.BindEnv(f.Name); err != nil {
			return
		}
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}

	return
}

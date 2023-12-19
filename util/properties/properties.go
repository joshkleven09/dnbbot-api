package properties

import (
	"fmt"
	"github.com/spf13/viper"
)

func New() {
	viper.SetConfigName("properties")
	viper.AddConfigPath("/dnbbotapi/properties")
	viper.AddConfigPath("./properties")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error properties file: %w", err))
	}
}

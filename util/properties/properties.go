package properties

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func New(env string, log *zerolog.Logger) {

	if env == "prod" {
		log.Info().Msg("Getting prod props...")
		viper.SetConfigName("properties-prod")
	} else {
		log.Info().Msg("Getting local props...")
		viper.SetConfigName("properties-local")
	}

	viper.AddConfigPath("/dnbbotapi/properties")
	viper.AddConfigPath("./properties")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error properties file: %w", err))
	}
}

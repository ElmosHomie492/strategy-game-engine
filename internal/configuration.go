package internal

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	viper "github.com/spf13/viper"
)

type EngineConfiguration struct {
	Logger zerolog.Logger `mapstructure:"logging"`
}

func BuildConfiguration(viperRef *viper.Viper, config string, envVars ...string) EngineConfiguration {
	var engineConfig EngineConfiguration

	viperRef.SetConfigName(config)
	viperRef.SetConfigType("yaml")
	viperRef.AddConfigPath(".")
	err := viperRef.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.GetViper().Unmarshal(&engineConfig)

	return engineConfig
}

func (config *EngineConfiguration) SetupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	config.Logger = log.With().Caller().Logger()
}

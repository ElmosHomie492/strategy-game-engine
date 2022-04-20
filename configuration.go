package main

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	viper "github.com/spf13/viper"
)

type EngineConfiguration struct {
	FirstValue  string `map:"firstValue"`
	SecondValue string `map:"secondValue"`
	Logger      zerolog.Logger
}

func initEngine(config string) EngineConfiguration {
	var engineConfig EngineConfiguration

	viper.SetConfigName(config)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.GetViper().Unmarshal(&engineConfig)

	engineConfig.Logger = setupLogging()

	engineConfig.Logger.Info().Msg("Starting Engine...")

	return engineConfig
}

func setupLogging() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.With().Caller().Logger()

	return log.Logger
}

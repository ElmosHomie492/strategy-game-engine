package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	viper "github.com/spf13/viper"
)

type EngineConfiguration struct {
	Logger    zerolog.Logger
	DebugMode bool `mapstructure:"debugMode"`
	Game      *Game
}

func BuildConfiguration(viperRef *viper.Viper, config string, envVars ...string) EngineConfiguration {
	var engineConfig EngineConfiguration

	// Configuration
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

func (config *EngineConfiguration) SetupLogging(debugMode bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", debugMode, "sets log level to debug")

	flag.Parse()

	// Default level is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	config.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

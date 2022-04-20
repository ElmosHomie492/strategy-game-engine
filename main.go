package main

import (
	"github.com/ElmosHomie492/strategy-game-engine/internal"
)

func main() {
	var config internal.EngineConfiguration

	config.Logger = config.SetupLogging()

	config.Logger.Info().Msg("Starting Engine...")
}

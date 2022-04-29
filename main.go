package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func configureGame() *Game {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Super Elmo's World")

	return game
}

func main() {
	var config EngineConfiguration
	config.SetupLogging(false)

	config.Logger.Info().Msg("Starting Engine...")
	config.Game = configureGame()

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(config.Game); err != nil {
		config.Logger.Fatal().Err(err)
	}
}

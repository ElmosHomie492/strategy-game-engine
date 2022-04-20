package main

func main() {
	config := initEngine("configuration")

	config.Logger.Info().Msg(config.FirstValue)
}

package main

func main() {
	config := initEngine("configuration")
	
	var config internal.Configuration

	setup.BuildConfiguration(
		viper.New(),
		&config,
		*internal.ConfigurationDefaults,
	)

	config.Logger.Info().Msg(config.FirstValue)
}

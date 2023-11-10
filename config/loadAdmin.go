package config

import (
	"log"
	"github.com/spf13/viper"
)

func LoadAdmin() {

	viper.AddConfigPath("./config/")
	viper.SetConfigName("defaultConfig") // Register config file name (no extension)
	viper.SetConfigType("json")          // Look for specific type

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	// Get the array data from the configuration
	jsonArray := viper.GetStringSlice("adminEmails")

	// Insert data into the map
	for _, item := range jsonArray {
		AdminEmails[item] = true
	}

	logger.Info().Msg("📢 Info message :" + "Admin data loaded")
}

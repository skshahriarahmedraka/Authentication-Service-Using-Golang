package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadAdmin() {
	// Set the configuration file name (data.json in this case)
	// viper.SetConfigFile("/config/defaultConfig.json")

	// viper.AddConfigPath("$GOPATH/src/github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/config/")
	// 	viper.AddConfigPath("./config/")
	// 	viper.SetConfigName("defaultConfig")
	viper.AddConfigPath("./config/")
    viper.SetConfigName("defaultConfig") // Register config file name (no extension)
    viper.SetConfigType("json")   // Look for specific type
    // viper.ReadInConfig()
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	// Get the array data from the configuration
	jsonArray := viper.GetStringSlice("adminEmails")
    fmt.Println("🚀 ~ file: loadAdmin.go ~ line 28 ~ funcLoadAdmin ~ jsonArray : ", jsonArray)

	// Create a map to store the data
	// dataMap := make(map[string]bool)

	// Insert data into the map
	for _, item := range jsonArray {
		AdminEmails[item] = true
	}

	// Print the resulting map
	fmt.Println("Data Map:", AdminEmails)
}

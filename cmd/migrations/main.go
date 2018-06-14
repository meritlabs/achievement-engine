package main

import (
	"fmt"

	"github.com/meritlabs/achievement-engine/pkg/db"
	"github.com/meritlabs/achievement-engine/pkg/migrations/data"
	"github.com/spf13/viper"
)

func initializeConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // lookup the config in the working dir
	viper.SetConfigType("yaml")

	// set explicit app defaults
	viper.SetDefault("db.connectionString", "localhost")

	viper.SetDefault("blockchain.network", "testnet")
	viper.SetDefault("blockchain.rpc.host", "localhost")
	viper.SetDefault("blockchain.rpc.user", "merit")
	viper.SetDefault("blockchain.rpc.password", "local321")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}
}

func main() {
	initializeConfig()
	session, err := db.WithDBSession(viper.GetString("db.connectionString"))

	if err != nil {
		panic(err)
	}
	defer session.Close()

	data.UpdateGoals(session)
}

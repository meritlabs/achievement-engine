package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/api/controllers"
	"github.com/meritlabs/achievement-engine/api/middleware"
	"github.com/meritlabs/achievement-engine/api/services"
	"github.com/meritlabs/achievement-engine/db/stores"
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

func createStores() *stores.Store {
	conStr := viper.GetString("db.connectionString")
	return stores.InitStore(conStr)
}

func createUserService(store *stores.Store, logger *log.Logger) services.UsersService {
	Net := chaincfg.Params{
		Name:             viper.GetString("blockchain.network"),
		PubKeyHashAddrID: 110,
		ScriptHashAddrID: 125,
	}

	BCClient := services.NewClient(
		viper.GetString("blockchain.rpc.host"),
		viper.GetString("blockchain.rpc.user"),
		viper.GetString("blockchain.rpc.password"),
		logger,
	)

	usersService := services.UsersService{Net, BCClient, store, store, store, store, store}
	return usersService
}

func main() {
	initializeConfig()

	store := createStores()
	defer store.ShutDownStore()

	logger := log.New(os.Stderr, "", log.LstdFlags)

	userService := createUserService(store, logger)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JSONErrorReporter())

	apiGroup := router.Group("/achivement-engine/api/v1")
	{
		apiGroup.POST("/sessions", controllers.CreateSession(userService))

		goals := apiGroup.Group("/goals")
		{
			goals.GET("/", controllers.ListGoals(store))
		}
		achievements := apiGroup.Group("/achievements", middleware.Auth(store, store))
		{
			achievements.GET("/", controllers.ListAchievements(store))
			achievements.GET("/:slug", controllers.GetAchievement(store))
			achievements.POST("/:slug/step/:step/complete", controllers.UpdateAchievement(store))
		}
		settings := apiGroup.Group("/settings", middleware.Auth(store, store))
		{
			settings.GET("/", controllers.GetSettings(store))
			settings.POST("/", controllers.UpdateSettings(store))
			settings.PUT("/", controllers.UpdateSettings(store))
		}
	}

	router.Run(viper.GetString("port"))
}

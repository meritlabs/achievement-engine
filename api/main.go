package main

import (
	"log"
	"os"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/api/controllers"
	"github.com/meritlabs/achievement-engine/api/middleware"
	"github.com/meritlabs/achievement-engine/api/services"
	"github.com/meritlabs/achievement-engine/db/stores"
)

func createStores() *stores.Store {
	return stores.InitStore()
}

func createUserService(store *stores.Store, logger *log.Logger) services.UsersService {
	Net := chaincfg.Params{
		Name:             "testnet",
		PubKeyHashAddrID: 110,
		ScriptHashAddrID: 125,
	}

	BCClient := services.NewClient(
		"localhost", //viper.GetString("blockchain.rpc.host"),
		"merit",     //viper.GetString("blockchain.rpc.user"),
		"local321",  //viper.GetString("blockchain.rpc.password"),
		logger,
	)

	usersService := services.UsersService{Net, BCClient, store, store, store, store}
	return usersService
}

func main() {
	store := createStores()
	defer store.ShutDownStore()

	logger := log.New(os.Stderr, "", log.LstdFlags)

	userService := createUserService(store, logger)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JSONErrorReporter())

	apiGroup := router.Group("/achivement-engine/api/v1")
	{
		goals := apiGroup.Group("/goals")
		{
			goals.GET("/", controllers.ListGoals(store))
		}
		achievements := apiGroup.Group("/achievements", middleware.Auth(userService, logger))
		{
			achievements.GET("/", controllers.ListAchievements(store))
			achievements.GET("/:slug", controllers.GetAchievement(store))
			achievements.POST("/:slug", controllers.UpdateAchievement(store))
		}
	}

	router.Run()
}

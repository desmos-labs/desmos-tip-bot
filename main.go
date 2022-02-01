package main

import (
	"fmt"
	"github.com/desmos-labs/desmostipbot/apis/streamlabs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/desmos-labs/desmostipbot/database"

	"github.com/desmos-labs/desmostipbot/cosmos"
	"github.com/desmos-labs/desmostipbot/tipper"
	"github.com/desmos-labs/desmostipbot/twitter"
	"github.com/desmos-labs/desmostipbot/types"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Please specify the config.toml file path")
	}

	configPath := os.Args[1]
	cfg, err := types.ParseConfig(configPath)
	if err != nil {
		panic(err)
	}

	// Build the database
	db, err := database.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	// Build the Cosmos client
	cosmosClient, err := cosmos.NewClient(cfg.Chain)
	if err != nil {
		panic(err)
	}

	// Build the tipper instance
	tipperClient := tipper.NewTipper(cosmosClient, db)

	// Setup the Twitter client
	twitterClient := twitter.NewClient(cfg.Twitter, tipperClient)
	err = twitterClient.StartListening()
	if err != nil {
		panic(err)
	}

	// Setup the rest server
	r := gin.Default()
	r.Use(gin.Recovery()) // Set panic errors to be 500

	// Setup CORS
	ginCfg := cors.DefaultConfig()
	ginCfg.AllowAllOrigins = true
	r.Use(cors.New(ginCfg))

	// Register the handlers
	streamlabs.RegisterHandlers(r, streamlabs.NewHandler(cfg.APIs.Streamlabs, db))

	// Run the server
	//port := cfg.API.Port
	//if port == 0 {
	//	port = 8080
	//}
	r.Run(fmt.Sprintf(":%d", 5000))

	// Listen for os signals and stop the clients
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	twitterClient.Stop()
}

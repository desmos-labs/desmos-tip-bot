package main

import (
	"github.com/desmos-labs/desmostipbot/cosmos"
	"github.com/desmos-labs/desmostipbot/tipper"
	"github.com/desmos-labs/desmostipbot/twitter"
	"github.com/desmos-labs/desmostipbot/types"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// Build the Cosmos client
	cosmosClient, err := cosmos.NewClient(cfg.Chain)
	if err != nil {
		panic(err)
	}

	// Build the tipper instance
	tipperClient := tipper.NewTipper(cosmosClient)

	// Setup the Twitter client
	twitterClient := twitter.NewClient(cfg.Twitter, tipperClient)
	err = twitterClient.StartListening()
	if err != nil {
		panic(err)
	}

	// Listen for os signals and stop the clients
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	twitterClient.Stop()

}

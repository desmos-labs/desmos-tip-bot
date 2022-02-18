package main

import (
	"fmt"
	"github.com/desmos-labs/plutus/apis/config"
	"github.com/desmos-labs/plutus/apis/user"
	"github.com/desmos-labs/plutus/coingecko"
	"log"
	"os"
	"os/signal"
	"syscall"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmosclient "github.com/desmos-labs/cosmos-go-wallet/client"
	cosmoswallet "github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/desmos-labs/plutus/apis/donations"
	"github.com/desmos-labs/plutus/apis/oauth"
	streamlabsclient "github.com/desmos-labs/plutus/integrations/streamlabs"
	"github.com/desmos-labs/plutus/integrations/twitter"
	notificationshandler "github.com/desmos-labs/plutus/notifications/handler"
	oauthhandler "github.com/desmos-labs/plutus/oauth/handler"

	"github.com/desmos-labs/plutus/database"

	desmosapp "github.com/desmos-labs/desmos/v2/app"

	"github.com/desmos-labs/plutus/desmos"
	"github.com/desmos-labs/plutus/types"
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

	// Setup the configurations
	cdc, amino := desmosapp.MakeCodecs()
	desmosapp.SetupConfig(sdk.GetConfig())

	// Build the various clients
	wallet, err := buildWallet(cfg)
	if err != nil {
		panic(err)
	}

	// Build the Desmos client
	coingeckoClient := coingecko.NewClient(cfg.Donations)
	desmosClient := desmos.NewDesmosClient(cfg.Chain.DesmosClientConfig, wallet, db)
	streamlabsClient := streamlabsclient.NewClient(cfg.Integrations.Streamlabs, coingeckoClient, db)
	twitterClient := twitter.NewClient(cfg.Integrations.Twitter)

	oAuthHandler := oauthhandler.NewHandler().
		RegisterClient(streamlabsClient)

	notificationsHandler := notificationshandler.NewHandler().
		RegisterClient(streamlabsClient).
		RegisterClient(twitterClient)

	// Setup the rest server
	r := gin.Default()
	r.Use(gin.Recovery()) // Set panic errors to be 500

	// Setup CORS
	ginCfg := cors.DefaultConfig()
	ginCfg.AllowAllOrigins = true
	r.Use(cors.New(ginCfg))

	// Register the handlers
	config.RegisterHandlers(r, config.NewHandler(desmosClient))
	donations.RegisterHandlers(r, donations.NewHandler(desmosClient, notificationsHandler, db))
	oauth.RegisterHandlers(r, oauth.NewHandler(cfg.Donations, oAuthHandler, cdc, amino, db))
	user.RegisterHandlers(r, user.NewHandler(desmosClient, cdc, db))

	// Run the server
	port := cfg.APIs.Port
	if port == 0 {
		port = 8080
	}
	_ = r.Run(fmt.Sprintf(":%d", port))

	// Setup the Twitter client
	if cfg.Integrations.Twitter.Enabled {
		err = twitterClient.StartListening()
		if err != nil {
			panic(err)
		}
	}

	// Listen for os signals and stop the clients
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	twitterClient.Stop()
}

func buildWallet(cfg *types.Config) (*cosmoswallet.Wallet, error) {
	encodingConfig := desmosapp.MakeTestEncodingConfig()

	// Build the Cosmos client
	cosmosClient, err := cosmosclient.NewClient(cfg.Chain.ChainConfig, encodingConfig.Marshaler)
	if err != nil {
		return nil, err
	}

	// Build the Cosmos wallet
	return cosmoswallet.NewWallet(cfg.Account, cosmosClient, encodingConfig.TxConfig)
}

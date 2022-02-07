package desmos

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/authz"
	cosmoswallet "github.com/desmos-labs/cosmos-go-wallet/wallet"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/hasura/go-graphql-client"

	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/types"
)

// Client represents the client used to interact with the Desmos chain
type Client struct {
	database  *database.Database
	cdc       codec.Codec
	gqlClient *graphql.Client

	cosmosClient   *cosmoswallet.Wallet
	authzClient    authz.QueryClient
	profilesClient profilestypes.QueryClient
	txClient       tx.ServiceClient
}

// NewDesmosClient allows to build a new Client instance
func NewDesmosClient(cfg *types.DesmosClientConfig, cosmosClient *cosmoswallet.Wallet, database *database.Database) *Client {
	return &Client{
		database:       database,
		cosmosClient:   cosmosClient,
		gqlClient:      graphql.NewClient(cfg.GraphQLAddr, nil),
		authzClient:    authz.NewQueryClient(cosmosClient.Client.GRPCConn),
		profilesClient: profilestypes.NewQueryClient(cosmosClient.Client.GRPCConn),
		txClient:       tx.NewServiceClient(cosmosClient.Client.GRPCConn),
	}
}

// ParseAddress parses the given address as a sdk.AccAddress instance
func (client *Client) ParseAddress(address string) (sdk.AccAddress, error) {
	return client.cosmosClient.Client.ParseAddress(address)
}

package desmos

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	cosmoswallet "github.com/desmos-labs/cosmos-go-wallet/wallet"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/hasura/go-graphql-client"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/types"
)

// Client represents the client used to interact with the Desmos chain
type Client struct {
	database       *database.Database
	cosmosClient   *cosmoswallet.Wallet
	gqlClient      *graphql.Client
	authzClient    authz.QueryClient
	profilesClient profilestypes.QueryClient
}

// NewDesmosClient allows to build a new Client instance
func NewDesmosClient(cfg *types.DesmosClientConfig, cosmosClient *cosmoswallet.Wallet, database *database.Database) *Client {
	return &Client{
		database:       database,
		cosmosClient:   cosmosClient,
		gqlClient:      graphql.NewClient(cfg.GraphQLAddr, nil),
		authzClient:    authz.NewQueryClient(cosmosClient.Client.GRPCConn),
		profilesClient: profilestypes.NewQueryClient(cosmosClient.Client.GRPCConn),
	}
}

// ParseAddress parses the given address as a sdk.AccAddress instance
func (client *Client) ParseAddress(address string) (sdk.AccAddress, error) {
	return client.cosmosClient.Client.ParseAddress(address)
}

// SearchDesmosAddress searches the Desmos address given a specific application and username.
// If the application is "desmos", then the provided username will be treated as a DTag.
// Otherwise, the GraphQL client will be used to search for an application link with the specified
// application and username case-insensitive.
func (client *Client) SearchDesmosAddress(application, username string) (sdk.AccAddress, error) {
	if strings.EqualFold(application, "desmos") {
		return client.getDesmosAddressFromDTag(username)
	}
	return client.getDesmosAddressFromApplication(application, username)
}

// getDesmosAddressFromDTag returns the Desmos address of the user that has the given DTag
func (client *Client) getDesmosAddressFromDTag(dTag string) (sdk.AccAddress, error) {
	res, err := client.profilesClient.Profile(context.Background(), &profilestypes.QueryProfileRequest{User: dTag})
	if err != nil {
		return nil, err
	}

	if res.Profile == nil {
		return nil, apiutils.WrapErr(http.StatusNotFound, fmt.Sprintf("User with DTag %s not found", dTag))
	}

	var account authtypes.AccountI
	err = client.cosmosClient.Client.Codec.UnpackAny(res.Profile, &account)
	if err != nil {
		return nil, err
	}

	return account.GetAddress(), nil
}

// profileByAppLinkQuery represents the GraphQL query used to search for a Desmos profile
// connected to a specific application account
type profileByAppLinkQuery struct {
	Links []struct {
		UserAddress string `graphql:"user_address"`
	} `graphql:"application_link(where: {application: {_ilike: $application}, username: {_ilike: $username}})"`
}

// getDesmosAddressFromApplication returns the Desmos address associated with the given application and username
func (client *Client) getDesmosAddressFromApplication(application, username string) (sdk.AccAddress, error) {
	var query profileByAppLinkQuery
	variables := map[string]interface{}{
		"application": graphql.String(application),
		"username":    graphql.String(username),
	}

	err := client.gqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}

	if len(query.Links) == 0 {
		return nil, apiutils.WrapErr(http.StatusNotFound, fmt.Sprintf(
			"No Desmos address connected to %s %s found", application, username))
	}

	return client.ParseAddress(query.Links[0].UserAddress)
}

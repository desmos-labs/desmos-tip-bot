package tipper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmostipbot/cosmos"
)

// Tipper represents the client used to tip the various accounts
type Tipper struct {
	cosmosClient *cosmos.Client
}

// NewTipper allows to build a new Tipper instance
func NewTipper(cosmosClient *cosmos.Client) *Tipper {
	return &Tipper{
		cosmosClient: cosmosClient,
	}
}

// Tip allows to tip the given user of the given amount.
// Returns the tip transaction hash if everything is correct or an error otherwise.
func (tipper *Tipper) Tip(amount sdk.Coins, user string) (txHash string, err error) {
	// TODO: Do this once the authz and feegrant module are enabled
	return "", nil
}

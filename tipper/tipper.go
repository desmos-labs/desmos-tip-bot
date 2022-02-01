package tipper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmostipbot/cosmos"
	"github.com/desmos-labs/desmostipbot/database"
	"github.com/desmos-labs/desmostipbot/types"
)

// Tipper represents the client used to tip the various accounts
type Tipper struct {
	cosmosClient *cosmos.Client
	database     *database.Database
}

// NewTipper allows to build a new Tipper instance
func NewTipper(cosmosClient *cosmos.Client, database *database.Database) *Tipper {
	return &Tipper{
		cosmosClient: cosmosClient,
		database:     database,
	}
}

// Tip allows to tip the given user of the given amount.
// Returns the tip transaction hash if everything is correct or an error otherwise.
func (tipper *Tipper) Tip(sender *types.AppAccount, receiver *types.AppAccount, amount sdk.Coins) (txHash string, err error) {
	//account, err := tipper.database.GetChainAccountByAppAccount(sender)
	//if err != nil {
	//	return "", err
	//}
	//
	//senderAddr, err := account.GetAccAddress()
	//if err != nil {
	//	return "", err
	//}
	//
	//// TODO: Change this to send the funds using a MsgExec from the authz module
	//// TODO: Allow to send funds to a user that does not have an account yet
	//msg := banktypes.NewMsgSend(senderAddr, nil, amount)
	//res, err := tipper.cosmosClient.BroadcastTx(msg)
	//if err != nil {
	//	return "", err
	//}
	//
	//return res.TxHash, nil

	return "", nil
}

package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmostipbot/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

// Client represents a Cosmos client that should be used to create and send transactions to the chain
type Client struct {
	cliCtx client.Context

	privKey cryptotypes.PrivKey

	fees string
}

// NewClient allows to build a new Client instance
func NewClient(chainCfg *types.ChainConfig) (*Client, error) {
	// Build the config
	sdkCfg := sdk.GetConfig()
	app.SetupConfig(sdkCfg)
	sdkCfg.Seal()

	// Get the private types
	algo := hd.Secp256k1
	derivedPriv, err := algo.Derive()(chainCfg.Mnemonic, "", sdkCfg.GetFullFundraiserPath())
	if err != nil {
		return nil, err
	}
	privKey := algo.Generate()(derivedPriv)

	// Build the RPC client
	rpcClient, err := rpchttp.New(chainCfg.NodeURI, "/websocket")
	if err != nil {
		return nil, err
	}

	// Build the context
	encodingConfig := app.MakeTestEncodingConfig()
	cliCtx := client.Context{}.
		WithNodeURI(chainCfg.NodeURI).
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastSync).
		WithClient(rpcClient).
		WithChainID(chainCfg.ChainID).
		WithFromAddress(sdk.AccAddress(privKey.PubKey().Address()))

	return &Client{
		cliCtx:  cliCtx,
		privKey: privKey,
		fees:    chainCfg.Fees,
	}, nil
}

// GetAccAddress returns the address of the account that is going to be used to sign the transactions
func (client *Client) GetAccAddress() sdk.AccAddress {
	return sdk.AccAddress(client.privKey.PubKey().Address())
}

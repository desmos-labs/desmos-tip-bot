package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ChainAccount represents the details of a chain account
type ChainAccount struct {
	Address string
}

// NewChainAccount allows to create a new ChainAccount instance
func NewChainAccount(address string) *ChainAccount {
	return &ChainAccount{
		Address: address,
	}
}

// GetAccAddress returns the AccAddress for this chain account
func (account *ChainAccount) GetAccAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(account.Address)
}

// --------------------------------------------------------------------------------------------------------------------

const (
	AppTwitter = "twitter"
)

// AppAccount contains the data about an application account
type AppAccount struct {
	AppName  string
	Username string
}

func NewAppAccount(appName, username string) *AppAccount {
	return &AppAccount{
		AppName:  appName,
		Username: username,
	}
}

package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// User contains the data about an application user
type User struct {
	DesmosAddress string
}

func NewUser(desmosAddress string) *User {
	return &User{
		DesmosAddress: desmosAddress,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type ServiceAccount struct {
	Service      string
	AccessToken  string
	RefreshToken string
}

func NewServiceAccount(service, accessToken, refreshToken string) *ServiceAccount {
	return &ServiceAccount{
		Service:      service,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type ApplicationAccount struct {
	Application string
	Username    string
}

func NewApplicationAccount(application, username string) *ApplicationAccount {
	return &ApplicationAccount{
		Application: application,
		Username:    username,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type DonationTx struct {
	Hash             string
	SenderAddress    string
	RecipientAddress string
	Amount           sdk.Coins
}

func NewDonationTx(hash, senderAddress, recipientAddress string, amount sdk.Coins) *DonationTx {
	return &DonationTx{
		Hash:             hash,
		SenderAddress:    senderAddress,
		RecipientAddress: recipientAddress,
		Amount:           amount,
	}
}

package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

type UserPreferences struct {
	Currency string
}

func NewUserPreferences(currency string) *UserPreferences {
	return &UserPreferences{
		Currency: currency,
	}
}

func DefaultPreferences() *UserPreferences {
	return NewUserPreferences(
		"usd",
	)
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
	Timestamp        time.Time
}

func NewDonationTx(hash, senderAddress, recipientAddress string, amount sdk.Coins, timestamp time.Time) *DonationTx {
	return &DonationTx{
		Hash:             hash,
		SenderAddress:    senderAddress,
		RecipientAddress: recipientAddress,
		Amount:           amount,
		Timestamp:        timestamp,
	}
}

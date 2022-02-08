package types

import (
	"strings"

	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

type DonationAmount struct {
	Currency string
	Amount   float64
}

func NewDonationAmount(amount float64, currency string) *DonationAmount {
	return &DonationAmount{
		Currency: currency,
		Amount:   amount,
	}
}

// Notification contains all the data that might be needed in order to send a notification
type Notification struct {
	// RecipientAppAccount contains the data of the application on which the donation was made (Twitter, Twitch, etc)
	RecipientAppAccount *ApplicationAccount

	// Tx contains the details of the donation transaction
	Tx *DonationTx

	// TipperProfile represents the profile of the user that made the donation
	TipperProfile *profilestypes.Profile

	// TipperUsername contains the username the tipper has chosen to use for this donation
	TipperUsername string

	// Donation message
	Message string
}

// NewNotification returns a new Notification instance
func NewNotification(
	recipientAppAccount *ApplicationAccount,
	donationTx *DonationTx,
	tipperProfile *profilestypes.Profile,
	tipperUsername,
	message string,
) *Notification {
	return &Notification{
		RecipientAppAccount: recipientAppAccount,
		Tx:                  donationTx,
		TipperProfile:       tipperProfile,
		TipperUsername:      strings.TrimSpace(tipperUsername),
		Message:             strings.TrimSpace(message),
	}
}

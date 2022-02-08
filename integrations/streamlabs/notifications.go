package streamlabs

import (
	"fmt"
	"net/http"

	"github.com/desmos-labs/plutus/notifications"
	"github.com/desmos-labs/plutus/types"
)

var _ notifications.Client = &Client{}

const (
	ServiceStreamlabs = "streamlabs"
)

// App implements notifications.Client
func (client *Client) App() string {
	return types.AppTwitch
}

// SendNotification implements notifications.Client
func (client *Client) SendNotification(data *types.Notification) error {
	// Get the authorization token
	serviceAccount, err := client.db.GetServiceAccount(ServiceStreamlabs, data.RecipientAppAccount)
	if err != nil {
		return err
	}

	// If the token does not exist, simply return
	if serviceAccount == nil {
		return nil
	}

	// Refresh the token
	serviceAccount, err = client.RefreshAccount(serviceAccount)
	if err != nil {
		return err
	}

	// Get the preferences
	preferences, err := client.db.GetUserPreferences(data.Tx.RecipientAddress)
	if err != nil {
		return err
	}

	// Get the donation amount
	amount, err := client.coingecko.GetDonationAmount(data.Tx.Amount, preferences)
	if err != nil {
		return err
	}

	// Send the alert notification
	return client.sendDonation(data, amount, serviceAccount)
}

// sendDonation sends the donation notification with the given data using the provided service account
func (client *Client) sendDonation(
	data *types.Notification, amount *types.DonationAmount, serviceAccount *types.ServiceAccount,
) error {
	var tipperName string
	switch {
	case data.TipperUsername != "":
		tipperName = data.TipperUsername
	case data.TipperProfile != nil && data.TipperProfile.Nickname != "":
		tipperName = data.TipperProfile.Nickname
	case data.TipperProfile != nil:
		tipperName = data.TipperProfile.DTag
	default:
		tipperName = "Anonymous"
	}

	req, err := http.NewRequest("POST", "https://streamlabs.com/api/v1.0/user", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("name", tipperName)
	q.Add("message", data.Message)
	q.Add("identifier", data.Tx.SenderAddress)
	q.Add("amount", fmt.Sprintf("%.2f", amount.Amount))
	q.Add("currency", amount.Currency)
	q.Add("created_at", fmt.Sprintf("%d", data.Tx.Timestamp.Unix()))
	q.Add("access_token", serviceAccount.AccessToken)
	req.URL.RawQuery = q.Encode()

	res, err := client.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return client.getBadRequestError(res)
}

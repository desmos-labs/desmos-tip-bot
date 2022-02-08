package coingecko

import (
	"encoding/json"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/plutus/types"
)

type Client struct {
	cfg  *types.DonationsConfig
	http *http.Client
}

func NewClient(cfg *types.DonationsConfig) *Client {
	return &Client{
		cfg:  cfg,
		http: &http.Client{},
	}
}

type marketResponse struct {
	CurrentPrice float64 `json:"current_price"`
}

// GetDonationAmount returns the DonationAmount related to the given sdk.Coins amount, using the specified currency
func (c *Client) GetDonationAmount(amount sdk.Coins, preferences *types.UserPreferences) (*types.DonationAmount, error) {
	var cost float64
	for _, coin := range amount {
		coinCost, err := c.getAmountCost(coin, preferences.Currency)
		if err != nil {
			return nil, err
		}
		cost += coinCost
	}
	return types.NewDonationAmount(cost, preferences.Currency), nil
}

// getAmountCost returns the cost of the specified amount, in the given currency
func (c *Client) getAmountCost(amount sdk.Coin, currency string) (float64, error) {
	data, ok := c.cfg.SupportedDenoms[amount.Denom]
	if !ok {
		return 0, nil
	}

	req, err := http.NewRequest("POST", "https://api.coingecko.com/api/v3/coins/markets", nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("vs_currency", currency)
	q.Add("ids", data.CoinGeckoID)
	req.URL.RawQuery = q.Encode()

	res, err := c.http.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var body []marketResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return 0, err
	}

	if len(body) == 0 {
		return 0, nil
	}

	return body[0].CurrentPrice * float64(amount.Amount.Uint64()), nil
}

package desmos

import (
	"context"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

// GetGrants returns the grants made from the given granter to the provided grantee, if any
func (client *Client) GetGrants(granterAddress string, granteeAddress string) ([]*authz.Grant, error) {
	res, err := client.authzClient.Grants(context.Background(), &authz.QueryGrantsRequest{
		Grantee: granteeAddress,
		Granter: granterAddress,
	})
	if err != nil {
		return nil, err
	}

	return res.Grants, nil
}

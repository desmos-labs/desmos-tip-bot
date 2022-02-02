package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmoswallettypes "github.com/desmos-labs/cosmos-go-wallet/types"
	apiutils "github.com/desmos-labs/desmostipbot/apis/utils"
)

// SendTip allows to send a tip with the given amount from the tipper to the recipient
func (client *DesmosClient) SendTip(tipper sdk.AccAddress, amount sdk.Coins, recipient sdk.AccAddress) (txHash string, err error) {
	// Get the address of the grantee (the tip bot)
	granteeAddr, err := client.ParseAddress(client.cosmosClient.AccAddress())
	if err != nil {
		return "", err
	}

	// Make sure the authorization is there
	res, err := client.authzClient.Grants(context.Background(), &authz.QueryGrantsRequest{
		Granter:    tipper.String(),
		Grantee:    granteeAddr.String(),
		MsgTypeUrl: sdk.MsgTypeURL(&banktypes.MsgSend{}),
	})
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return "", apiutils.WrapErr(http.StatusNotFound, fmt.Sprintf("No grants found for user %s", granteeAddr.String()))
		}
		return "", err
	}

	// Read the authorization as a SendAuthorization
	var authorization authz.Authorization
	err = client.cosmosClient.Client.Codec.UnpackAny(res.Grants[0].Authorization, &authorization)
	if err != nil {
		return "", err
	}

	sendAuth, ok := authorization.(*banktypes.SendAuthorization)
	if !ok {
		return "", fmt.Errorf("invalid authorization type: expected %T, got %T", banktypes.SendAuthorization{}, authorization)
	}

	// Make sure the spend limit is enough for the tip
	if sendAuth.SpendLimit.IsAllLT(amount) {
		return "", apiutils.WrapErr(http.StatusBadRequest, fmt.Sprintf("Not enough money. Left amount: %s", sendAuth.SpendLimit))
	}

	// Build the message that will execute the authorization to send the tokens
	msgExec := authz.NewMsgExec(granteeAddr, []sdk.Msg{banktypes.NewMsgSend(tipper, recipient, amount)})
	txRes, err := client.cosmosClient.BroadcastTxSync(cosmoswallettypes.NewTransactionData(&msgExec))
	if err != nil {
		return "", err
	}

	return txRes.TxHash, nil
}

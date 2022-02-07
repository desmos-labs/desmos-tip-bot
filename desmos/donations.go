package desmos

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/desmos-labs/plutus/types"
)

// GetDonationDetails returns the details of the donation transaction that has the given hash
func (client *Client) GetDonationDetails(txHash string) (*types.DonationTx, error) {
	res, err := client.txClient.GetTx(context.Background(), &tx.GetTxRequest{Hash: txHash})
	if err != nil {
		return nil, err
	}

	if res.TxResponse.Code != 0 {
		return nil, apiutils.WrapErr(http.StatusBadRequest, fmt.Sprintf(
			"Error while sending the donation transaction: %s", res.TxResponse.RawLog))
	}

	var sendMsgs []*banktypes.MsgSend
	for _, msgAny := range res.Tx.Body.Messages {
		var msg sdk.Msg
		err = client.cdc.UnpackAny(msgAny, &msg)
		if err != nil {
			return nil, err
		}

		if sendMsg, ok := msg.(*banktypes.MsgSend); ok {
			sendMsgs = append(sendMsgs, sendMsg)
		}
	}

	if len(sendMsgs) > 1 {
		return nil, apiutils.WrapErr(http.StatusBadRequest, fmt.Sprintf(
			"Transaction cannot have multiple send messages inside. Required 1, found %d", len(sendMsgs)))
	}

	return types.NewDonationTx(
		res.TxResponse.TxHash,
		sendMsgs[0].FromAddress,
		sendMsgs[0].ToAddress,
		sendMsgs[0].Amount,
	), nil
}

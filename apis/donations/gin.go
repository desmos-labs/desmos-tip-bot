package donations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apiutils "github.com/desmos-labs/desmostipbot/apis/utils"
	"github.com/gin-gonic/gin"
)

type DonationRequest struct {
	// TipperAddress represents the Desmos address of the tipper
	TipperAddress string `json:"tipper_address"`

	// Amount represents the amount to be tipped, serializes as a sdk.Coins instance
	Amount string `json:"amount"`

	// DonationMessage represents the message to be sent along with the donation
	DonationMessage string `json:"donation_message"`

	Platform string `json:"platform"`
	Username string `json:"username"`
}

type DonationResponse struct {
	TxHash string `json:"tx_hash"`
}

func RegisterHandlers(r *gin.Engine, handler *Handler) {
	r.POST("/donations", func(c *gin.Context) {
		bodyBz, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			apiutils.HandleError(c, err)
			return
		}

		var reqBody DonationRequest
		err = json.Unmarshal(bodyBz, &reqBody)
		if err != nil {
			apiutils.HandleError(c, err)
			return
		}

		// Perform the donation
		txHash, err := handler.HandleDonationRequest(reqBody)
		if err != nil {
			apiutils.HandleError(c, err)
			return
		}

		// Send the response
		c.JSON(http.StatusOK, &DonationResponse{
			TxHash: txHash,
		})
	})
}

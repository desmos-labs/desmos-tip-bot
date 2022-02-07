package donations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
)

type DonationRequest struct {
	// TipperUsername is the username the tipper has decided to show inside the donation
	TipperUsername string `json:"tipper_username"`

	// DonationMessage represents the message to be sent along with the donation
	DonationMessage string `json:"donation_message"`

	// RecipientApplication is where the donation was made (Twitch, Twitter, etc)
	RecipientApplication string `json:"application"`

	// RecipientUsername is the recipient on the above specified application
	RecipientUsername string `json:"username"`

	// TxHash is the transaction hash used for the donation
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
		err = handler.HandleDonationRequest(reqBody)
		if err != nil {
			apiutils.HandleError(c, err)
			return
		}

		// Send the response
		c.String(http.StatusOK, "Donation sent successfully")
	})
}

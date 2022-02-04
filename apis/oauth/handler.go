package oauth

import (
	"encoding/hex"
	"net/http"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/oauth/handler"
)

// Handler represents the class that allows to handle various requests
type Handler struct {
	db      *database.Database
	cdc     codec.Codec
	handler *handler.OAuthHandler
}

// NewHandler returns a new Handler instance
func NewHandler(handler *handler.OAuthHandler, cdc codec.Codec, db *database.Database) *Handler {
	return &Handler{
		db:      db,
		cdc:     cdc,
		handler: handler,
	}
}

// HandleAuthenticationTokenRequest allows to handle a request for an authentication token
func (h *Handler) HandleAuthenticationTokenRequest(request TokenRequest) error {
	// Read the public key
	pubKeyBz, err := hex.DecodeString(request.PubKeyBytes)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid public key bytes encoding")
	}

	var pubkey cryptotypes.PubKey
	err = h.cdc.UnmarshalInterface(pubKeyBz, &pubkey)
	if err != nil {
		return err
	}

	// Verify the public key matches the address
	sdkAddr, err := sdk.AccAddressFromBech32(request.DesmosAddress)
	if err != nil {
		return err
	}

	if !sdkAddr.Equals(sdk.AccAddress(pubkey.Address())) {
		return apiutils.WrapErr(http.StatusBadRequest, "Desmos address does not match public key")
	}

	// Verify the signature
	msgBz, err := hex.DecodeString(request.SignedBytes)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signed bytes encoding")
	}

	sigBz, err := hex.DecodeString(request.SignatureBytes)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signature bytes encoding")
	}

	if !pubkey.VerifySignature(msgBz, sigBz) {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signature")
	}

	// Get the token
	token, err := h.handler.GetAuthenticationToken(request.Platform, request.DesmosAddress, request.OAuthCode)
	if err != nil {
		return err
	}

	// Store the token in the database
	return h.db.SaveOAuthToken(token)
}

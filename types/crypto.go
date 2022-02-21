package types

import (
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
)

// SignedRequest contains the data that is sent to the APIs for a signed request
type SignedRequest struct {
	DesmosAddress  string `json:"desmos_address"`
	SignedBytes    string `json:"signed_bytes"`
	PubKeyBytes    string `json:"pubkey_bytes"`
	SignatureBytes string `json:"signature_bytes"`
}

func (r SignedRequest) Verify(nonce string, cdc codec.Codec, amino *codec.LegacyAmino) error {
	// Read the public key
	pubKeyBz, err := hex.DecodeString(r.PubKeyBytes)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid public key bytes encoding")
	}

	var pubkey cryptotypes.PubKey
	err = cdc.UnmarshalInterface(pubKeyBz, &pubkey)
	if err != nil {
		return err
	}

	// Verify the public key matches the address
	sdkAddr, err := sdk.AccAddressFromBech32(r.DesmosAddress)
	if err != nil {
		return err
	}

	if !sdkAddr.Equals(sdk.AccAddress(pubkey.Address())) {
		return apiutils.WrapErr(http.StatusBadRequest, "Desmos address does not match public key")
	}

	msgBz, err := hex.DecodeString(r.SignedBytes)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signed bytes encoding")
	}

	sigBz, err := hex.DecodeString(r.SignatureBytes)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signature bytes encoding")
	}

	if !pubkey.VerifySignature(msgBz, sigBz) {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signature")
	}

	if isValid := verifyDirectSignature(msgBz, nonce, cdc); !isValid {
		err = verifyAminoSignature(msgBz, nonce, amino)
		if err != nil {
			return err
		}
	}

	return nil
}

// verifyDirectSignature tries verifying the request as one being signed using SIGN_MODE_DIRECT.
// Returns true if the signature is valid, false otherwise.
func verifyDirectSignature(msgBz []byte, expectedMemo string, cdc codec.Codec) bool {
	// Verify the signed value contains the OAuth code inside the memo field
	var signDoc tx.SignDoc
	err := cdc.Unmarshal(msgBz, &signDoc)
	if err != nil {
		return false
	}

	var txBody tx.TxBody
	err = cdc.Unmarshal(signDoc.BodyBytes, &txBody)
	if err != nil {
		return false
	}

	if !strings.EqualFold(txBody.Memo, expectedMemo) {
		return false
	}

	return true
}

// verifyAminoSignature tries verifying the request as one being signed using SIGN_MODE_AMINO_JSON.
// Returns an error if something is wrong, nil otherwise.
func verifyAminoSignature(msgBz []byte, expectedMemo string, cdc *codec.LegacyAmino) error {
	var signDoc legacytx.StdSignDoc
	err := cdc.UnmarshalJSON(msgBz, &signDoc)
	if err != nil {
		return apiutils.WrapErr(http.StatusBadRequest, "Invalid signed value. Must be StdSignDoc or SignDoc")
	}

	if !strings.EqualFold(signDoc.Memo, expectedMemo) {
		return apiutils.WrapErr(http.StatusBadRequest, "Signed memo must be equals to OAuth code")
	}

	return nil
}

package user

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/desmos"
	"github.com/desmos-labs/plutus/types"
)

// Handler allows handling user-related requests.
type Handler struct {
	cdc    codec.Codec
	amino  *codec.LegacyAmino
	desmos *desmos.Client
	db     *database.Database
}

// NewHandler returns a new Handler instance
func NewHandler(desmos *desmos.Client, cdc codec.Codec, amino *codec.LegacyAmino, db *database.Database) *Handler {
	return &Handler{
		cdc:    cdc,
		amino:  amino,
		desmos: desmos,
		db:     db,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// DataResponse represents the response for a user details request
type DataResponse struct {
	EnabledIntegrations []string  `json:"enabled_integrations"`
	GrantedAmount       sdk.Coins `json:"granted_amount"`
}

// GetUserData returns the user data for the user having the given address.
func (h *Handler) GetUserData(desmosAddress string) (*DataResponse, error) {
	// Check the address validity
	_, err := sdk.AccAddressFromBech32(desmosAddress)
	if err != nil {
		return nil, apiutils.WrapErr(http.StatusBadRequest, "Invalid Desmos address")
	}

	// Get the grants from the given user towards the bot address
	grants, err := h.desmos.GetGrants(desmosAddress, h.desmos.GetAddress())
	if err != nil {
		return nil, err
	}

	// Get the amount granted
	var grantedAmount sdk.Coins
	for _, grant := range grants {
		var authorization authz.Authorization
		err = h.cdc.UnpackAny(grant.Authorization, &authorization)
		if err != nil {
			return nil, err
		}

		sendAuth, ok := authorization.(*banktypes.SendAuthorization)
		if !ok {
			// Skip non-send authorizations
			continue
		}

		grantedAmount = grantedAmount.Add(sendAuth.SpendLimit...)
	}

	// Get the services accounts
	servicesAccounts, err := h.db.GetServicesAccounts(desmosAddress)
	if err != nil {
		return nil, err
	}

	return &DataResponse{
		EnabledIntegrations: getServices(servicesAccounts),
		GrantedAmount:       grantedAmount,
	}, nil
}

// getServices returns the name of all the services looking into the given accounts, without duplicates
func getServices(accounts []*types.ServiceAccount) []string {
	existingServices := map[string]bool{}
	for _, account := range accounts {
		existingServices[account.Service] = true
	}

	var services []string
	for service := range existingServices {
		services = append(services, service)
	}

	return services
}

// --------------------------------------------------------------------------------------------------------------------

// HandleDisconnectRequest handles the given DisconnectRequest properly
func (h *Handler) HandleDisconnectRequest(request DisconnectRequest) error {
	// Verify the request
	err := request.Verify(request.Platform, h.cdc, h.amino)
	if err != nil {
		return err
	}

	// Delete the service account
	return h.db.DeleteServiceAccount(request.Platform, request.DesmosAddress)
}

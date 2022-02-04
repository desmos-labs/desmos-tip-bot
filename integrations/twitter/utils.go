package twitter

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/plutus/types"
)

// ParseText parses the given text and returns the amount and user specified.
func ParseText(text string) (amount sdk.Coins, user string, err error) {
	if !types.DesmosTipRegEx.MatchString(text) {
		return nil, "", types.ErrInvalidCommand
	}

	tipCmdGroup := types.DesmosTipRegEx.FindString(text)
	parts := strings.Split(tipCmdGroup, " ")

	amt, ok := sdk.NewIntFromString(parts[2])
	if !ok {
		return nil, "", types.ErrInvalidAmount
	}
	amt = amt.MulRaw(1_000_000)

	return sdk.NewCoins(sdk.NewCoin("udaric", amt)), parts[3], nil
}

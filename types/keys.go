package types

import (
	"fmt"
	"regexp"
)

const (
	AppTwitch = "twitch"
)

var (
	DesmosTipRegEx = regexp.MustCompile("^@desmostipbot tip [0-9]+ \\S*")
)

func TipSentMessage(txHash string) string {
	return fmt.Sprintf("Your tip has been sent successfully. Here is your transaction hash: %s", txHash)
}

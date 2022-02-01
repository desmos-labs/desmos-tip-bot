package types

const (
	ServiceStreamlabs     = "streamlabs"
	ServiceStreamElements = "streamelements"
)

type OAuthToken struct {
	DesmosAddress string
	Service       string
	AccessToken   string
	RefreshToken  string
}

func NewOAuthToken(desmosAddress, service, accessToken, refreshToken string) *OAuthToken {
	return &OAuthToken{
		Service:       service,
		DesmosAddress: desmosAddress,
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
	}
}

package oauth

// TokenRequest represents the request body that must be used to get an authorization token
type TokenRequest struct {
	Platform       string `json:"platform"`
	OAuthCode      string `json:"oauth_code"`
	DesmosAddress  string `json:"desmos_address"`
	SignedBytes    string `json:"signed_bytes"`
	PubKeyBytes    string `json:"pubkey_bytes"`
	SignatureBytes string `json:"signature_bytes"`
}

package types

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	cosmoswallettypes "github.com/desmos-labs/cosmos-go-wallet/types"
)

// ParseConfig parses the contents of the file at the given path as a Config instance
func ParseConfig(filePath string) (*Config, error) {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(bz, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Config contains the data used to configure the bot
type Config struct {
	Chain        *ChainConfig                     `yaml:"chain"`
	Account      *cosmoswallettypes.AccountConfig `yaml:"account"`
	Database     *DatabaseConfig                  `yaml:"database"`
	APIs         *APIsConfig                      `yaml:"apis"`
	Integrations *IntegrationsConfig              `yaml:"integrations"`
	Donations    *DonationsConfig                 `yaml:"donations"`
}

type DonationsConfig struct {
	SupportedApps   []string `yaml:"supported_apps"`
	SupportedDenoms map[string]struct {
		CoinGeckoID string `yaml:"coin_gecko_id"`
	} `yaml:"supported_denoms"`
}

type ChainConfig struct {
	*cosmoswallettypes.ChainConfig `yaml:"-,inline"`
	*DesmosClientConfig            `yaml:"-,inline"`
}

// DesmosClientConfig represents the configuration for a DesmosClient instance
type DesmosClientConfig struct {
	GraphQLAddr string `yaml:"gql_addr"`
}

// APIsConfig contains the configuration for the REST APIs provided by the bot
type APIsConfig struct {
	Port uint64 `yaml:"port"`
}

// DatabaseConfig contains the configuration data to connect to a PostgreSQL database
type DatabaseConfig struct {
	URI string `yaml:"uri"`
}

// IntegrationsConfig contains the configuration for the various integrations
type IntegrationsConfig struct {
	Streamlabs *StreamlabsConfig `yaml:"streamlabs"`
	Twitter    *TwitterConfig    `yaml:"twitter"`
}

// StreamlabsConfig contains the configuration data to integrate the bot with Streamlabs
type StreamlabsConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
}

// TwitterConfig contains all the data used to configure the Twitter integration
type TwitterConfig struct {
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
	AccessToken    string `yaml:"access_token"`
	AccessSecret   string `yaml:"access_secret"`
}

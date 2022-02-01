package types

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

// ParseConfig parses the contents of the file at the given path as a Config instance
func ParseConfig(filePath string) (*Config, error) {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(bz, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Config contains the data used to configure the bot
type Config struct {
	Chain    *ChainConfig    `toml:"chain"`
	Database *DatabaseConfig `toml:"database"`
	Twitter  *TwitterConfig  `toml:"twitter"`
	APIs     *APIsConfig     `toml:"apis"`
}

type APIsConfig struct {
	Streamlabs *StreamlabsConfig `toml:"streamlabs"`
}

type StreamlabsConfig struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	RedirectURI  string `toml:"redirect_uri"`
}

// DatabaseConfig contains the configuration data to connect to a PostgreSQL database
type DatabaseConfig struct {
	URI string `toml:"uri"`
}

// ChainConfig contains the configuration data of the chain for which the tipper will run
type ChainConfig struct {
	NodeURI  string `toml:"node_uri"`
	ChainID  string `toml:"chain_id"`
	Mnemonic string `toml:"mnemonic"`
	Fees     string `toml:"fees"`
}

// TwitterConfig contains all the data used to configure the Twitter integration
type TwitterConfig struct {
	ConsumerKey    string `toml:"consumer_key"`
	ConsumerSecret string `toml:"consumer_secret"`
	AccessToken    string `toml:"access_token"`
	AccessSecret   string `toml:"access_secret"`
}

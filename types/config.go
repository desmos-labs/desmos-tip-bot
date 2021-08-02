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
	Chain   *ChainConfig   `toml:"chain"`
	Twitter *TwitterConfig `toml:"twitter"`
}

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

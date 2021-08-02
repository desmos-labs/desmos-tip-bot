module github.com/desmos-labs/desmostipbot

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/desmos-labs/desmos v0.17.4
	github.com/dghubble/go-twitter v0.0.0-20210609183100-2fdbf421508e
	github.com/dghubble/oauth1 v0.7.0
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/pelletier/go-toml v1.9.3
	github.com/rs/zerolog v1.23.0
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.11
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/desmos-labs/cosmos-sdk v0.42.5-0.20210727121119-42c87671b67d

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

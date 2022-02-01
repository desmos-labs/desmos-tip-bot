module github.com/desmos-labs/desmostipbot

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/desmos-labs/desmos v0.17.4
	github.com/dghubble/go-twitter v0.0.0-20210609183100-2fdbf421508e
	github.com/dghubble/oauth1 v0.7.0
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/jmoiron/sqlx v1.3.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lib/pq v1.10.2
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pelletier/go-toml v1.9.3
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.23.0
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.11
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/desmos-labs/cosmos-sdk v0.42.5-0.20210727121119-42c87671b67d

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

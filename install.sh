cd cmd/mttd
VERSION=$(git describe --always --match "v*")
go install -ldflags "-X github.com/cosmos/cosmos-sdk/version.Version=$VERSION -X github.com/cosmos/cosmos-sdk/version.AppName=mttd"

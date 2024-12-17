package contracts

import (
	_ "embed" // embed compiled smart contracts
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"

	evmtypes "github.com/mtt-labs/mtt-chain/x/evm/types"

	"github.com/mtt-labs/mtt-chain/x/erc20/types"
)

var (
	//go:embed compiled_contracts/MttBridge.json
	BridgeJSON []byte // nolint: golint

	// BridgeContract is the compiled erc20 contracts
	BridgeContract evmtypes.CompiledContract

	// BridgeAddress is the erc20 module address
	BridgeAddress common.Address
)

func init() {
	BridgeAddress = types.ModuleAddress

	err := json.Unmarshal(BridgeJSON, &BridgeContract)
	if err != nil {
		panic(err)
	}

	if len(BridgeContract.Bin) == 0 {
		panic("load contracts failed")
	}
}

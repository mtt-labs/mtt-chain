package contracts

import (
	_ "embed" // embed compiled smart contracts
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"

	evmtypes "github.com/mtt-labs/mtt-chain/x/evm/types"

	"github.com/mtt-labs/mtt-chain/x/erc20/types"
)

var (
	//go:embed compiled_contracts/ERC20MinterBurnerDecimals.json
	ERC20MinterBurnerDecimalsJSON []byte // nolint: golint

	// ERC20MinterBurnerDecimalsContract is the compiled erc20 contracts
	ERC20MinterBurnerDecimalsContract evmtypes.CompiledContract

	// ERC20MinterBurnerDecimalsAddress is the erc20 module address
	ERC20MinterBurnerDecimalsAddress common.Address
)

func init() {
	ERC20MinterBurnerDecimalsAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20MinterBurnerDecimalsJSON, &ERC20MinterBurnerDecimalsContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20MinterBurnerDecimalsContract.Bin) == 0 {
		panic("load contracts failed")
	}
}

package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	feemarketkeeper "github.com/mtt-labs/mtt-chain/x/feemarket/keeper"
)

func FixMinGasPrice(ctx sdk.Context, k *feemarketkeeper.Keeper) {
	ctx.Logger().Info("Applying Mtt-chain v2 upgrade." +
		" Fixing fee market min_gas_price.")
	p := k.GetParams(ctx)
	p.MinGasPrice = sdk.NewDec(100000000000)
	err := k.SetParams(ctx, p)
	if err != nil {
		panic(err)
	}
}

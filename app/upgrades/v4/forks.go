package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
)

func UpdateGovParams(ctx sdk.Context, k *govkeeper.Keeper) {
	ctx.Logger().Info("Applying Mtt-chain v4 upgrade. Fixing gov params")
	p := k.GetParams(ctx)
	oneGWei := sdk.NewInt(1000000000)
	tenGWei := sdk.NewInt(10000000000)
	p.MinDeposit = sdk.NewCoins(sdk.Coin{
		Denom:  "amtt",
		Amount: oneGWei.Mul(tenGWei),
	})
	k.SetParams(ctx, p)
}

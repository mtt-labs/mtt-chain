package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	erc20ketkeeper "mtt/x/erc20/keeper"
)

func UpdateErc20Admin(ctx sdk.Context, k *erc20ketkeeper.Keeper) {
	ctx.Logger().Info("Applying Mtt-chain v3 upgrade." +
		" Fixing fee market min_gas_price.")
	p := k.GetParams(ctx)
	p.Admin = "mtt1tz8a2h5wt7fw34jftl6h0z5n97hkt5gn0pzl2x"
	k.SetParams(ctx, p)
}

package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	erc20ketkeeper "github.com/mtt-labs/mtt-chain/x/erc20/keeper"
)

func UpdateErc20Admin(ctx sdk.Context, k *erc20ketkeeper.Keeper) {
	ctx.Logger().Info("Applying Mtt-chain v3 upgrade. Fixing erc20 admin address")
	p := k.GetParams(ctx)
	p.Admin = "mtt1mg9m7dsgceuwmd3e3lz8gr300npxd9kgdkvzdj"
	k.SetParams(ctx, p)
}

package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v3 "mtt/app/upgrades/v3"

	v2 "mtt/app/upgrades/v2"
)

// BeginBlockForks is intended to be ran in a chain upgrade.
func BeginBlockForks(ctx sdk.Context, app *App) {
	switch ctx.BlockHeight() {
	case v2.UpgradeHeight:
		v2.FixMinGasPrice(ctx, &app.FeeMarketKeeper)
	case v3.UpgradeHeight:
		v3.UpdateErc20Admin(ctx, &app.Erc20Keeper)
	}
}

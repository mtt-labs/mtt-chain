package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/mtt-labs/mtt-chain/app/upgrades/v2"
	v3 "github.com/mtt-labs/mtt-chain/app/upgrades/v3"
	v4 "github.com/mtt-labs/mtt-chain/app/upgrades/v4"
	v5 "github.com/mtt-labs/mtt-chain/app/upgrades/v5"
	v6 "github.com/mtt-labs/mtt-chain/app/upgrades/v6"
)

// BeginBlockForks is intended to be ran in a chain upgrade.
func BeginBlockForks(ctx sdk.Context, app *App) {
	switch ctx.BlockHeight() {
	case v2.UpgradeHeight:
		v2.FixMinGasPrice(ctx, &app.FeeMarketKeeper)
	case v3.UpgradeHeight:
		v3.UpdateErc20Admin(ctx, &app.Erc20Keeper)
	case v4.UpgradeHeight:
		v4.UpdateGovParams(ctx, &app.GovKeeper)
	case v5.UpgradeHeight:
		v5.UpdateStakingSlashParams(ctx, app.StakingKeeper, &app.SlashingKeeper)
	case v6.UpgradeHeight:
		v6.UpdateStakingEvmParams(ctx, app.StakingKeeper, app.EvmKeeper)
	}
}

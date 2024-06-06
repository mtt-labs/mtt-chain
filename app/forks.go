package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	//v2 "mtt/cmd/mttd/upgrades/v2"
)

// ScheduleForkUpgrade executes any necessary fork logic for based upon the current
// block height. It sets an upgrade plan once
// the chain reaches the pre-defined upgrade height.
//
// CONTRACT: for this logic to work properly it is required to:
//
//  1. Release a non-breaking patch version so that the chain can set the scheduled upgrade plan at upgrade-height.
//  2. Release the software defined in the upgrade-info
func (app *App) ScheduleForkUpgrade(ctx sdk.Context) {
	upgradePlan := upgradetypes.Plan{
		Height: ctx.BlockHeight(),
	}

	// handle forks
	switch ctx.BlockHeight() {
	//case v2.Height:
	//upgradePlan.Name = v2.UpgradeName
	default:
		// No-op
		return
	}

	if err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan); err != nil {
		panic(
			fmt.Errorf(
				"failed to schedule upgrade %s during BeginBlock at height %d: %w",
				upgradePlan.Name, ctx.BlockHeight(), err,
			),
		)
	}
}

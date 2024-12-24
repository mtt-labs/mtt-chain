package v5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"time"
)

func UpdateStakingSlashParams(ctx sdk.Context, stakingKeeper *stakingkeeper.Keeper, slashKeeper *slashingkeeper.Keeper) {
	ctx.Logger().Info("Applying Mtt-chain v5 upgrade. update Staking Slash params")
	stakingParams := stakingKeeper.GetParams(ctx)
	stakingParams.UnbondingTime = time.Duration(604800) * time.Second
	stakingParams.GlobalMinSelfDelegation = sdk.Coin{
		Denom:  "amtt",
		Amount: sdk.NewInt(10000).Mul(sdk.NewInt(100000000000000000)),
	}
	err := stakingKeeper.SetParams(ctx, stakingParams)
	if err != nil {
		panic(err)
		return
	}

	slashParams := slashKeeper.GetParams(ctx)
	slashParams.SignedBlocksWindow = 50000
	err = slashKeeper.SetParams(ctx, slashParams)
	if err != nil {
		panic(err)
		return
	}
}

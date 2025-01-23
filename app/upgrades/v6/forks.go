package v6

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	evmkeeper "github.com/mtt-labs/mtt-chain/x/evm/keeper"
)

func UpdateStakingEvmParams(ctx sdk.Context, stakingKeeper *stakingkeeper.Keeper, evmKeeper *evmkeeper.Keeper) {
	ctx.Logger().Info("Applying Mtt-chain v6 upgrade. update Staking Evm params")
	oneMtt, _ := sdk.NewIntFromString("1000000000000000000")
	minAmount := sdk.NewInt(10000).Mul(oneMtt)

	stakingParams := stakingKeeper.GetParams(ctx)
	stakingParams.MinCommissionRate, _ = sdk.NewDecFromStr("0.05")
	stakingParams.GlobalMinSelfDelegation = sdk.Coin{
		Denom:  "amtt",
		Amount: minAmount,
	}

	validators := stakingKeeper.GetAllValidators(ctx)
	for _, validator := range validators {
		if validator.MinSelfDelegation.LT(minAmount) {
			validator.MinSelfDelegation = minAmount
		}

		if validator.Commission.Rate.LT(stakingParams.MinCommissionRate) {
			validator.Commission.Rate = stakingParams.MinCommissionRate
		}
		stakingKeeper.SetValidator(ctx, validator)
	}

	err := stakingKeeper.SetParams(ctx, stakingParams)
	if err != nil {
		panic(err)
		return
	}

	evmParams := evmKeeper.GetParams(ctx)
	evmParams.AllowUnprotectedTxs = true
	err = evmKeeper.SetParams(ctx, evmParams)
	if err != nil {
		panic(err)
		return
	}
}

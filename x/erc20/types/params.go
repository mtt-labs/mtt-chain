package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store key
var (
	ParamStoreKeyEnableErc20   = []byte("EnableErc20")
	ParamStoreKeyEnableEVMHook = []byte("EnableEVMHook")
	ParamStoreKeyBridgeAdmin   = []byte("BridgeAdmin")
	ParamStoreKeyBridgeAddress = []byte("BridgeAddress")
)

var _ paramtypes.ParamSet = &Params{}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object
func NewParams(
	enableErc20 bool,
	enableEVMHook bool,
) Params {
	return Params{
		EnableErc20:   enableErc20,
		EnableEVMHook: enableEVMHook,
	}
}

func DefaultParams() Params {
	return Params{
		EnableErc20:   true,
		EnableEVMHook: true,
		Admin:         "",
		Address:       "",
	}
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateAcc(i interface{}) error {
	if addr, ok := i.(string); !ok {
		return errorsmod.Wrapf(
			errortypes.ErrInvalidAddress, "address '%s' is not a valid ethereum hex address",
			addr,
		)
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateString(i interface{}) error {
	if addr, ok := i.(string); !ok {
		return errorsmod.Wrapf(
			errortypes.ErrInvalidAddress, "address '%s' is not a valid ethereum hex address",
			addr,
		)
	}
	return nil
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyEnableErc20, &p.EnableErc20, validateBool),
		paramtypes.NewParamSetPair(ParamStoreKeyEnableEVMHook, &p.EnableEVMHook, validateBool),
		paramtypes.NewParamSetPair(ParamStoreKeyBridgeAdmin, &p.Admin, validateAcc),
		paramtypes.NewParamSetPair(ParamStoreKeyBridgeAddress, &p.Address, validateString),
	}
}

func (p Params) Validate() error { return nil }

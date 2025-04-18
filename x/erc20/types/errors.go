package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// errors
var (
	ErrERC20Disabled          = sdkerrors.Register(ModuleName, 2, "erc20 module is disabled")
	ErrInternalTokenPair      = sdkerrors.Register(ModuleName, 3, "internal ethereum token mapping error")
	ErrTokenPairNotFound      = sdkerrors.Register(ModuleName, 4, "token pair not found")
	ErrTokenPairAlreadyExists = sdkerrors.Register(ModuleName, 5, "token pair already exists")
	ErrUndefinedOwner         = sdkerrors.Register(ModuleName, 6, "undefined owner of contracts pair")
	ErrBalanceInvariance      = sdkerrors.Register(ModuleName, 7, "post transfer balance invariant failed")
	ErrUnexpectedEvent        = sdkerrors.Register(ModuleName, 8, "unexpected event")
	ErrABIPack                = sdkerrors.Register(ModuleName, 9, "contracts ABI pack failed")
	ErrABIUnpack              = sdkerrors.Register(ModuleName, 10, "contracts ABI unpack failed")
	ErrEVMDenom               = sdkerrors.Register(ModuleName, 11, "EVM denomination registration")
	ErrEVMCall                = sdkerrors.Register(ModuleName, 12, "EVM call unexpected error")
	ErrERC20TokenPairDisabled = sdkerrors.Register(ModuleName, 13, "erc20 token pair is disabled")
	ErrCallerNotAdmin         = sdkerrors.Register(ModuleName, 14, "caller not admin")
)

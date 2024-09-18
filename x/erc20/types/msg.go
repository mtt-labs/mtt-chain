package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = &MsgConvertCoin{}
	_ sdk.Msg = &MsgConvertERC20{}
)

const (
	TypeMsgConvertCoin   = "convert_coin"
	TypeMsgConvertERC20  = "convert_ERC20"
	TypeMsgSetBridge     = "set_bridge"
	TypeMsgSetAdmin      = "set_admin"
	TypeMsgSetBeginBlock = "set_begin_block"
	TypeMsgFundMint      = "set_fund_mint"
)

// NewMsgConvertCoin creates a new instance of MsgConvertCoin
func NewMsgConvertCoin(coin sdk.Coin, receiver common.Address, sender sdk.AccAddress) *MsgConvertCoin { // nolint: interfacer
	return &MsgConvertCoin{
		Coin:     coin,
		Receiver: receiver.Hex(),
		Sender:   sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgConvertCoin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertCoin) Type() string { return TypeMsgConvertCoin }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertCoin) ValidateBasic() error {
	if err := ValidateErc20Denom(msg.Coin.Denom); err != nil {
		if err := ibctransfertypes.ValidateIBCDenom(msg.Coin.Denom); err != nil {
			return err
		}
	}

	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(msg.Receiver) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver hex address %s", msg.Receiver)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgConvertCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertCoin) GetSigners() []sdk.AccAddress {
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// NewMsgConvertERC20 creates a new instance of MsgConvertERC20
func NewMsgConvertERC20(amount sdk.Int, receiver sdk.AccAddress, contract, sender common.Address) *MsgConvertERC20 { // nolint: interfacer
	return &MsgConvertERC20{
		ContractAddress: contract.String(),
		Amount:          amount,
		Receiver:        receiver.String(),
		Sender:          sender.Hex(),
	}
}

// Route should return the name of the module
func (msg MsgConvertERC20) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertERC20) Type() string { return TypeMsgConvertERC20 }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertERC20) ValidateBasic() error {
	if !common.IsHexAddress(msg.ContractAddress) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contracts hex address '%s'", msg.ContractAddress)
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid receiver address")
	}
	if !common.IsHexAddress(msg.Sender) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender hex address %s", msg.Sender)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgConvertERC20) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertERC20) GetSigners() []sdk.AccAddress {
	addr := common.HexToAddress(msg.Sender)
	return []sdk.AccAddress{addr.Bytes()}
}

// NewMsgSetBridge creates a new instance of MsgSetBridge
func NewMsgSetBridge(fromAddress sdk.AccAddress, address common.Address) *MsgSetBridge { // nolint: interfacer
	return &MsgSetBridge{
		FromAddress: fromAddress.String(),
		Address:     address.String(),
	}
}

// Route should return the name of the module
func (msg MsgSetBridge) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetBridge) Type() string { return TypeMsgSetBridge }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetBridge) ValidateBasic() error {
	if !common.IsHexAddress(msg.Address) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contracts hex address '%s'", msg.Address)
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid bridge address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetBridge) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetBridge) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// NewMsgSetAdmin creates a new instance of MsgSetAdmin
func NewMsgSetAdmin(fromAddress sdk.AccAddress, address common.Address) *MsgSetAdmin { // nolint: interfacer
	return &MsgSetAdmin{
		FromAddress: fromAddress.String(),
		Address:     address.String(),
	}
}

// Route should return the name of the module
func (msg MsgSetAdmin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetAdmin) Type() string { return TypeMsgSetAdmin }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid admin address")
	}
	_, err = sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid admin address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetAdmin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// NewMsgSetBeginBlock creates a new instance of MsgSetBeginBlock
func NewMsgSetBeginBlock(fromAddress sdk.AccAddress, height uint64) *MsgSetBeginBlock { // nolint: interfacer
	return &MsgSetBeginBlock{
		FromAddress: fromAddress.String(),
		Height:      height,
	}
}

// Route should return the name of the module
func (msg MsgSetBeginBlock) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetBeginBlock) Type() string { return TypeMsgSetBeginBlock }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetBeginBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid receiver address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetBeginBlock) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetBeginBlock) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// NewMsgFundMint creates a new instance of MsgFundMint
func NewMsgFundMint(fromAddress sdk.AccAddress, amount uint64) *MsgFundMint { // nolint: interfacer
	return &MsgFundMint{
		FromAddress: fromAddress.String(),
		Amount:      amount,
	}
}

// Route should return the name of the module
func (msg MsgFundMint) Route() string { return RouterKey }

// Type should return the action
func (msg MsgFundMint) Type() string { return TypeMsgFundMint }

// ValidateBasic runs stateless checks on the message
func (msg MsgFundMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid receiver address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgFundMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgFundMint) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

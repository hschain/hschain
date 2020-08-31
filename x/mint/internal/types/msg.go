package types

import (
	sdk "hschain/types"
)

// RouterKey is they name of the mint module
const RouterKey = ModuleName

// MsgBurn - high level transaction of the coin module
type MsgBurn struct {
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

var _ sdk.Msg = MsgBurn{}

// NewMsgSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgBurn(fromAddr sdk.AccAddress, amount sdk.Coins) MsgBurn {
	return MsgBurn{FromAddress: fromAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgBurn) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgBurn) Type() string { return "burn" }

// ValidateBasic Implements Msg.
func (msg MsgBurn) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgIssue - high level transaction of the coin module
type MsgIssue struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	ToAddress sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
}

var _ sdk.Msg = MsgIssue{}

// NewMsgIssue - construct arbitrary multi-in, multi-out issue msg.
func NewMsgIssue(sender sdk.AccAddress, toAddr sdk.AccAddress, amount sdk.Coins) MsgIssue {
	return MsgIssue{Sender: sender, ToAddress: toAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgIssue) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgIssue) Type() string { return "issue" }

// ValidateBasic Implements Msg.
func (msg MsgIssue) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	if msg.ToAddress.Empty() {
		return sdk.ErrInvalidAddress("missing to address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssue) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

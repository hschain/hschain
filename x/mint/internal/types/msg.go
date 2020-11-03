package types

import (
	"fmt"

	sdk "github.com/hschain/hschain/types"
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

// MsgDistory - high level transaction of the coin module
type MsgDestory struct {
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

var _ sdk.Msg = MsgDestory{}

// NewMsgSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgDestory(fromAddr sdk.AccAddress, amount sdk.Coins) MsgDestory {
	return MsgDestory{FromAddress: fromAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgDestory) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDestory) Type() string { return "conversion-rate" }

// ValidateBasic Implements Msg.
func (msg MsgDestory) ValidateBasic() sdk.Error {
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
func (msg MsgDestory) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDestory) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

type MsgConversionRate struct {
	Fromaddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Rate        sdk.Coins      `json:"amount" yaml:"amount"`
}

var _ sdk.Msg = MsgConversionRate{}

// NewMsgSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgConversionRate(Fromaddress sdk.AccAddress, Rate sdk.Coins) MsgConversionRate {
	return MsgConversionRate{Fromaddress: Fromaddress, Rate: Rate}
}

// Route Implements Msg.
func (msg MsgConversionRate) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgConversionRate) Type() string { return "destory" }

// ValidateBasic Implements Msg.
func (msg MsgConversionRate) ValidateBasic() sdk.Error {
	if msg.Fromaddress.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	if !msg.Rate.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Rate.String())
	}
	if !msg.Rate.IsAllPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgConversionRate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgConversionRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Fromaddress}
}

type MsgAddressPermissions struct {
	Address sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Command string         `json:"command" yaml:"command"`
	Status  int            `json:"status" yaml:"status"`
}

func (Permissions MsgAddressPermissions) String() string {
	return fmt.Sprintf("%v%v%v", Permissions.Address, Permissions.Command, Permissions.Status)
}

type MsgPermissions struct {
	FromAddress sdk.AccAddress        `json:"from_address" yaml:"from_address"`
	Permissions MsgAddressPermissions `json:"address_permissions" yaml:"address_permissions"`
}

var _ sdk.Msg = MsgPermissions{}

// NewMsgSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgPermissions(FromAddress sdk.AccAddress, ToAddress sdk.AccAddress, Command string, Status int) MsgPermissions {
	return MsgPermissions{FromAddress: FromAddress, Permissions: MsgAddressPermissions{Address: ToAddress, Command: Command, Status: Status}}
}

// Route Implements Msg.
func (msg MsgPermissions) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgPermissions) Type() string { return "permissions" }

// ValidateBasic Implements Msg.
func (msg MsgPermissions) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress("missing from address")
	}

	if msg.Permissions.Address.Empty() {
		return sdk.ErrInvalidAddress("missing to address")
	}

	if len(msg.Permissions.Command) <= 0 {
		return sdk.ErrInvalidCoins("Command null")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgPermissions) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgPermissions) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgIssue - high level transaction of the coin module
type MsgAddSysAddress struct {
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Address     sdk.AccAddress `json:"address" yaml:"address"`
	Command     string         `json:"command" yaml:"command"`
}

var _ sdk.Msg = MsgIssue{}

// NewMsgIssue - construct arbitrary multi-in, multi-out issue msg.
func NewMsgAddSysAddress(FromAddress sdk.AccAddress, Address sdk.AccAddress, Command string) MsgAddSysAddress {
	return MsgAddSysAddress{FromAddress: FromAddress, Address: Address, Command: Command}
}

// Route Implements Msg.
func (msg MsgAddSysAddress) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgAddSysAddress) Type() string { return "add sys Address" }

// ValidateBasic Implements Msg.
func (msg MsgAddSysAddress) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress("missing from Address")
	}
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress("missing to address")
	}

	if len(msg.Command) <= 0 {
		return sdk.ErrInvalidAddress("missing to Command")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgAddSysAddress) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgAddSysAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgIssue - high level transaction of the coin module
type MsgSupplement struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"`
	Amount sdk.Coins      `json:"amount" yaml:"amount"`
}

var _ sdk.Msg = MsgIssue{}

// NewMsgIssue - construct arbitrary multi-in, multi-out issue msg.
func NewMsgSupplement(sender sdk.AccAddress, amount sdk.Coins) MsgSupplement {
	return MsgSupplement{Sender: sender, Amount: amount}
}

// Route Implements Msg.
func (msg MsgSupplement) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSupplement) Type() string { return "issue" }

// ValidateBasic Implements Msg.
func (msg MsgSupplement) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
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
func (msg MsgSupplement) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgSupplement) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

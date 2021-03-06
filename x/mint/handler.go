package mint

import (
	"fmt"

	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/mint/internal/keeper"
	"github.com/hschain/hschain/x/mint/internal/types"
)

// NewHandler returns a handler for "bank" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgBurn:
			return handleMsgBurn(ctx, k, msg)

		case types.MsgIssue:
			return handleMsgIssue(ctx, k, msg)

		case types.MsgDestory:
			return handleMsgDestory(ctx, k, msg)

		case types.MsgDestoryUser:
			return handleMsgDestoryUser(ctx, k, msg)

		case types.MsgConversionRate:
			return handleMsgConversionRate(ctx, k, msg)

		case types.MsgPermissions:
			return handleMsgPermissions(ctx, k, msg)

		case types.MsgAddSysAddress:
			return handleMsgAddSysAddress(ctx, k, msg)

		case types.MsgSupplement:
			return handleMsgSupplement(ctx, k, msg)

		case types.MsgVanish:
			return handleMsgVanish(ctx, k, msg)

		case types.MsgVanishUser:
			return handleMsgVanishUser(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized bank message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgBurn MsgBurn.
func handleMsgBurn(ctx sdk.Context, k keeper.Keeper, msg types.MsgBurn) sdk.Result {

	err := k.BurnCoins(ctx, msg.FromAddress, msg.Amount)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// handleMsgBurn MsgBurn.
func handleMsgPermissions(ctx sdk.Context, k keeper.Keeper, msg types.MsgPermissions) sdk.Result {

	if k.GetBalance(ctx, msg.FromAddress).AmountOf(k.BondDenom(ctx)).IsZero() {
		errMsg := fmt.Sprintf("sender must hold %s", k.BondDenom(ctx))
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	k.SetPermissions(ctx, msg.Permissions.Address, msg.Permissions.Command, msg.Permissions.Status)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// handleMsgBurn MsgBurn.
func handleMsgAddSysAddress(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddSysAddress) sdk.Result {

	if k.GetBalance(ctx, msg.FromAddress).AmountOf(k.BondDenom(ctx)).IsZero() {
		errMsg := fmt.Sprintf("sender must hold %s", k.BondDenom(ctx))
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	k.SetSysAddress(ctx, msg.Command, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// MsgConversionRate MsgConversionRate.
func handleMsgConversionRate(ctx sdk.Context, k keeper.Keeper, msg types.MsgConversionRate) sdk.Result {

	if k.GetBalance(ctx, msg.Fromaddress).AmountOf(k.BondDenom(ctx)).IsZero() {

		status := k.GetPermissions(ctx, msg.Fromaddress, "conversion-rate")
		if status != 1 {
			errMsg := fmt.Sprintf("from address not permissions")
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}

	params := k.GetParams(ctx)

	rate := msg.Rate.AmountOf(params.MintDenom)
	if rate.IsZero() {
		return sdk.ErrInvalidAddress("rate must be >0 number").Result()
	}

	k.SetConversionRates(ctx, params.MintDenom, sdk.NewCoin(params.MintDenom, rate))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// MsgDestory MsgDestory.
func handleMsgDestoryUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgDestoryUser) sdk.Result {

	if k.GetBalance(ctx, msg.Sender).AmountOf(k.BondDenom(ctx)).IsZero() {

		status := k.GetPermissions(ctx, msg.Sender, "destory")
		if status != 1 {
			errMsg := fmt.Sprintf("from address not permissions")
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}

	err := k.DestoryCoins(ctx, msg.ToAddress, msg.Amount)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// MsgDestory MsgDestory.
func handleMsgDestory(ctx sdk.Context, k keeper.Keeper, msg types.MsgDestory) sdk.Result {
	err := k.DestoryCoins(ctx, msg.FromAddress, msg.Amount)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// handleMsgIssue MsgIssue.
func handleMsgIssue(ctx sdk.Context, k keeper.Keeper, msg types.MsgIssue) sdk.Result {

	if k.GetBalance(ctx, msg.Sender).AmountOf(k.BondDenom(ctx)).IsZero() {

		status := k.GetPermissions(ctx, msg.Sender, "issue")
		if status != 1 {
			errMsg := fmt.Sprintf("from address not permissions")
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}

	if msg.Amount.Empty() {
		// skip as no coins need to be issue
		errMsg := fmt.Sprintf("no denom found")
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
	for _, coin := range msg.Amount {
		if !k.GetDenomSupply(ctx, coin.Denom).IsZero() {
			errMsg := fmt.Sprintf("denom %s is exist", coin.Denom)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}

	if err := k.IssueCoins(ctx, msg.ToAddress, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSupplement(ctx sdk.Context, k keeper.Keeper, msg types.MsgSupplement) sdk.Result {

	if k.GetBalance(ctx, msg.Sender).AmountOf(k.BondDenom(ctx)).IsZero() {
		errMsg := fmt.Sprintf("from address not permissions")
		return sdk.ErrUnknownRequest(errMsg).Result()

	}
	if msg.Amount.Empty() {
		// skip as no coins need to be supplement
		errMsg := fmt.Sprintf("no denom found")
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	if err := k.SupplementCoins(ctx, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgVanish(ctx sdk.Context, k keeper.Keeper, msg types.MsgVanish) sdk.Result {

	if k.GetBalance(ctx, msg.Sender).AmountOf(k.BondDenom(ctx)).IsZero() {
		errMsg := fmt.Sprintf("from address not permissions")
		return sdk.ErrUnknownRequest(errMsg).Result()

	}
	if msg.Amount.Empty() {
		// skip as no coins need to be vanish
		errMsg := fmt.Sprintf("no denom found")
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	if err := k.VanishCoins(ctx, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgVanishUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgVanishUser) sdk.Result {

	if k.GetBalance(ctx, msg.Sender).AmountOf(k.BondDenom(ctx)).IsZero() {
		errMsg := fmt.Sprintf("from address not permissions")
		return sdk.ErrUnknownRequest(errMsg).Result()

	}
	if msg.Amount.Empty() {
		// skip as no coins need to be vanish
		errMsg := fmt.Sprintf("no denom found")
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	if err := k.VanishUCoins(ctx, msg.FromAddress, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

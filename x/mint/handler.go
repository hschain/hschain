package mint

import (
	"fmt"

	sdk "hschain/types"
	"hschain/x/mint/internal/keeper"
	"hschain/x/mint/internal/types"
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

		case types.MsgConversionRate:
			return handleMsgConversionRate(ctx, k, msg)

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

// MsgConversionRate MsgConversionRate.
func handleMsgConversionRate(ctx sdk.Context, k keeper.Keeper, msg types.MsgConversionRate) sdk.Result {

	if k.GetBalance(ctx, msg.Sender).AmountOf(k.BondDenom(ctx)).IsZero() {
		errMsg := fmt.Sprintf("sender must hold %s", k.BondDenom(ctx))
		return sdk.ErrUnknownRequest(errMsg).Result()
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
		errMsg := fmt.Sprintf("sender must hold %s", k.BondDenom(ctx))
		return sdk.ErrUnknownRequest(errMsg).Result()
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

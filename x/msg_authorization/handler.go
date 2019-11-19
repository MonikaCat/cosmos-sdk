package msg_authorization

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgGrantAuthorization:
			return handleMsgGrantAuthorization(ctx, msg, k)
		case MsgRevokeAuthorization:
			return handleMsgRevokeAuthorization(ctx, msg, k)
		case MsgExecDelegated:
			return handleMsgExecDelegated(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized authorization message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgGrantAuthorization(ctx sdk.Context, msg MsgGrantAuthorization, k Keeper) sdk.Result {
	//TODO
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgRevokeAuthorization(ctx sdk.Context, msg MsgRevokeAuthorization, k Keeper) sdk.Result {
	//TODO
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgExecDelegated(ctx sdk.Context, msg MsgExecDelegated, k Keeper) sdk.Result {
	//TODO
	return sdk.Result{Events: ctx.EventManager().Events()}
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SwapMsg defines a simple interface for getting the token denoms on a swap message route.
type SwapMsgRoute interface {
	TokenInDenom() string
	TokenOutDenom() string
	TokenDenomsOnPath() []string
}

type SwapMsgAmountOut interface {
	SwapMsgRoute

	GetExactTokenOut() sdk.Coin
}

type SwapMsgAmountIn interface {
	SwapMsgRoute

	GetExactTokenIn() sdk.Coin
}

var (
	_ SwapMsgRoute = MsgSwapExactAmountOut{}
	_ SwapMsgRoute = MsgSwapExactAmountIn{}
)

func (msg MsgSwapExactAmountOut) TokenInDenom() string {
	return msg.Routes[0].GetTokenInDenom()
}

func (msg MsgSwapExactAmountOut) TokenOutDenom() string {
	return msg.TokenOut.Denom
}

func (msg MsgSwapExactAmountOut) TokenDenomsOnPath() []string {
	denoms := make([]string, 0, len(msg.Routes)+1)
	for i := 0; i < len(msg.Routes); i++ {
		denoms = append(denoms, msg.Routes[i].TokenInDenom)
	}
	denoms = append(denoms, msg.TokenOutDenom())
	return denoms
}

func (msg MsgSwapExactAmountOut) GetExactTokenOut() sdk.Coin {
	return msg.GetTokenOut()
}

func (msg MsgSwapExactAmountIn) TokenInDenom() string {
	return msg.TokenIn.Denom
}

func (msg MsgSwapExactAmountIn) TokenOutDenom() string {
	lastRouteIndex := len(msg.Routes) - 1
	return msg.Routes[lastRouteIndex].GetTokenOutDenom()
}

func (msg MsgSwapExactAmountIn) TokenDenomsOnPath() []string {
	denoms := make([]string, 0, len(msg.Routes)+1)
	denoms = append(denoms, msg.TokenInDenom())
	for i := 0; i < len(msg.Routes); i++ {
		denoms = append(denoms, msg.Routes[i].TokenOutDenom)
	}
	return denoms
}

func (msg MsgSwapExactAmountIn) GetExactTokenIn() sdk.Coin {
	return msg.GetTokenIn()
}

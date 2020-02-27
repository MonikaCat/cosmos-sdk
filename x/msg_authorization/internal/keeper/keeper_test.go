package keeper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/internal/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

type TestSuite struct {
	suite.Suite
	ctx           sdk.Context
	accountKeeper auth.AccountKeeper
	paramsKeeper  params.Keeper
	bankKeeper    bank.Keeper
	keeper        Keeper
	router        baseapp.Router
}

func (s *TestSuite) SetupTest() {
	s.ctx, s.accountKeeper, s.paramsKeeper, s.bankKeeper, s.keeper, s.router = SetupTestInput()
}

func (s *TestSuite) TestKeeper() {
	err := s.bankKeeper.SetBalances(s.ctx, granterAddr, sdk.NewCoins(sdk.NewInt64Coin("steak", 10000)))
	s.Require().Nil(err)
	s.Require().True(s.bankKeeper.GetBalance(s.ctx, granterAddr, "steak").IsEqual(sdk.NewCoin("steak", sdk.NewInt(10000))))

	s.T().Log("verify that no authorization returns nil")
	authorization, expiration := s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().Nil(authorization)
	s.Require().Zero(expiration)
	now := s.ctx.BlockHeader().Time
	s.Require().NotNil(now)

	newCoins := sdk.NewCoins(sdk.NewInt64Coin("steak", 100))
	s.T().Log("verify if expired authorization is rejected")
	x := types.SendAuthorization{SpendLimit: newCoins}
	s.keeper.Grant(s.ctx, granterAddr, granteeAddr, x, now.Add(-1*time.Hour))
	authorization, _ = s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().Nil(authorization)

	s.T().Log("verify if authorization is accepted")
	x = types.SendAuthorization{SpendLimit: newCoins}
	s.keeper.Grant(s.ctx, granteeAddr, granterAddr, x, now.Add(time.Hour))
	authorization, _ = s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().NotNil(authorization)
	s.Require().Equal(authorization.MsgType(), bank.MsgSend{}.Type())

	s.T().Log("verify fetching authorization with wrong msg type fails")
	authorization, _ = s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgMultiSend{}.Type())
	s.Require().Nil(authorization)

	s.T().Log("verify fetching authorization with wrong grantee fails")
	authorization, _ = s.keeper.GetAuthorization(s.ctx, recipientAddr, granterAddr, bank.MsgMultiSend{}.Type())
	s.Require().Nil(authorization)

	s.T().Log("")

	s.T().Log("verify revoke fails with wrong information")
	s.keeper.Revoke(s.ctx, recipientAddr, granterAddr, bank.MsgSend{}.Type())
	authorization, _ = s.keeper.GetAuthorization(s.ctx, recipientAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().Nil(authorization)

	s.T().Log("verify revoke executes with correct information")
	s.keeper.Revoke(s.ctx, recipientAddr, granterAddr, bank.MsgSend{}.Type())
	authorization, _ = s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().NotNil(authorization)

}

func (s *TestSuite) TestKeeperFees() {
	err := s.bankKeeper.SetBalances(s.ctx, granterAddr, sdk.NewCoins(sdk.NewInt64Coin("steak", 10000)))
	s.Require().Nil(err)
	s.Require().True(s.bankKeeper.GetBalance(s.ctx, granterAddr, "steak").IsEqual(sdk.NewCoin("steak", sdk.NewInt(10000))))

	now := s.ctx.BlockHeader().Time
	s.Require().NotNil(now)

	smallCoin := sdk.NewCoins(sdk.NewInt64Coin("steak", 20))
	someCoin := sdk.NewCoins(sdk.NewInt64Coin("steak", 123))
	//lotCoin := sdk.NewCoins(sdk.NewInt64Coin("steak", 4567))

	msgs := types.MsgExecAuthorized{
		Grantee: granteeAddr,
		Msgs: []sdk.Msg{
			bank.MsgSend{
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("steak", 2)),
				FromAddress: granterAddr,
				ToAddress:   recipientAddr,
			},
		},
	}

	s.T().Log("verify dispatch fails with invalid authorization")
	result, error := s.keeper.DispatchActions(s.ctx, granteeAddr, msgs.Msgs)
	s.Require().Nil(result)
	s.Require().NotNil(error)

	s.T().Log("verify dispatch executes with correct information")
	// grant authorization
	s.keeper.Grant(s.ctx, granteeAddr, granterAddr, types.SendAuthorization{SpendLimit: smallCoin}, now)
	authorization, expiration := s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().NotNil(authorization)
	s.Require().Zero(expiration)
	s.Require().Equal(authorization.MsgType(), bank.MsgSend{}.Type())
	result, error = s.keeper.DispatchActions(s.ctx, granteeAddr, msgs.Msgs)
	s.Require().NotNil(result)
	s.Require().Nil(error)

	authorization, _ = s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().NotNil(authorization)

	s.T().Log("verify dispatch fails with overlimit")
	// grant authorization

	msgs = types.MsgExecAuthorized{
		Grantee: granteeAddr,
		Msgs: []sdk.Msg{
			bank.MsgSend{
				Amount:      someCoin,
				FromAddress: granterAddr,
				ToAddress:   recipientAddr,
			},
		},
	}

	result, error = s.keeper.DispatchActions(s.ctx, granteeAddr, msgs.Msgs)
	s.Require().Nil(result)
	s.Require().NotNil(error)

	authorization, _ = s.keeper.GetAuthorization(s.ctx, granteeAddr, granterAddr, bank.MsgSend{}.Type())
	s.Require().NotNil(authorization)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

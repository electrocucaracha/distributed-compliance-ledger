package keeper

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAccountHeaders = "accounts"
	QueryAccount        = "account"
)

func NewQuerier(accKeeper types.AccountKeeper, authzKeeper authz.Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAccountHeaders:
			return queryAccountHeaders(ctx, req, accKeeper, authzKeeper, cdc)
		case QueryAccount:
			return queryAccount(ctx, req, accKeeper, authzKeeper, cdc, path[1:])
		default:
			return nil, sdk.ErrUnknownRequest("unknown authnext query endpoint")
		}
	}
}

func queryAccountHeaders(ctx sdk.Context, req abci.RequestQuery, accKeeper types.AccountKeeper,
	authzKeeper authz.Keeper, cdc *codec.Codec) (res []byte, err sdk.Error) {
	var params types.QueryAccountHeadersParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.QueryAccountHeadersResult{
		Total: CountTotal(ctx, accKeeper),
		Items: []types.AccountHeader{},
	}

	skipped := 0

	accKeeper.IterateAccounts(ctx, func(account exported.Account) (stop bool) {
		if account.GetPubKey() == nil {
			return false
		}

		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			header := ToAccountHeader(ctx, authzKeeper, account)
			result.Items = append(result.Items, header)
			return false
		}

		return true
	})

	res = codec.MustMarshalJSONIndent(cdc, result)

	return res, nil
}

func queryAccount(ctx sdk.Context, req abci.RequestQuery, accKeeper types.AccountKeeper, authzKeeper authz.Keeper,
	cdc *codec.Codec, path []string) (res []byte, err sdk.Error) {
	accAddr, err_ := sdk.AccAddressFromBech32(path[0])
	if err_ != nil {
		return nil, sdk.ErrInvalidAddress(err_.Error())
	}

	acc := accKeeper.GetAccount(ctx, accAddr)
	header := ToAccountHeader(ctx, authzKeeper, acc)

	res = codec.MustMarshalJSONIndent(cdc, header)

	return res, nil
}

func ToAccountHeader(ctx sdk.Context, authzKeeper authz.Keeper, account exported.Account) types.AccountHeader {
	bechPubKey, err := sdk.Bech32ifyAccPub(account.GetPubKey())
	if err != nil {
		bechPubKey = ""
	}

	header := types.AccountHeader{
		Address:       account.GetAddress(),
		PubKey:        bechPubKey,
		Roles:         authzKeeper.GetAccountRoles(ctx, account.GetAddress()).Roles,
		Coins:         account.GetCoins(),
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	return header
}

func CountTotal(ctx sdk.Context, accKeeper types.AccountKeeper) int {
	res := 0

	accKeeper.IterateAccounts(ctx, func(account exported.Account) (stop bool) {
		if account.GetPubKey() != nil {
			res++
		}

		return false
	})

	return res
}
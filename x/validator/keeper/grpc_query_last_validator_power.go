// Copyright 2022 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LastValidatorPowerAll(c context.Context, req *types.QueryAllLastValidatorPowerRequest) (*types.QueryAllLastValidatorPowerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lastValidatorPowers []types.LastValidatorPower
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lastValidatorPowerStore := prefix.NewStore(store, types.KeyPrefix(types.LastValidatorPowerKeyPrefix))

	pageRes, err := query.Paginate(lastValidatorPowerStore, req.Pagination, func(key []byte, value []byte) error {
		var lastValidatorPower types.LastValidatorPower
		if err := k.cdc.Unmarshal(value, &lastValidatorPower); err != nil {
			return err
		}

		lastValidatorPowers = append(lastValidatorPowers, lastValidatorPower)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLastValidatorPowerResponse{LastValidatorPower: lastValidatorPowers, Pagination: pageRes}, nil
}

func (k Keeper) LastValidatorPower(c context.Context, req *types.QueryGetLastValidatorPowerRequest) (*types.QueryGetLastValidatorPowerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Owner == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLastValidatorPower(
		ctx,
		valAddr,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLastValidatorPowerResponse{LastValidatorPower: val}, nil
}

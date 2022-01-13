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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ComplianceInfoAll(c context.Context, req *types.QueryAllComplianceInfoRequest) (*types.QueryAllComplianceInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var complianceInfos []types.ComplianceInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	complianceInfoStore := prefix.NewStore(store, types.KeyPrefix(types.ComplianceInfoKeyPrefix))

	pageRes, err := query.Paginate(complianceInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var complianceInfo types.ComplianceInfo
		if err := k.cdc.Unmarshal(value, &complianceInfo); err != nil {
			return err
		}

		complianceInfos = append(complianceInfos, complianceInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllComplianceInfoResponse{ComplianceInfo: complianceInfos, Pagination: pageRes}, nil
}

func (k Keeper) ComplianceInfo(c context.Context, req *types.QueryGetComplianceInfoRequest) (*types.QueryGetComplianceInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetComplianceInfo(
		ctx,
		req.Vid,
		req.Pid,
		req.SoftwareVersion,
		req.CertificationType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetComplianceInfoResponse{ComplianceInfo: val}, nil
}

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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProposedCertificateAll(c context.Context, req *types.QueryAllProposedCertificateRequest) (*types.QueryAllProposedCertificateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposedCertificates []types.ProposedCertificate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proposedCertificateStore := prefix.NewStore(store, types.KeyPrefix(types.ProposedCertificateKeyPrefix))

	pageRes, err := query.Paginate(proposedCertificateStore, req.Pagination, func(key []byte, value []byte) error {
		var proposedCertificate types.ProposedCertificate
		if err := k.cdc.Unmarshal(value, &proposedCertificate); err != nil {
			return err
		}

		proposedCertificates = append(proposedCertificates, proposedCertificate)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposedCertificateResponse{ProposedCertificate: proposedCertificates, Pagination: pageRes}, nil
}

func (k Keeper) ProposedCertificate(c context.Context, req *types.QueryGetProposedCertificateRequest) (*types.QueryGetProposedCertificateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProposedCertificate(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProposedCertificateResponse{ProposedCertificate: val}, nil
}

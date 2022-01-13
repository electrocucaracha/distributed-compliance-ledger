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

package keeper_test

/* TODO issue 99
import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestProposedCertificateQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProposedCertificate(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetProposedCertificateRequest
		response *types.QueryGetProposedCertificateResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProposedCertificateRequest{
				Subject:      msgs[0].Subject,
				SubjectKeyId: msgs[0].SubjectKeyId,
			},
			response: &types.QueryGetProposedCertificateResponse{ProposedCertificate: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProposedCertificateRequest{
				Subject:      msgs[1].Subject,
				SubjectKeyId: msgs[1].SubjectKeyId,
			},
			response: &types.QueryGetProposedCertificateResponse{ProposedCertificate: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProposedCertificateRequest{
				Subject:      strconv.Itoa(100000),
				SubjectKeyId: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ProposedCertificate(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestProposedCertificateQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProposedCertificate(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProposedCertificateRequest {
		return &types.QueryAllProposedCertificateRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProposedCertificateAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProposedCertificate), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProposedCertificate),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProposedCertificateAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProposedCertificate), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProposedCertificate),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProposedCertificateAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ProposedCertificate),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProposedCertificateAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
*/

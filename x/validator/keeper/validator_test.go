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

/* FIXME issue 99

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)

		keeper.SetValidator(ctx, items[i])
	}
	return items
}

func TestValidatorGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidator(ctx,
			item.Owner,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestValidatorRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidator(ctx,
			item.Owner,
		)
		_, found := keeper.GetValidator(ctx,
			item.Owner,
		)
		require.False(t, found)
	}
}

func TestValidatorGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllValidator(ctx))
}
*/

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
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	dclauthkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func ValidatorKeeper(t testing.TB, dclauthK *dclauthkeeper.Keeper) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	// TODO issue 99: might be not the best solution
	if dclauthK != nil {
		stateStore.MountStoreWithDB(dclauthK.StoreKey(), sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(dclauthK.MemKey(), sdk.StoreTypeMemory, nil)
	}

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
		dclauthK,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}

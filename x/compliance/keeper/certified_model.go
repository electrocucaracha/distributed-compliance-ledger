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
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetCertifiedModel set a specific certifiedModel in the store from its index.
func (k Keeper) SetCertifiedModel(ctx sdk.Context, certifiedModel types.CertifiedModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertifiedModelKeyPrefix))
	b := k.cdc.MustMarshal(&certifiedModel)
	store.Set(types.CertifiedModelKey(
		certifiedModel.Vid,
		certifiedModel.Pid,
		certifiedModel.SoftwareVersion,
		certifiedModel.CertificationType,
	), b)
}

// GetCertifiedModel returns a certifiedModel from its index.
func (k Keeper) GetCertifiedModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) (val types.CertifiedModel, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertifiedModelKeyPrefix))

	b := store.Get(types.CertifiedModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCertifiedModel removes a certifiedModel from the store.
func (k Keeper) RemoveCertifiedModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertifiedModelKeyPrefix))
	store.Delete(types.CertifiedModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
}

// GetAllCertifiedModel returns all certifiedModel.
func (k Keeper) GetAllCertifiedModel(ctx sdk.Context) (list []types.CertifiedModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertifiedModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CertifiedModel
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

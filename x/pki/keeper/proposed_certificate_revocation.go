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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetProposedCertificateRevocation set a specific proposedCertificateRevocation in the store from its index.
func (k Keeper) SetProposedCertificateRevocation(ctx sdk.Context, proposedCertificateRevocation types.ProposedCertificateRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	b := k.cdc.MustMarshal(&proposedCertificateRevocation)
	store.Set(types.ProposedCertificateRevocationKey(
		proposedCertificateRevocation.Subject,
		proposedCertificateRevocation.SubjectKeyId,
	), b)
}

// GetProposedCertificateRevocation returns a proposedCertificateRevocation from its index.
func (k Keeper) GetProposedCertificateRevocation(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.ProposedCertificateRevocation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))

	b := store.Get(types.ProposedCertificateRevocationKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposedCertificateRevocation removes a proposedCertificateRevocation from the store.
func (k Keeper) RemoveProposedCertificateRevocation(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	store.Delete(types.ProposedCertificateRevocationKey(
		subject,
		subjectKeyId,
	))
}

// GetAllProposedCertificateRevocation returns all proposedCertificateRevocation.
func (k Keeper) GetAllProposedCertificateRevocation(ctx sdk.Context) (list []types.ProposedCertificateRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedCertificateRevocation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Check if the Proposed Certificate Revocation record associated with a
// Subject/SubjectKeyID combination is present in the store.
func (k Keeper) IsProposedCertificateRevocationPresent(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	return store.Has(types.ProposedCertificateRevocationKey(
		subject,
		subjectKeyId,
	))
}

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

package pki

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the approvedCertificates
	for _, elem := range genState.ApprovedCertificatesList {
		k.SetApprovedCertificates(ctx, elem)
	}
	// Set all the proposedCertificate
	for _, elem := range genState.ProposedCertificateList {
		k.SetProposedCertificate(ctx, elem)
	}
	// Set all the childCertificates
	for _, elem := range genState.ChildCertificatesList {
		k.SetChildCertificates(ctx, elem)
	}
	// Set all the proposedCertificateRevocation
	for _, elem := range genState.ProposedCertificateRevocationList {
		k.SetProposedCertificateRevocation(ctx, elem)
	}
	// Set all the revokedCertificates
	for _, elem := range genState.RevokedCertificatesList {
		k.SetRevokedCertificates(ctx, elem)
	}
	// Set all the uniqueCertificate
	for _, elem := range genState.UniqueCertificateList {
		k.SetUniqueCertificate(ctx, elem)
	}
	// Set if defined
	if genState.ApprovedRootCertificates != nil {
		k.SetApprovedRootCertificates(ctx, *genState.ApprovedRootCertificates)
	}
	// Set if defined
	if genState.RevokedRootCertificates != nil {
		k.SetRevokedRootCertificates(ctx, *genState.RevokedRootCertificates)
	}
	// Set all the approvedCertificatesBySubject
	for _, elem := range genState.ApprovedCertificatesBySubjectList {
		k.SetApprovedCertificatesBySubject(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ApprovedCertificatesList = k.GetAllApprovedCertificates(ctx)
	genesis.ProposedCertificateList = k.GetAllProposedCertificate(ctx)
	genesis.ChildCertificatesList = k.GetAllChildCertificates(ctx)
	genesis.ProposedCertificateRevocationList = k.GetAllProposedCertificateRevocation(ctx)
	genesis.RevokedCertificatesList = k.GetAllRevokedCertificates(ctx)
	genesis.UniqueCertificateList = k.GetAllUniqueCertificate(ctx)
	// Get all approvedRootCertificates
	approvedRootCertificates, found := k.GetApprovedRootCertificates(ctx)
	if found {
		genesis.ApprovedRootCertificates = &approvedRootCertificates
	}
	// Get all revokedRootCertificates
	revokedRootCertificates, found := k.GetRevokedRootCertificates(ctx)
	if found {
		genesis.RevokedRootCertificates = &revokedRootCertificates
	}
	genesis.ApprovedCertificatesBySubjectList = k.GetAllApprovedCertificatesBySubject(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

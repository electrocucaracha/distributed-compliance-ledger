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

package pki_test

/* TODO issue 99
import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ApprovedCertificatesList: []types.ApprovedCertificates{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		ProposedCertificateList: []types.ProposedCertificate{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		ChildCertificatesList: []types.ChildCertificates{
			{
				Issuer:         "0",
				AuthorityKeyId: "0",
			},
			{
				Issuer:         "1",
				AuthorityKeyId: "1",
			},
		},
		ProposedCertificateRevocationList: []types.ProposedCertificateRevocation{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		RevokedCertificatesList: []types.RevokedCertificates{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		UniqueCertificateList: []types.UniqueCertificate{
			{
				Issuer:       "0",
				SerialNumber: "0",
			},
			{
				Issuer:       "1",
				SerialNumber: "1",
			},
		},
		ApprovedRootCertificates: &types.ApprovedRootCertificates{
			Certs: []*types.CertificateIdentifier{},
		},
		RevokedRootCertificates: &types.RevokedRootCertificates{
			Certs: []*types.CertificateIdentifier{},
		},
		ApprovedCertificatesBySubjectList: []types.ApprovedCertificatesBySubject{
			{
				Subject: "0",
			},
			{
				Subject: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PkiKeeper(t)
	pki.InitGenesis(ctx, *k, genesisState)
	got := pki.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.ApprovedCertificatesList, got.ApprovedCertificatesList)
	require.ElementsMatch(t, genesisState.ProposedCertificateList, got.ProposedCertificateList)
	require.ElementsMatch(t, genesisState.ChildCertificatesList, got.ChildCertificatesList)
	require.ElementsMatch(t, genesisState.ProposedCertificateRevocationList, got.ProposedCertificateRevocationList)
	require.ElementsMatch(t, genesisState.RevokedCertificatesList, got.RevokedCertificatesList)
	require.ElementsMatch(t, genesisState.UniqueCertificateList, got.UniqueCertificateList)
	require.Equal(t, genesisState.ApprovedRootCertificates, got.ApprovedRootCertificates)
	require.Equal(t, genesisState.RevokedRootCertificates, got.RevokedRootCertificates)
	require.ElementsMatch(t, genesisState.ApprovedCertificatesBySubjectList, got.ApprovedCertificatesBySubjectList)
	// this line is used by starport scaffolding # genesis/test/assert
}
*/

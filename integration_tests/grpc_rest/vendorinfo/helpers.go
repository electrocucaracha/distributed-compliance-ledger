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

package vendorinfo

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	vendorinfotypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func NewMsgCreateVendorInfo(vid int32, signer string) *vendorinfotypes.MsgCreateVendorInfo {

	return &vendorinfotypes.MsgCreateVendorInfo{
		Creator:              signer,
		VendorID:             vid,
		VendorName:           testconstants.VendorName,
		CompanyLegalName:     testconstants.CompanyLegalName,
		CompanyPrefferedName: testconstants.CompanyPreferredName,
		VendorLandingPageURL: testconstants.VendorLandingPageUrl,
	}
}

func NewMsgUpdateVendorInfo(vid int32, signer string) *vendorinfotypes.MsgUpdateVendorInfo {

	return &vendorinfotypes.MsgUpdateVendorInfo{
		Creator:              signer,
		VendorID:             vid,
		VendorName:           testconstants.VendorName + "/new",
		CompanyLegalName:     testconstants.CompanyLegalName + "/new",
		CompanyPrefferedName: testconstants.CompanyPreferredName + "/new",
		VendorLandingPageURL: testconstants.VendorLandingPageUrl + "/new",
	}
}

func AddVendorInfo(
	suite *utils.TestSuite,
	msg *vendorinfotypes.MsgCreateVendorInfo,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {

	msg.Creator = suite.GetAddress(signerName).String()
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func GetVendorInfo(
	suite *utils.TestSuite,
	vid int32,
) (*vendorinfotypes.VendorInfo, error) {

	var res vendorinfotypes.VendorInfo

	if suite.Rest {
		var resp vendorinfotypes.QueryGetVendorInfoResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/vendorinfo/vendors/%v", vid), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		vendorinfoClient := vendorinfotypes.NewQueryClient(grpcConn)
		resp, err := vendorinfoClient.VendorInfo(
			context.Background(),
			&vendorinfotypes.QueryGetVendorInfoRequest{VendorID: vid},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	}

	return &res, nil
}

func GetVendorInfos(suite *utils.TestSuite) (res []vendorinfotypes.VendorInfo, err error) {
	if suite.Rest {
		var resp vendorinfotypes.QueryAllVendorInfoResponse
		err := suite.QueryREST("/dcl/vendorinfo/vendors", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		vendorinfoClient := vendorinfotypes.NewQueryClient(grpcConn)
		resp, err := vendorinfoClient.VendorInfoAll(
			context.Background(),
			&vendorinfotypes.QueryAllVendorInfoRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	}

	return res, nil
}

func VendorInfoDemo(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// New vendor adds first vendorinfo
	createFirstVendorInfoMsg := NewMsgCreateVendorInfo(vid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createFirstVendorInfoMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check first vendorinfo is added
	receivedVendorInfo, err := GetVendorInfo(suite, createFirstVendorInfoMsg.VendorID)
	require.NoError(suite.T, err)
	require.Equal(suite.T, createFirstVendorInfoMsg.VendorID, receivedVendorInfo.VendorID)
	require.Equal(suite.T, createFirstVendorInfoMsg.VendorName, receivedVendorInfo.VendorName)
	require.Equal(suite.T, createFirstVendorInfoMsg.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(suite.T, createFirstVendorInfoMsg.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(suite.T, createFirstVendorInfoMsg.VendorLandingPageURL, receivedVendorInfo.VendorLandingPageURL)

	// Get all vendorinfos
	_, err = GetVendorInfos(suite)
	require.NoError(suite.T, err)
}

/* Error cases */

func AddVendorInfoByNonVendor(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// register new account without Vendor role
	testHouseName := utils.RandString()
	vid := int32(tmrand.Uint16())
	testHouseAccount := test_dclauth.CreateAccount(
		suite,
		testHouseName,
		dclauthtypes.AccountRoles{dclauthtypes.TestHouse},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
	)

	require.NotContains(suite.T, testHouseAccount.Roles, dclauthtypes.Vendor)

	// try to add createVendorInfoMsg
	createVendorInfoMsg := NewMsgCreateVendorInfo(vid, testHouseAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createVendorInfoMsg}, testHouseName, testHouseAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnauthorized.Is(err))
}

func AddVendorInfoByDifferentVendor(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// register new Vendor account
	vendorName := utils.RandString()
	vid := int32(tmrand.Uint16())
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid+1,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
	)

	// try to add createVendorInfoMsg
	createVendorInfoMsg := NewMsgCreateVendorInfo(vid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createVendorInfoMsg}, vendorName, vendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnauthorized.Is(err))
}

func AddVendorInfoTwice(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// register new Vendor account
	vendorName := utils.RandString()
	vid := int32(tmrand.Uint16())
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
	)

	// add vendorinfo
	createVendorInfoMsg := NewMsgCreateVendorInfo(vid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createVendorInfoMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// add the same vendorinfo second time
	_, err = AddVendorInfo(suite, createVendorInfoMsg, vendorName, vendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrInvalidRequest.Is(err))
}

func GetVendorInfoForUnknown(suite *utils.TestSuite) {
	_, err := GetVendorInfo(suite, int32(tmrand.Uint16()))
	require.Error(suite.T, err)
	suite.AssertNotFound(err)
}

func GetVendorInfoForInvalidVid(suite *utils.TestSuite) {
	// zero vid
	_, err := GetVendorInfo(suite, 0)
	require.Error(suite.T, err)
	// FIXME: Consider adding validation for queries.
	// require.True(suite.T, sdkerrors.ErrInvalidRequest.Is(err))
	suite.AssertNotFound(err)
}

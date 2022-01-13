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

package model

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the vendorProducts
	for _, elem := range genState.VendorProductsList {
		k.SetVendorProducts(ctx, elem)
	}
	// Set all the model
	for _, elem := range genState.ModelList {
		k.SetModel(ctx, elem)
	}
	// Set all the modelVersion
	for _, elem := range genState.ModelVersionList {
		k.SetModelVersion(ctx, elem)
	}
	// Set all the modelVersions
	for _, elem := range genState.ModelVersionsList {
		k.SetModelVersions(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.VendorProductsList = k.GetAllVendorProducts(ctx)
	genesis.ModelList = k.GetAllModel(ctx)
	genesis.ModelVersionList = k.GetAllModelVersion(ctx)
	genesis.ModelVersionsList = k.GetAllModelVersions(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

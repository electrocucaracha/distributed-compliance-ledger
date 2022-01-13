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

package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmtypes "github.com/tendermint/tendermint/types"
	validatortypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index.
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GenTxs: []json.RawMessage{},

		// this line is used by starport scaffolding # genesis/types/default
	}
}

// FIXME issue 99 review
// NewGenesisState creates a new GenesisState object.
func NewGenesisState(genTxs []json.RawMessage) *GenesisState {
	// Ensure genTxs is never nil, https://github.com/cosmos/cosmos-sdk/issues/5086
	if len(genTxs) == 0 {
		genTxs = make([]json.RawMessage, 0)
	}
	return &GenesisState{
		GenTxs: genTxs,
	}
}

// FIXME issue 99 review
// NewGenesisStateFromTx creates a new GenesisState object
// from auth transactions.
func NewGenesisStateFromTx(txJSONEncoder sdk.TxEncoder, genTxs []sdk.Tx) *GenesisState {
	genTxsBz := make([]json.RawMessage, len(genTxs))
	for i, genTx := range genTxs {
		var err error
		genTxsBz[i], err = txJSONEncoder(genTx)
		if err != nil {
			panic(err)
		}
	}
	return NewGenesisState(genTxsBz)
}

// GetGenesisStateFromAppState gets the genutil genesis state from the expected app state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}
	return &genesisState
}

// SetGenesisStateInAppState sets the genutil genesis state within the expected app state.
func SetGenesisStateInAppState(
	cdc codec.JSONCodec, appState map[string]json.RawMessage, genesisState *GenesisState,
) map[string]json.RawMessage {
	genesisStateBz := cdc.MustMarshalJSON(genesisState)
	appState[ModuleName] = genesisStateBz
	return appState
}

// GenesisStateFromGenDoc creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenDoc(genDoc tmtypes.GenesisDoc) (genesisState map[string]json.RawMessage, err error) {
	if err = json.Unmarshal(genDoc.AppState, &genesisState); err != nil {
		return genesisState, err
	}
	return genesisState, nil
}

// FIXME issue 99 review
// GenesisStateFromGenFile creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenFile(genFile string) (genesisState map[string]json.RawMessage, genDoc *tmtypes.GenesisDoc, err error) {
	if !tmos.FileExists(genFile) {
		return genesisState, genDoc,
			fmt.Errorf("%s does not exist, run `init` first", genFile)
	}

	genDoc, err = tmtypes.GenesisDocFromFile(genFile)
	if err != nil {
		return genesisState, genDoc, err
	}

	genesisState, err = GenesisStateFromGenDoc(*genDoc)
	return genesisState, genDoc, err
}

// ValidateGenesis validates GenTx transactions.
func ValidateGenesis(genesisState *GenesisState, txJSONDecoder sdk.TxDecoder) error {
	for i, genTx := range genesisState.GenTxs {
		var tx sdk.Tx
		tx, err := txJSONDecoder(genTx)
		if err != nil {
			return err
		}

		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return errors.New(
				"must provide genesis Tx with exactly 1 CreateValidator message")
		}

		if _, ok := msgs[0].(*validatortypes.MsgCreateValidator); !ok {
			return fmt.Errorf(
				"genesis transaction %v does not contain a MsgCreateValidator", i)
		}
	}
	return nil
}

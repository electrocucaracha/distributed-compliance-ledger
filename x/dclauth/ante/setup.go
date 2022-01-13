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

package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Sets Infinite Gas Meter instead of default one in SetUpContextDecorator
type InfiniteGasSetUpContextDecorator struct{}

func NewInfiniteGasSetUpContextDecorator() InfiniteGasSetUpContextDecorator {
	return InfiniteGasSetUpContextDecorator{}
}

func (sud InfiniteGasSetUpContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	return next(newCtx, tx, simulate)
}

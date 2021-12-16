/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { JsonObject, JsonProperty } from 'json2typescript';
import { StdTxnValue } from './std-txn-value';

@JsonObject('StdTxn')
export class StdTxn {

  @JsonProperty('type', String)
  type = 'cosmos-sdk/StdTx';

  @JsonProperty('value', StdTxnValue)
  value: StdTxnValue = new StdTxnValue();

  constructor(init?: Partial<StdTxn>) {
    Object.assign(this, init);
  }
}
#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# FIXME issue 99: enable once implemented
exit 0

set -euo pipefail
source integration_tests/cli/common.sh

# Preparation of Actors

vid=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

echo "Create other CertificationCenter account"
create_new_account second_zb_account "CertificationCenter"

# Body


pid=$RANDOM
sv=$RANDOM
svs=$RANDOM
echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $svs $vendor_account

test_divider

echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv"
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | dcld tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Revoke Certification for uncertificate Model with VID: $vid PID: $pid"
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
certification_type="zigbee"
result=$(echo "test1234" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="$certification_type" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="zigbee")
check_response "$result" "\"value\": false"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} and SV: ${sv} before compliance record was created"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="zigbee")
check_response "$result" "\"value\": true"
echo "$result"



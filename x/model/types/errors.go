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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// Model Error Codes.
	ErrModelAlreadyExists       = sdkerrors.Register(ModuleName, 501, "model already exists")
	ErrModelDoesNotExist        = sdkerrors.Register(ModuleName, 502, "model does not exist")
	ErrVendorProductsDoNotExist = sdkerrors.Register(ModuleName, 504, "vendor products do not exist")

	// Model Version Error Codes.
	ErrSoftwareVersionStringInvalid = sdkerrors.Register(ModuleName, 511, "software version string invalid")
	ErrFirmwareDigestsInvalid       = sdkerrors.Register(ModuleName, 512, "firmware digests invalid")
	ErrCDVersionNumberInvalid       = sdkerrors.Register(ModuleName, 513, "CD version number invalid")
	ErrOtaURLInvalid                = sdkerrors.Register(ModuleName, 514, "OTA URL invalid")
	ErrOtaMissingInformation        = sdkerrors.Register(ModuleName, 515, "OTA missing information")
	ErrReleaseNotesURLInvalid       = sdkerrors.Register(ModuleName, 516, "release notes URL invalid")
	ErrModelVersionDoesNotExist     = sdkerrors.Register(ModuleName, 517, "model version does not exist")
	ErrNoModelVersionsExist         = sdkerrors.Register(ModuleName, 518, "no model versions exist")
	ErrModelVersionAlreadyExists    = sdkerrors.Register(ModuleName, 519, "model version already exists")
	ErrOtaURLCannotBeSet            = sdkerrors.Register(ModuleName, 520, "OTA URL cannot be set")
	ErrMaxSVLessThanMinSV           = sdkerrors.Register(ModuleName, 521, "max software version less than min software version")
)

func NewErrModelAlreadyExists(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(ErrModelAlreadyExists,
		"Model associated with vid=%v and pid=%v already exists on the ledger",
		vid, pid)
}

func NewErrModelDoesNotExist(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(ErrModelDoesNotExist,
		"No model associated with vid=%v and pid=%v exist on the ledger",
		vid, pid)
}

func NewErrVendorProductsDoNotExist(vid interface{}) error {
	return sdkerrors.Wrapf(ErrVendorProductsDoNotExist,
		"No vendor products associated with vid=%v exist on the ledger",
		vid)
}

func NewErrSoftwareVersionStringInvalid(softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrSoftwareVersionStringInvalid,
		"SoftwareVersionString %v is invalid. It should be greater then 1 and less then 64 character long",
		softwareVersion)
}

func NewErrFirmwareDigestsInvalid(firmwareDigests interface{}) error {
	return sdkerrors.Wrapf(ErrFirmwareDigestsInvalid,
		"firmwareDigests %v is of invalid length. Maximum length should be less then 512",
		firmwareDigests)
}

func NewErrCDVersionNumberInvalid(cdVersionNumber interface{}) error {
	return sdkerrors.Wrapf(ErrCDVersionNumberInvalid,
		"CDVersionNumber %v is invalid. It should be a 16 bit unsigned integer",
		cdVersionNumber)
}

func NewErrOtaURLInvalid(otaURL interface{}) error {
	return sdkerrors.Wrapf(ErrOtaURLInvalid,
		"OtaURL %v is invalid. Maximum length should be less then 256",
		otaURL)
}

func NewErrOtaMissingInformation() error {
	return sdkerrors.Wrap(ErrOtaMissingInformation,
		"OtaFileSize, OtaChecksum and OtaChecksumType are required if OtaUrl is provided")
}

func NewErrReleaseNotesURLInvalid(releaseNotesURL interface{}) error {
	return sdkerrors.Wrapf(ErrReleaseNotesURLInvalid,
		"ReleaseNotesURLInvalid %v is invalid. Maximum length should be less then 256",
		releaseNotesURL)
}

func NewErrModelVersionDoesNotExist(vid interface{}, pid interface{}, softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrModelVersionDoesNotExist,
		"No model version associated with vid=%v, pid=%v and softwareVersion=%v exist on the ledger",
		vid, pid, softwareVersion)
}

func NewErrNoModelVersionsExist(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(ErrNoModelVersionsExist,
		"No versions associated with vid=%v and pid=%v exist on the ledger",
		vid, pid)
}

func NewErrModelVersionAlreadyExists(vid interface{}, pid interface{}, softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrModelVersionAlreadyExists,
		"Model Version already exists on ledger with vid=%v pid=%v and softwareVersion=%v exist on the ledger",
		vid, pid, softwareVersion)
}

func NewErrOtaURLCannotBeSet(vid interface{}, pid interface{}, softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrOtaURLCannotBeSet,
		"OTA URL cannot be set for model version associated with vid=%v, pid=%v "+
			"and softwareVersion=%v because OTA was not set for this model initially",
		vid, pid, softwareVersion)
}

func NewErrMaxSVLessThanMinSV(minApplicableSoftwareVersion interface{},
	maxApplicableSoftwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrMaxSVLessThanMinSV,
		"MaxApplicableSoftwareVersion %v is less than MinApplicableSoftwareVersion %v",
		maxApplicableSoftwareVersion, minApplicableSoftwareVersion)
}

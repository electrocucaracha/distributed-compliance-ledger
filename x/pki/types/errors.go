package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pki module sentinel errors
const (
	DefaultCodespace string = ModuleName

	CodeProposedCertificateAlreadyExists           uint32 = 401
	CodeProposedCertificateDoesNotExist            uint32 = 402
	CodeCertificateAlreadyExists                   uint32 = 403
	CodeCertificateDoesNotExist                    uint32 = 404
	CodeProposedCertificateRevocationAlreadyExists uint32 = 405
	CodeProposedCertificateRevocationDoesNotExist  uint32 = 406
	CodeRevokedCertificateDoesNotExist             uint32 = 407
	CodeInappropriateCertificateType               uint32 = 408
	CodeInvalidCertificate                         uint32 = 409
)

func ErrProposedCertificateAlreadyExists(subject string, subjectKeyID string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeProposedCertificateAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateDoesNotExist(subject string, subjectKeyID string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeProposedCertificateDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger. "+
			"The cerificate either does not exists or already approved.", subject, subjectKeyID))
}

func ErrCertificateAlreadyExists(issuer string, serialNumber string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeCertificateAlreadyExists,
		fmt.Sprintf("X509 certificate associated with the combination of "+
			"issuer=%v and serialNumber=%v already exists on the ledger", issuer, serialNumber))
}

func ErrCertificateDoesNotExist(subject string, subjectKeyID string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeCertificateDoesNotExist,
		fmt.Sprintf("No X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateRevocationAlreadyExists(subject string, subjectKeyID string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeProposedCertificateRevocationAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate revocation associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateRevocationDoesNotExist(subject string, subjectKeyID string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeProposedCertificateRevocationDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate revocation associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger.", subject, subjectKeyID))
}

func ErrRevokedCertificateDoesNotExist(subject string, subjectKeyID string) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeRevokedCertificateDoesNotExist,
		fmt.Sprintf("No revoked X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrInappropriateCertificateType(error interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeInappropriateCertificateType, fmt.Sprintf("%v", error))
}

func ErrCodeInvalidCertificate(error interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeInvalidCertificate, fmt.Sprintf("%v", error))
}

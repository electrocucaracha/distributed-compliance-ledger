# PKI types

#    plain ones
# starport scaffold --module pki type CertificateIdentifier subject subjectKeyId 
# starport scaffold --module pki type Certificate pemCert serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner subject subjectKeyId

#    messages
starport scaffold --module pki message ProposeAddX509RootCert cert --signer signer
starport scaffold --module pki message ApproveAddX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message AddX509Cert cert --signer signer
starport scaffold --module pki message ProposeRevokeX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message ApproveRevokeX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message RevokeX509Cert subject subjectKeyId --signer signer

# CRUD data types
starport scaffold --module pki map ApprovedCertificates pemCert serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner --index subject,subjectKeyId --no-message
starport scaffold --module pki map ProposedCertificate pemCert serialNumber owner approvals:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map ChildCertificates subject subjectKeyId --index issuer,authorityKeyId --no-message
starport scaffold --module pki map ProposedCertificateRevocation  approvals:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map RevokedCertificates pemCert serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner --index subject,subjectKeyId --no-message
# PKI types

#    plain ones
# starport scaffold --module pki type CertificateIdentifier subject subjectKeyID 
# starport scaffold --module pki type Certificate pemCert serialNumber issuer authorityKeyID rootSubject rootSubjectKeyID isRoot:bool owner subject subjectKeyID

#    messages
starport scaffold --module pki message ProposeAddX509RootCert cert --signer signer
starport scaffold --module pki message ApproveAddX509RootCert subject subjectKeyID --signer signer
starport scaffold --module pki message AddX509Cert cert --signer signer
starport scaffold --module pki message ProposeRevokeX509RootCert subject subjectKeyID --signer signer
starport scaffold --module pki message ApproveRevokeX509RootCert subject subjectKeyID --signer signer
starport scaffold --module pki message RevokeX509Cert subject subjectKeyID --signer signer

# CRUD data types
starport scaffold --module pki list ApprovedCertificates pemCert serialNumber issuer authorityKeyID rootSubject rootSubjectKeyID isRoot:bool owner subject subjectKeyID --no-message
starport scaffold --module pki map ProposedCertificate pemCert serialNumber owner approvals:strings --index subject,subjectKeyID --no-message
starport scaffold --module pki list ChildCertificates subject subjectKeyID issuer authorityKeyID --no-message
starport scaffold --module pki map ProposedCertificateRevocation  approvals:strings --index subject,subjectKeyID --no-message
starport scaffold --module pki list RevokedCertificates pemCert serialNumber issuer authorityKeyID rootSubject rootSubjectKeyID isRoot:bool owner subject subjectKeyID --no-message
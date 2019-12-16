package certificate

type CertificateType byte

const (
	RSA CertificateType = iota
	EC
)

var (
	certTypeToStringMap = map[CertificateType]string{
		RSA: "RSA",
		EC:  "EC",
	}
)

func (value CertificateType) String() string {
	return certTypeToStringMap[value]
}

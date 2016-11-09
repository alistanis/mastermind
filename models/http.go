package models

type HttpConfig struct {
	Host                 string
	Port                 string
	Cert                 string
	Key                  string
	CertificateAuthority string
	UseSSL               bool
}

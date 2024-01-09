package models

type Credential interface {
	GetUserOcid() string
	GetTenancyOcid() string
	GetCompartmentOcid() string
	GetPrivateKey() string
	GetFingerprint() string
}

type MyCredential struct {
	UserOcid        string
	TenancyOcid     string
	CompartmentOcid string
	PrivateKey      string
	Fingerprint     string
}

func (c MyCredential) GetUserOcid() string {
	return c.UserOcid
}

func (c MyCredential) GetTenancyOcid() string {
	return c.TenancyOcid
}

func (c MyCredential) GetCompartmentOcid() string {
	return c.CompartmentOcid
}

func (c MyCredential) GetPrivateKey() string {
	return c.PrivateKey
}

func (c MyCredential) GetFingerprint() string {
	return c.Fingerprint
}

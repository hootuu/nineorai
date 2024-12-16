package domains

import "github.com/hootuu/gelato/crtpto/sha256x"

type Password string

func NewPassword(str string) Password {
	return Password(sha256x.SHA256(str))
}

func (p Password) String() string {
	return string(p)
}

func (p Password) IsValid() bool {
	return sha256x.IsSHA256(p.String())
}

func (p Password) EncryptedPassword() EncryptedPassword {
	return NewEncryptedPassword(p)
}

type EncryptedPassword string

func NewEncryptedPassword(pwd Password) EncryptedPassword {
	return EncryptedPassword(sha256x.SHA256(string(pwd)))
}

func (ep EncryptedPassword) Match(pwd Password) bool {
	return ep == pwd.EncryptedPassword()
}

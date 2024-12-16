package trustee

import (
	"github.com/hootuu/gelato/crtpto/rand"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/idx"
	"github.com/hootuu/nineorai/keys"
	"strings"
)

type Create struct {
	Trustee  bool
	BoundID  string
	Password string
}

func NewRandCreate(trustee bool) *Create {
	if !trustee {
		return &Create{Trustee: false}
	}
	pwd, err := rand.String(16)
	if err != nil {
		pwd = idx.New()
	}
	return &Create{
		Trustee:  true,
		BoundID:  idx.New(),
		Password: pwd,
	}
}

type CreateResult struct {
	Key keys.Key
}

func (c Create) Validate() *errors.Error {
	if len(strings.TrimSpace(c.BoundID)) == 0 {
		return errors.Verify("require bound id")
	}
	if len(strings.TrimSpace(c.Password)) == 0 {
		return errors.Verify("require password")
	}
	return nil
}

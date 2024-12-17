package domains

import "github.com/hootuu/gelato/errors"

type Signature string

func (s Signature) Validate() *errors.Error {
	if len(s) == 0 {
		return errors.Verify("invalid signature")
	}
	return nil
}

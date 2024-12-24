package trustee

import (
	"github.com/hootuu/gelato/errors"
	"strings"
)

const (
	ByLink    = "link"
	ByAddress = "address"
)

type Exists struct {
	By   string `json:"by" bson:"by"` // link | address
	Para string `json:"para" bson:"para"`
}

type ExistsResult struct {
	Exists bool `json:"exists" bson:"exists"`
}

func (e Exists) Validate() *errors.Error {
	switch e.By {
	case ByLink, ByAddress:
		if len(strings.TrimSpace(e.Para)) == 0 {
			return errors.Verify("require para")
		}
	default:
		return errors.Verify("by must be either link or address")
	}
	return nil
}

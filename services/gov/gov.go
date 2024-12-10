package gov

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
)

type Governance interface {
	Initial() (domains.StakeAddr, io.Error)
}

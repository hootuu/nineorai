package gov

import (
	"nineorai/domains"
	"nineorai/io"
)

type Governance interface {
	Initial() (domains.StakeAddr, io.Error)
}

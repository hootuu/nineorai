package trustee

import (
	"github.com/hootuu/nineorai/io"
)

type Service interface {
	Create(req *io.Request[Create]) *io.Response[CreateResult]
	Exists(req *io.Request[Exists]) *io.Response[ExistsResult]
}

package identity

import (
	"github.com/hootuu/nineorai/io"
)

type Service interface {
	Create(req *io.Request[Create]) *io.Response[CreateResult]
	Get(req *io.Request[Get]) *io.Response[GetResult]
	GetByLink(req *io.Request[GetByLink]) *io.Response[GetResult]
	GetByNineoraID(req *io.Request[GetByNineoraID]) *io.Response[GetResult]
}

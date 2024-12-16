package asset

import "github.com/hootuu/nineorai/io"

type Service interface {
	Create(req *io.Request[Create]) *io.Response[CreateResult]
}

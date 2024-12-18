package token

import "github.com/hootuu/nineorai/io"

type Service interface {
	Create(req *io.Request[Create]) *io.Response[CreateResult]
	Mint(req *io.Request[Mint]) *io.Response[MintResult]

	AccLoadByAuth(req *io.Request[AccLoadByAuth]) *io.Response[AccLoadResult]
}

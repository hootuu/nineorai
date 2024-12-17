package token

import "github.com/hootuu/nineorai/io"

type Service interface {
	Create(req *io.Request[Create]) *io.Response[CreateResult]
	Mint(req *io.Request[Mint]) *io.Response[MintResult]

	LoadAccountByAuthority(req *io.Request[LoadByAuthority]) *io.Response[LoadResult]
}

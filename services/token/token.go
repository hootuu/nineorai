package token

import "github.com/hootuu/nineorai/io"

type Service interface {
	Create(req *io.Request[Create]) *io.Response[CreateResult]
	Mint(req *io.Request[Mint]) *io.Response[MintResult]

	AccLoadByAuth(req *io.Request[AccLoadByAuth]) *io.Response[AccLoadResult]
	AccLoadByLink(req *io.Request[AccLoadByLink]) *io.Response[AccLoadResult]
	AccCreate(req *io.Request[AccountCreate]) *io.Response[AccountCreateResult]

	Transfer(req *io.Request[Transfer]) *io.Response[TransferResult]

	TxLoad(req *io.Request[TxLoad]) *io.Response[TxLoadResult]
}

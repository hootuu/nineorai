package token

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
)

type Service interface {
	io.CtrlManager
	io.TagManager
	io.MetaManager

	Create(req Create) (domains.TokenAddr, io.Error)

	Mint(req Mint, ctx MintCtx) (domains.Tx, io.Error)

	Transfer(req Transfer, ctx TransferCtx) (domains.Tx, io.Error)

	GetInfo(stake domains.TokenAddr) (domains.Token, io.Error)

	GetSupply(stake domains.TokenAddr) (uint64, io.Error)

	GetAccount(stake domains.TokenAddr, id domains.IdentityAddr) (domains.TokenAccount, io.Error)
}

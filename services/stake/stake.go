package stake

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
)

type Service interface {
	io.CtrlManager
	io.TagManager
	io.MetaManager

	Create(req Create, ctx CreateCtx) (domains.StakeAddr, io.Error)

	Mint(req Mint, ctx MintCtx) (domains.Tx, io.Error)

	Transfer(req Transfer, ctx TransferCtx) (domains.Tx, io.Error)

	GetInfo(stake domains.StakeAddr) (domains.Stake, io.Error)

	GetSupply(stake domains.Stake) (uint64, io.Error)

	GetAccount(stake domains.StakeAddr, id domains.IdentityAddr) (domains.StakeAccount, io.Error)
}

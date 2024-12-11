package identity

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
)

type Service interface {
	io.CtrlManager
	io.TagManager
	io.MetaManager

	Create(req Create, ctx CreateCtx) (*keys.NineoraID, *domains.IdentityAddr, io.Error)

	Networks(who domains.IdentityAddr) (io.Pagination[domains.Network], io.Error)

	TokenAccounts(
		who domains.IdentityAddr,
		network domains.NetworkAddr,
		page io.Page,
	) (io.Pagination[domains.TokenAccount], io.Error)

	StakeAccounts(
		who domains.IdentityAddr,
		network domains.NetworkAddr,
		page io.Page,
	) (io.Pagination[domains.StakeAccount], io.Error)

	AssetAccounts(who domains.IdentityAddr,
		addr domains.NetworkAddr,
		page io.Page,
	) (io.Pagination[domains.AssetAccount], io.Error)

	TokenTx(
		who domains.IdentityAddr,
		network domains.NetworkAddr,
		page io.Page,
	) (io.Pagination[domains.Tx], io.Error)

	StakeTx(
		who domains.IdentityAddr,
		network domains.NetworkAddr,
		page io.Page,
	) (io.Pagination[domains.Tx], io.Error)

	Tx(
		who domains.IdentityAddr,
		network domains.NetworkAddr,
		page io.Page,
	) (io.Pagination[domains.Tx], io.Error)
}

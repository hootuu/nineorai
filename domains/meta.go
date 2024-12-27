package domains

import "github.com/hootuu/gelato/errors"

const (
	MetaName        = "name"
	MetaIco         = "ico"
	MetaUri         = "uri"
	MetaDescription = "description"
	MetaImages      = "images"
	MetaVideos      = "videos"

	MetaAssetTypeId   = "asset:id"
	MetaAssetTypeCode = "asset:type:code"
	MetaAssetTypeName = "asset:type:name"
	MetaAssetStatus   = "asset:status"
	MetaAssetLocation = "asset:location"
	MetaAssetLocMemo  = "asset:loc:memo"

	MetaAssetSTOC  = "asset:stoc"  // The address for Short-term order commission
	MetaAssetMTRPD = "asset:mtrpd" // The address for Mid-term reward pool dividend
	MetaAssetLTEI  = "asset:ltei"  //The address for Long-term equity incentive
)

type MetaID string
type MetaRef string

type Meta = Dict

func NewMeta(values ...map[string]interface{}) (Meta, *errors.Error) {
	return NewDict(values...)
}

func MustNewMeta(values ...map[string]interface{}) Meta {
	meta, _ := NewMeta(values...)
	return meta
}

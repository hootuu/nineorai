package domains

import "github.com/hootuu/gelato/errors"

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

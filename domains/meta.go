package domains

type MetaID string
type MetaRef string

type Meta = Dict

func NewMeta(values ...map[string]interface{}) (Meta, error) {
	return NewDict(values...)
}

const (
	MetaName = "name"
	MetaUri  = "uri"
)

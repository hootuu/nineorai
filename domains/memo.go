package domains

type Memo = Dict

func NewMemo() Memo {
	return MustNewDict()
}

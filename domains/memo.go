package domains

const (
	MemoMemo        = "memo"
	MemoDescription = "description"
)

type Memo = Dict

func NewMemo() Memo {
	return MustNewDict()
}

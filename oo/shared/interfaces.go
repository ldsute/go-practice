package shared

type LLNode interface {
	SetValue(v int) error
	GetValue() int
}

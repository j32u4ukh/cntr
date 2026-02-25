package cntr

type IHashable[T any] interface {
	GetHash() string
	Compare(other T) bool
}

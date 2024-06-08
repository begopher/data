package simple

type Source[T any] interface {
	Change(value T) error
	Value() (T, error)
}

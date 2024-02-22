package data

type Source[T any] interface {
	Value() (T, error)
	Change(T) error
}

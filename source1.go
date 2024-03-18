package data

type Source1[T any] interface {
	Change(value T) error
	Value() (T, error)
}

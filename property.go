package data

type Property[T any] interface {
	Change(value T) (error, bool)
	Value() (T, error)
}

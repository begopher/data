package data

type Source2[T any] interface {
	Change(value T, checksum uint32) error
	Value() (T, error)
}

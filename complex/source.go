package complex

type Source[T any] interface {
	Change(value T, checksum uint32) error
	Value() (T, error)
}

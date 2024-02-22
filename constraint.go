package data

type Constraint[T any] interface {
	Evaluate(T) error
}

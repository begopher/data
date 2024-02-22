package data

func Constraints[T any](many ...Constraint[T]) Constraint[T] {
	for _, c := range many {
		if c == nil {
			panic("data.Constraints: nil is not allowed")
		}
	}
	return constraints[T]{many}
}

type constraints[T any] struct {
	many []Constraint[T]
}

func (c constraints[T]) Evaluate(value T) error {
	for _, constraint := range c.many {
		if err := constraint.Evaluate(value); err != nil {
			return err
		}
	}
	return nil
}

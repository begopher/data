package simple

import(
	"github.com/begopher/data"
)

func Property[T comparable] (
	value *T,
	constraint data.Constraint[T],
	source Source[T],
) data.Property[T] {
        if constraint == nil {
	   panic("simple.Property: constraint cannot be nil")
	}
	if source == nil {
	   panic("simple.Property: source cannot be nil")
	}
	return &property[T]{
	       value:      value,
	       constraint: constraint,
	       source:     source,
	}
}

type property[T comparable] struct {
	value *T
	constraint data.Constraint[T]
	source Source[T]
}

func (s *property[T]) Change(value T) (error, bool) {
	got, err:= s.Value()
	if err != nil{
		return err, false
	}
	if got == value {
		return nil, false
	}
	if err := s.constraint.Evaluate(value); err != nil {
		return err, false
	}
	if err := s.source.Change(value); err != nil {
		return err, false
	}
	s.value = &value
	return nil, true
}

func (s *property[T]) Value() (T, error) {
	if s.value != nil {
		return *s.value, nil
	}
	value, err := s.source.Value()
	if err != nil {
		s.value = &value
	}
	return value, err
}

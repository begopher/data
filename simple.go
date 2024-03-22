package data

func Simple[T comparable] (
	value *T,
	constraint Constraint[T],
	source Source1[T],
) Property[T] {
        if constraint == nil {
	   panic("data.Simple: constraint cannot be nil")
	}
	if source == nil {
	   panic("data.Simple: source cannot be nil")
	}
	return &simple[T]{
	       value:      value,
	       constraint: constraint,
	       source:     source,
	}
}

type simple[T comparable] struct {
	value *T
	constraint Constraint[T]
	source Source1[T]
}

func (s *simple[T]) Change(value T) (error, bool) {
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

func (s *simple[T]) Value() (T, error) {
	if s.value != nil {
		return *s.value, nil
	}
	value, err := s.source.Value()
	if err != nil {
		s.value = &value
	}
	return value, err
}

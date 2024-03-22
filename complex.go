package data

func Complex[T interface { Checksum() uint32 }] (
	value *T,
	checksum uint32,
	constraint Constraint[T],
	source Source2[T],
) Property[T] {
        if constraint == nil {
	   panic("data.Complex: constraint cannot be nil")
	}
	if source == nil {
	   panic("data.Complex: source cannot be nil")
	}
	return &complex[T]{
	       value:      value,
	       checksum:   checksum,
	       constraint: constraint,
	       source:     source,
	}
}

type complex[T interface{ Checksum() uint32 }] struct {
	value *T
	checksum uint32
	constraint Constraint[T]
	source Source2[T]
}

func (e *complex[T]) Change(value T) (error, bool) {
	checksum := value.Checksum()
	if checksum == e.checksum {
		return nil, false
	}
	if err := e.constraint.Evaluate(value); err != nil {
		return err, false
	}
	err := e.source.Change(value, checksum)
	if err == nil {
		e.checksum = checksum
		e.value = &value
		return nil, true
	}
	return err, false
}

func (e *complex[T]) Value() (T, error) {
	if e.value != nil {
		return *e.value, nil
	}
	value, err := e.source.Value()
	if err != nil {
		e.value = &value
	}
	return value, err
}

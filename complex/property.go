package complex

import(
	"github.com/begopher/data"
)

func Property[T interface { Checksum() uint32 }] (
	value *T,
	checksum uint32,
	constraint data.Constraint[T],
	source Source[T],
) data.Property[T] {
        if constraint == nil {
	   panic("complex.Property: constraint cannot be nil")
	}
	if source == nil {
	   panic("complex.Property: source cannot be nil")
	}
	return &property[T]{
	       value:      value,
	       checksum:   checksum,
	       constraint: constraint,
	       source:     source,
	}
}

type property[T interface{ Checksum() uint32 }] struct {
	value *T
	checksum uint32
	constraint data.Constraint[T]
	source Source[T]
}

func (e *property[T]) Change(value T) (error, bool) {
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

func (e *property[T]) Value() (T, error) {
	if e.value != nil {
		return *e.value, nil
	}
	value, err := e.source.Value()
	if err != nil {
		e.value = &value
	}
	return value, err
}

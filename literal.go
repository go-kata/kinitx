package kinitx

import (
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

// Literal represents a bootstrapper that registers an object to use directly instead of it creation.
type Literal struct {
	// t specifies the type of an object that is registered by this literal.
	t reflect.Type
	// object specifies the object to register.
	object reflect.Value
}

// NewLiteral returns a new literal.
func NewLiteral(x interface{}) (*Literal, error) {
	if x == nil {
		return nil, kerror.New(kerror.ERuntime, "value expected, nil given")
	}
	return &Literal{
		t:      reflect.TypeOf(x),
		object: reflect.ValueOf(x),
	}, nil
}

// MustNewLiteral is a variant of the NewLiteral that panics on error.
func MustNewLiteral(x interface{}) *Literal {
	l, err := NewLiteral(x)
	if err != nil {
		panic(err)
	}
	return l
}

// Bootstrap implements the kinit.Bootstrapper interface.
func (l *Literal) Bootstrap(arena *kinit.Arena) error {
	return arena.Register(l.t, l.object, kdone.Noop)
}

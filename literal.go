package kinitx

import (
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

// Literal represents an initializer that registers an object to use directly instead of it creation.
type Literal struct {
	// t specifies the type of an object that is registered by this initializer.
	t reflect.Type
	// object specifies the object to register.
	object reflect.Value
}

// NewLiteral returns a new initializer.
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
	i, err := NewLiteral(x)
	if err != nil {
		panic(err)
	}
	return i
}

// Initialize implements the kinit.Initializer interface.
func (i *Literal) Initialize(arena *kinit.Arena) error {
	return arena.Register(i.t, i.object, kdone.Noop)
}

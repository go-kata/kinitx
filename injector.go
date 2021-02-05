package kinitx

import (
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

// Injector represents a functor that provides an object to use directly instead of it creation.
type Injector struct {
	// t specifies the type of an object that is provided by this injector.
	t reflect.Type
	// object specifies the provided object.
	object reflect.Value
}

// NewInjector returns a new injector.
//
// The argument x must not be nil.
func NewInjector(x interface{}) (*Injector, error) {
	if x == nil {
		return nil, kerror.New(kerror.EViolation, "value expected, nil given")
	}
	return &Injector{
		t:      reflect.TypeOf(x),
		object: reflect.ValueOf(x),
	}, nil
}

// MustNewInjector is a variant of the NewInjector that panics on error.
func MustNewInjector(x interface{}) *Injector {
	f, err := NewInjector(x)
	if err != nil {
		panic(err)
	}
	return f
}

// Parameters implements the kinit.Functor interface.
func (f *Injector) Parameters() []reflect.Type {
	if f == nil {
		return nil
	}
	return []reflect.Type{runtimeType}
}

// Call implements the kinit.Functor interface.
func (f *Injector) Call(a ...reflect.Value) ([]kinit.Functor, error) {
	if f == nil {
		return nil, nil
	}
	if len(a) != 1 {
		return nil, kerror.Newf(kerror.EViolation,
			"%s injector expects %d argument(s), %d given", f.t, 1, len(a))
	}
	if a[0].Type() != runtimeType {
		return nil, kerror.Newf(kerror.EViolation,
			"%s injector expects argument %d to be of %s type, %s given",
			f.t, 1, runtimeType, a[0].Type())
	}
	if err := a[0].Interface().(*kinit.Runtime).Register(f.t, f.object, kdone.Noop); err != nil {
		return nil, err
	}
	return nil, nil
}

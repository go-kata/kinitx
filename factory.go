package kinitx

import (
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
)

// Factory represents a constructor based on a function.
type Factory struct {
	// t specifies the type of an object that is created by this constructor.
	t reflect.Type
	// function specifies the reflection to a function value.
	function reflect.Value
	// inTypes specifies types of function input parameters.
	inTypes []reflect.Type
	// objectOutIndex specifies the index of a function output parameter that contains a created object.
	objectOutIndex int
	// destructorOutIndex specifies the index of a function output parameter that contains a destructor.
	// The value -1 means that a function doesn't return a destructor.
	destructorOutIndex int
	// errorOutIndex specifies the index of a function output parameter that contains an error.
	// The value -1 means that a function doesn't return an error.
	errorOutIndex int
}

// NewFactory returns a new constructor.
//
// The argument x must be a function that is compatible with one of following signatures
// (T is an arbitrary Go type):
//
//     func(...) T;
//
//     func(...) (T, error);
//
//     func(...) (T, kdone.Destructor, error).
//
func NewFactory(x interface{}) (*Factory, error) {
	if x == nil {
		return nil, kerror.New(kerror.ERuntime, "function expected, nil given")
	}
	ft := reflect.TypeOf(x)
	fv := reflect.ValueOf(x)
	if ft.Kind() != reflect.Func {
		return nil, kerror.Newf(kerror.ERuntime, "function expected, %s given", ft)
	}
	if fv.IsNil() {
		return nil, kerror.New(kerror.ERuntime, "function expected, nil given")
	}
	c := &Factory{
		function: fv,
	}
	numIn := ft.NumIn()
	if ft.IsVariadic() {
		numIn--
	}
	c.inTypes = make([]reflect.Type, numIn)
	for i := 0; i < numIn; i++ {
		c.inTypes[i] = ft.In(i)
	}
	switch ft.NumOut() {
	default:
		return nil, kerror.Newf(kerror.ERuntime, "function %s is not a factory", ft)
	case 1:
		c.t = ft.Out(0)
		c.objectOutIndex = 0
		c.destructorOutIndex = -1
		c.errorOutIndex = -1
	case 2:
		if ft.Out(1) != errorType {
			return nil, kerror.Newf(kerror.ERuntime, "function %s is not a factory", ft)
		}
		c.t = ft.Out(0)
		c.objectOutIndex = 0
		c.destructorOutIndex = -1
		c.errorOutIndex = 1
	case 3:
		if ft.Out(1) != destructorType || ft.Out(2) != errorType {
			return nil, kerror.Newf(kerror.ERuntime, "function %s is not a factory", ft)
		}
		c.t = ft.Out(0)
		c.objectOutIndex = 0
		c.destructorOutIndex = 1
		c.errorOutIndex = 2
	}
	return c, nil
}

// MustNewFactory is a variant of the NewFactory that panics on error.
func MustNewFactory(x interface{}) *Factory {
	c, err := NewFactory(x)
	if err != nil {
		panic(err)
	}
	return c
}

// Type implements the kinit.Constructor interface.
func (c *Factory) Type() reflect.Type {
	if c == nil {
		return nil
	}
	return c.t
}

// Parameters implements the kinit.Constructor interface.
func (c *Factory) Parameters() []reflect.Type {
	if c == nil {
		return nil
	}
	types := make([]reflect.Type, len(c.inTypes))
	copy(types, c.inTypes)
	return types
}

// Create implements the kinit.Constructor interface.
func (c *Factory) Create(a ...reflect.Value) (reflect.Value, kdone.Destructor, error) {
	if c == nil {
		return reflect.Value{}, kdone.Noop, nil
	}
	if len(a) != len(c.inTypes) {
		return reflect.Value{}, kdone.Noop, kerror.Newf(kerror.ERuntime,
			"%s constructor expects %d argument(s), %d given",
			c.t, len(c.inTypes), len(a))
	}
	for i, v := range a {
		if v.Type() != c.inTypes[i] {
			return reflect.Value{}, kdone.Noop, kerror.Newf(kerror.ERuntime,
				"%s constructor expects argument %d to be of %s type, %s given",
				c.t, i+1, c.inTypes[i], v.Type())
		}
	}
	out := c.function.Call(a)
	obj := out[c.objectOutIndex]
	var dtor kdone.Destructor = kdone.Noop
	if c.destructorOutIndex >= 0 {
		if v := out[c.destructorOutIndex].Interface(); v != nil {
			dtor = v.(kdone.Destructor)
		}
	}
	var err error
	if c.errorOutIndex >= 0 {
		if v := out[c.errorOutIndex].Interface(); v != nil {
			err = v.(error)
		}
	}
	return obj, dtor, err
}

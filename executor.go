package kinitx

import (
	"reflect"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

// Executor represents an executor based on a function.
type Executor struct {
	// function specifies the reflection to a function value.
	function reflect.Value
	// inTypes specifies types of function input parameters.
	inTypes []reflect.Type
	// executorOutIndex specifies the index of a function output parameter that contains an executor.
	// The value -1 means that a function doesn't return an executor.
	executorOutIndex int
	// errorOutIndex specifies the index of a function output parameter that contains an error.
	// The value -1 means that a function doesn't return an error.
	errorOutIndex int
}

// NewExecutor returns a new executor.
// The argument x must be a function that is compatible with one of following signatures:
//
//     func(...)
//
//     func(...) error
//
//     func(...) (kinit.Executor, error)
//
func NewExecutor(x interface{}) (*Executor, error) {
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
	e := &Executor{
		function: fv,
	}
	numIn := ft.NumIn()
	if ft.IsVariadic() {
		numIn--
	}
	e.inTypes = make([]reflect.Type, numIn)
	for i := 0; i < numIn; i++ {
		e.inTypes[i] = ft.In(i)
	}
	switch ft.NumOut() {
	default:
		return nil, kerror.Newf(kerror.ERuntime, "function %s is not an executor", ft)
	case 0:
		e.executorOutIndex = -1
		e.errorOutIndex = -1
	case 1:
		if ft.Out(0) != errorType {
			return nil, kerror.Newf(kerror.ERuntime, "function %s is not an executor", ft)
		}
		e.executorOutIndex = -1
		e.errorOutIndex = 0
	case 2:
		if ft.Out(0) != executorType || ft.Out(1) != errorType {
			return nil, kerror.Newf(kerror.ERuntime, "function %s is not an executor", ft)
		}
		e.executorOutIndex = 0
		e.errorOutIndex = 1
	}
	return e, nil
}

// MustNewExecutor is a variant of the NewExecutor that panics on error.
func MustNewExecutor(x interface{}) *Executor {
	e, err := NewExecutor(x)
	if err != nil {
		panic(err)
	}
	return e
}

// Parameters implements the kinit.Executor interface.
func (e *Executor) Parameters() []reflect.Type {
	if e == nil {
		return nil
	}
	types := make([]reflect.Type, len(e.inTypes))
	copy(types, e.inTypes)
	return types
}

// Execute implements the kinit.Executor interface.
func (e *Executor) Execute(a ...reflect.Value) (kinit.Executor, error) {
	if e == nil {
		return nil, nil
	}
	if len(a) != len(e.inTypes) {
		return nil, kerror.Newf(kerror.ERuntime,
			"executor expects %d argument(s), %d given", len(e.inTypes), len(a))
	}
	for i, v := range a {
		if v.Type() != e.inTypes[i] {
			return nil, kerror.Newf(kerror.ERuntime,
				"executor expects argument %d to be of %s type, %s given",
				i+1, e.inTypes[i], v.Type())
		}
	}
	out := e.function.Call(a)
	var exec kinit.Executor
	if e.executorOutIndex >= 0 {
		if v := out[e.executorOutIndex].Interface(); v != nil {
			exec = v.(kinit.Executor)
		}
	}
	var err error
	if e.errorOutIndex >= 0 {
		if v := out[e.errorOutIndex].Interface(); v != nil {
			err = v.(error)
		}
	}
	return exec, err
}

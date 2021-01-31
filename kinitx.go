// Package kinitx provides extensions for the package github.com/go-kata/kinit.
package kinitx

import (
	"io"
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kinit"
)

// errorType specifies the reflection to the error interface.
var errorType = reflect.TypeOf((*error)(nil)).Elem()

// closerType specifies the reflection to the io.Closer interface.
var closerType = reflect.TypeOf((*io.Closer)(nil)).Elem()

// destructorType specifies the reflection to the kdone.Destructor interface.
var destructorType = reflect.TypeOf((*kdone.Destructor)(nil)).Elem()

// executorType specifies the reflection to the kinit.Executor interface.
var executorType = reflect.TypeOf((*kinit.Executor)(nil)).Elem()

// Provide calls the kinit.Provide passing a constructor based on the given entity.
//
// See the documentation for the MakeConstructor to find out possible values of the argument x.
func Provide(x interface{}) error {
	ctor, err := MakeConstructor(x)
	if err != nil {
		return err
	}
	return kinit.Provide(ctor)
}

// MustProvide is a variant of the Provide that panics on error.
func MustProvide(x interface{}) {
	if err := Provide(x); err != nil {
		panic(err)
	}
}

// Apply calls kinit.Apply passing a processor based on the given entity.
//
// See the documentation for the NewProcessor to find out possible values of the argument x.
func Apply(x interface{}) error {
	proc, err := NewProcessor(x)
	if err != nil {
		return err
	}
	return kinit.Apply(proc)
}

// MustApply is a variant of the Apply that panics on error.
func MustApply(x interface{}) {
	if err := Apply(x); err != nil {
		panic(err)
	}
}

// Invoke calls the kinit.Invoke passing an executor and initializers based on given entities.
//
// See the documentation for the NewExecutor and NewLiteral to find out possible values
// of the arguments x and xx, respectively.
func Invoke(x interface{}, xx ...interface{}) error {
	exec, err := NewExecutor(x)
	if err != nil {
		return err
	}
	bootstrappers := make([]kinit.Bootstrapper, len(xx))
	for i, v := range xx {
		boot, err := NewLiteral(v)
		if err != nil {
			return err
		}
		bootstrappers[i] = boot
	}
	return kinit.Invoke(exec, bootstrappers...)
}

// MustInvoke is a variant of the Invoke that panics on error.
func MustInvoke(x interface{}, xx ...interface{}) {
	if err := Invoke(x, xx...); err != nil {
		panic(err)
	}
}

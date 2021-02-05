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

// functorType specifies the reflection to the kinit.Functor interface.
var functorType = reflect.TypeOf((*kinit.Functor)(nil)).Elem()

// functorSliceType specifies the reflection to the slice of functors.
var functorSliceType = reflect.SliceOf(functorType)

// runtimeType specifies the reflection to the kinit.Runtime interface.
var runtimeType = reflect.TypeOf((*kinit.Runtime)(nil))

// Provide calls the kinit.Provide passing a constructor based on the given entity.
//
// The x argument will be parsed corresponding to following rules:
//
// If x is a function it will be parsed using the NewOpener only when returns
// an implementation of the io.Closer interface at the first position and, optionally,
// error at the second. All other functions will be parsed using the NewConstructor.
//
// If x is a struct or pointer it will be parsed using the NewInitializer.
//
// All other variants of x are unacceptable.
func Provide(x interface{}) error {
	ctor, err := newProperConstructor(x)
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

// Attach calls kinit.Attach passing a processor based on the given entity.
//
// See the documentation for the NewProcessor to find out possible values of the argument x.
func Attach(x interface{}) error {
	proc, err := NewProcessor(x)
	if err != nil {
		return err
	}
	return kinit.Attach(proc)
}

// MustAttach is a variant of the Attach that panics on error.
func MustAttach(x interface{}) {
	if err := Attach(x); err != nil {
		panic(err)
	}
}

// Run calls the kinit.Run passing functors based on given entities.
//
// Items of the xx argument (let's name each item as x) will be parsed corresponding to following rules:
//
// It is unacceptable x to be nil.
//
// If x is a function it will be parsed using the NewFunctor.
//
// Otherwise x will be parsed using the NewInjector.
//
func Run(xx ...interface{}) error {
	functors := make([]kinit.Functor, len(xx))
	for i, x := range xx {
		fun, err := newProperFunctor(x)
		if err != nil {
			return err
		}
		functors[i] = fun
	}
	return kinit.Run(functors...)
}

// MustRun is a variant of the Run that panics on error.
func MustRun(xx ...interface{}) {
	if err := Run(xx...); err != nil {
		panic(err)
	}
}

package kinitx

import (
	"reflect"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

// MakeConstructor returns a new constructor based on the given entity.
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
func MakeConstructor(x interface{}) (kinit.Constructor, error) {
	if x == nil {
		return nil, kerror.New(kerror.ERuntime, "function, struct or struct pointer expected, nil given")
	}
	var ctor kinit.Constructor
	var err error
	t := reflect.TypeOf(x)
	switch t.Kind() {
	default:
		return nil, kerror.Newf(kerror.ERuntime, "function, struct or struct pointer expected, %s given", t)
	case reflect.Func:
		var isOpener bool
		switch t.NumOut() {
		case 2:
			if t.Out(1) != errorType {
				break
			}
			fallthrough
		case 1:
			isOpener = t.Out(0).Implements(closerType)
		}
		if isOpener {
			ctor, err = NewOpener(x)
		} else {
			ctor, err = NewConstructor(x)
		}
	case reflect.Struct, reflect.Ptr:
		ctor, err = NewInitializer(x)
	}
	return ctor, err
}

// MustMakeConstructor is a variant of the MakeConstructor that panics on error.
func MustMakeConstructor(x interface{}) kinit.Constructor {
	ctor, err := MakeConstructor(x)
	if err != nil {
		panic(err)
	}
	return ctor
}

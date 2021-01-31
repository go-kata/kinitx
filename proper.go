package kinitx

import (
	"reflect"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

// newProperConstructor returns a new constructor based on the given entity.
//
// See the documentation for the Provide to find out possible values of the argument x.
func newProperConstructor(x interface{}) (kinit.Constructor, error) {
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

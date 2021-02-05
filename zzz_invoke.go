package kinitx

import "github.com/go-kata/kinit"

// Invoke calls the kinit.Invoke passing an executor and initializers based on given entities.
//
// See the documentation for the NewExecutor and NewLiteral to find out possible values
// of the arguments x and xx, respectively.
//
// Deprecated: since 0.4.0, use Run instead.
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
//
// Deprecated: since 0.4.0, use MustRun instead.
func MustInvoke(x interface{}, xx ...interface{}) {
	if err := Invoke(x, xx...); err != nil {
		panic(err)
	}
}

package kinitx

import "github.com/go-kata/kinit"

// Apply calls kinit.Apply passing a processor based on the given entity.
//
// See the documentation for the NewProcessor to find out possible values of the argument x.
//
// Deprecated: since 0.4.0, use Attach instead.
func Apply(x interface{}) error {
	proc, err := NewProcessor(x)
	if err != nil {
		return err
	}
	return kinit.Apply(proc)
}

// MustApply is a variant of the Apply that panics on error.
//
// Deprecated: since 0.4.0, use MustAttach instead.
func MustApply(x interface{}) {
	if err := Apply(x); err != nil {
		panic(err)
	}
}

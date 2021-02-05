package kinitx

import (
	"testing"

	"github.com/go-kata/kerror"
)

func TestApply__NilProcessor(t *testing.T) {
	err := Apply(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestMustApply__NilProcessor(t *testing.T) {
	err := kerror.Try(func() error {
		MustApply(nil)
		return nil
	})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

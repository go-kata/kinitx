package kinitx

import (
	"testing"

	"github.com/go-kata/kerror"
)

func TestProvide__NilConstructor(t *testing.T) {
	err := Provide(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestMustProvide__NilConstructor(t *testing.T) {
	err := kerror.Try(func() error {
		MustProvide(nil)
		return nil
	})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestAttach__NilProcessor(t *testing.T) {
	err := Attach(nil)
	t.Logf("%+v", err)
	if err == nil {
		t.Fail()
		return
	}
}

func TestMustAttach__NilProcessor(t *testing.T) {
	err := kerror.Try(func() error {
		MustAttach(nil)
		return nil
	})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestRun__NilFunctor(t *testing.T) {
	err := Run(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestMustRun__NilFunctor(t *testing.T) {
	err := kerror.Try(func() error {
		MustRun(nil)
		return nil
	})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

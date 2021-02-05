package kinitx

import (
	"testing"

	"github.com/go-kata/kerror"
)

func TestInvoke__NilExecutor(t *testing.T) {
	err := Invoke(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestMustInvoke__NilExecutor(t *testing.T) {
	err := kerror.Try(func() error {
		MustInvoke(nil)
		return nil
	})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestInvoke__NilBootstrapper(t *testing.T) {
	err := Invoke(func() {}, nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestMustInvoke__NilBootstrapper(t *testing.T) {
	err := kerror.Try(func() error {
		MustInvoke(func() {}, nil)
		return nil
	})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

package kinitx

import (
	"reflect"
	"testing"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

func TestExecutor0(t *testing.T) {
	var c int
	exec := MustNewExecutor(func(v *int) { *v++ })
	t.Logf("%+v", exec.Parameters())
	if _, err := exec.Execute(reflect.ValueOf(&c)); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if c != 1 {
		t.Fail()
		return
	}
}

func TestExecutor1(t *testing.T) {
	var c int
	exec := MustNewExecutor(func(v *int) error {
		*v++
		return nil
	})
	t.Logf("%+v", exec.Parameters())
	if _, err := exec.Execute(reflect.ValueOf(&c)); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if c != 1 {
		t.Fail()
		return
	}
}

func TestExecutor2(t *testing.T) {
	var c int
	exec := MustNewExecutor(func(v *int) (kinit.Executor, error) {
		*v++
		return MustNewExecutor(func() {}), nil
	})
	t.Logf("%+v", exec.Parameters())
	if _, err := exec.Execute(reflect.ValueOf(&c)); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if c != 1 {
		t.Fail()
		return
	}
}

func TestExecutorWithWrongXType(t *testing.T) {
	_, err := NewExecutor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutorWithWrongSignature(t *testing.T) {
	_, err := NewExecutor(func() int { return 0 })
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutorWithWrongNumberOfArguments(t *testing.T) {
	var c int
	exec := MustNewExecutor(func(v *int) { *v++ })
	t.Logf("%+v", exec.Parameters())
	_, err := exec.Execute(reflect.ValueOf(&c), reflect.ValueOf(0))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutorWithWrongArgumentType(t *testing.T) {
	exec := MustNewExecutor(func(v *int) { *v++ })
	t.Logf("%+v", exec.Parameters())
	_, err := exec.Execute(reflect.ValueOf(""))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

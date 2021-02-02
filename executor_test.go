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

func TestExecutor_NewWithNil(t *testing.T) {
	_, err := NewExecutor(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutor_NewWithNilFunction(t *testing.T) {
	_, err := NewExecutor((func())(nil))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutor_NewWithWrongType(t *testing.T) {
	_, err := NewExecutor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutor_NewWithWrongSignature(t *testing.T) {
	_, err := NewExecutor(func() int { return 0 })
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestExecutor_ExecuteWithWrongNumberOfArguments(t *testing.T) {
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

func TestExecutor_ExecuteWithWrongArgumentType(t *testing.T) {
	exec := MustNewExecutor(func(v *int) { *v++ })
	t.Logf("%+v", exec.Parameters())
	_, err := exec.Execute(reflect.ValueOf(""))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestNilExecutor_Parameters(t *testing.T) {
	if (*Executor)(nil).Parameters() != nil {
		t.Fail()
		return
	}
}

func TestNilExecutor_Execute(t *testing.T) {
	exec, err := (*Executor)(nil).Execute()
	if exec != nil {
		t.Fail()
		return
	}
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
}

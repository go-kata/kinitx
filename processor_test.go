package kinitx

import (
	"reflect"
	"testing"

	"github.com/go-kata/kerror"
)

func TestProcessor0(t *testing.T) {
	var c int
	proc := MustNewProcessor(func(v *int, i int8) { *v += int(i) })
	t.Logf("%+v %+v", proc.Type(), proc.Parameters())
	if err := proc.Process(reflect.ValueOf(&c), reflect.ValueOf(int8(1))); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if c != 1 {
		t.Fail()
		return
	}
}

func TestProcessor1(t *testing.T) {
	var c int
	proc := MustNewProcessor(func(v *int, i int8) error {
		*v += int(i)
		return nil
	})
	t.Logf("%+v %+v", proc.Type(), proc.Parameters())
	if err := proc.Process(reflect.ValueOf(&c), reflect.ValueOf(int8(1))); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if c != 1 {
		t.Fail()
		return
	}
}

func TestProcessor_NewWithNil(t *testing.T) {
	_, err := NewProcessor(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessor_NewWithNilFunction(t *testing.T) {
	_, err := NewProcessor((func())(nil))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessor_NewWithWrongType(t *testing.T) {
	_, err := NewProcessor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessor_NewWithWrongSignature(t *testing.T) {
	_, err := NewProcessor(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessor_ProcessWithWrongObjectType(t *testing.T) {
	proc := MustNewProcessor(func(v *int, i int8) { *v += int(i) })
	t.Logf("%+v %+v", proc.Type(), proc.Parameters())
	err := proc.Process(reflect.ValueOf(""), reflect.ValueOf(int8(1)))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessor_ProcessWithWrongNumberOfArguments(t *testing.T) {
	var c int
	proc := MustNewProcessor(func(v *int, i int8) { *v += int(i) })
	t.Logf("%+v %+v", proc.Type(), proc.Parameters())
	err := proc.Process(reflect.ValueOf(&c))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessor_ProcessWithWrongArgumentType(t *testing.T) {
	var c int
	proc := MustNewProcessor(func(v *int, i int8) { *v += int(i) })
	t.Logf("%+v %+v", proc.Type(), proc.Parameters())
	err := proc.Process(reflect.ValueOf(&c), reflect.ValueOf(""))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestNilProcessor_Type(t *testing.T) {
	if (*Processor)(nil).Type() != nil {
		t.Fail()
		return
	}
}

func TestNilProcessor_Parameters(t *testing.T) {
	if (*Processor)(nil).Parameters() != nil {
		t.Fail()
		return
	}
}

func TestNilProcessor_Process(t *testing.T) {
	if err := (*Processor)(nil).Process(reflect.Value{}); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
}

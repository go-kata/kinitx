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

func TestProcessorWithWrongXType(t *testing.T) {
	_, err := NewProcessor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessorWithWrongSignature(t *testing.T) {
	_, err := NewProcessor(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessorWithWrongObjectType(t *testing.T) {
	proc := MustNewProcessor(func(v *int, i int8) { *v += int(i) })
	t.Logf("%+v %+v", proc.Type(), proc.Parameters())
	err := proc.Process(reflect.ValueOf(""), reflect.ValueOf(int8(1)))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestProcessorWithWrongNumberOfArguments(t *testing.T) {
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

func TestProcessorWithWrongArgumentType(t *testing.T) {
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

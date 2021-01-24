package kinitx

import (
	"reflect"
	"testing"

	"github.com/go-kata/kerror"
)

type testStructT1 struct{}

type testStructT2 struct{}

type testStructT3 struct {
	Obj1 *testStructT1
	Obj2 *testStructT2
}

func TestStructWithStruct(t *testing.T) {
	ctor := MustNewStruct(testStructT3{})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testStructT1{}
	obj2 := &testStructT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(testStructT3)
	if !ok {
		t.Logf("%+v", o3)
		t.Fail()
		return
	}
	if obj3.Obj1 != obj1 || obj3.Obj2 != obj2 {
		t.Fail()
		return
	}
}

func TestStructWithPointer(t *testing.T) {
	ctor := MustNewStruct((*testStructT3)(nil))
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testStructT1{}
	obj2 := &testStructT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(*testStructT3)
	if !ok {
		t.Logf("%+v", o3)
		t.Fail()
		return
	}
	if obj3.Obj1 != obj1 || obj3.Obj2 != obj2 {
		t.Fail()
		return
	}
}

func TestStructWithWrongXType(t *testing.T) {
	_, err := NewStruct(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestStructWithWrongArgumentNumber(t *testing.T) {
	ctor := MustNewStruct((*testStructT3)(nil))
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	_, _, err := ctor.Create(reflect.ValueOf(&testStructT1{}))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestStructWithWrongArgumentType(t *testing.T) {
	ctor := MustNewStruct((*testStructT3)(nil))
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	_, _, err := ctor.Create(reflect.ValueOf(&testStructT1{}), reflect.ValueOf(0))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

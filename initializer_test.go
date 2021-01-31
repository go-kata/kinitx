package kinitx

import (
	"reflect"
	"testing"

	"github.com/go-kata/kerror"
)

type testInitializerT1 struct{}

type testInitializerT2 struct{}

type testInitializerT3 struct {
	Obj1 *testInitializerT1
	Obj2 *testInitializerT2
}

func TestInitializerWithStruct(t *testing.T) {
	ctor := MustNewInitializer(testInitializerT3{})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testInitializerT1{}
	obj2 := &testInitializerT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(testInitializerT3)
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

func TestInitializerWithStructPointer(t *testing.T) {
	ctor := MustNewInitializer((*testInitializerT3)(nil))
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testInitializerT1{}
	obj2 := &testInitializerT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(*testInitializerT3)
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

func TestInitializerWithWrongXType(t *testing.T) {
	_, err := NewInitializer(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestInitializerWithWrongArgumentNumber(t *testing.T) {
	ctor := MustNewInitializer((*testInitializerT3)(nil))
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	_, _, err := ctor.Create(reflect.ValueOf(&testInitializerT1{}))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestInitializerWithWrongArgumentType(t *testing.T) {
	ctor := MustNewInitializer((*testInitializerT3)(nil))
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	_, _, err := ctor.Create(reflect.ValueOf(&testInitializerT1{}), reflect.ValueOf(0))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

package kinitx

import (
	"reflect"
	"testing"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
)

type testFactoryT1 struct{}

type testFactoryT2 struct{}

type testFactoryT3 struct {
	obj1 *testFactoryT1
	obj2 *testFactoryT2
}

func TestFactory1(t *testing.T) {
	ctor := MustNewFactory(func(
		obj1 *testFactoryT1,
		obj2 *testFactoryT2,
	) *testFactoryT3 {
		return &testFactoryT3{obj1, obj2}
	})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testFactoryT1{}
	obj2 := &testFactoryT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(*testFactoryT3)
	if !ok {
		t.Logf("%+v", o3)
		t.Fail()
		return
	}
	if obj3.obj1 != obj1 || obj3.obj2 != obj2 {
		t.Fail()
		return
	}
}

func TestFactory2(t *testing.T) {
	ctor := MustNewFactory(func(
		obj1 *testFactoryT1,
		obj2 *testFactoryT2,
	) (
		*testFactoryT3,
		error,
	) {
		return &testFactoryT3{obj1, obj2}, nil
	})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testFactoryT1{}
	obj2 := &testFactoryT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(*testFactoryT3)
	if !ok {
		t.Logf("%+v", o3)
		t.Fail()
		return
	}
	if obj3.obj1 != obj1 || obj3.obj2 != obj2 {
		t.Fail()
		return
	}
}

func TestFactory3(t *testing.T) {
	ctor := MustNewFactory(func(
		obj1 *testFactoryT1,
		obj2 *testFactoryT2,
	) (
		*testFactoryT3,
		kdone.Destructor,
		error,
	) {
		return &testFactoryT3{obj1, obj2}, kdone.Noop, nil
	})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	obj1 := &testFactoryT1{}
	obj2 := &testFactoryT2{}
	o3, dtor, err := ctor.Create(reflect.ValueOf(obj1), reflect.ValueOf(obj2))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	defer dtor.MustDestroy()
	obj3, ok := o3.Interface().(*testFactoryT3)
	if !ok {
		t.Logf("%+v", o3)
		t.Fail()
		return
	}
	if obj3.obj1 != obj1 || obj3.obj2 != obj2 {
		t.Fail()
		return
	}
}

func TestFactoryWithWrongXType(t *testing.T) {
	_, err := NewFactory(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestFactoryWithWrongSignature(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Logf("%+v", v)
			t.Fail()
			return
		}
	}()
	_, err := NewFactory(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestFactoryWithWrongNumberOfArguments(t *testing.T) {
	ctor := MustNewFactory(func(
		obj1 *testFactoryT1,
		obj2 *testFactoryT2,
	) *testFactoryT3 {
		return &testFactoryT3{obj1, obj2}
	})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	_, _, err := ctor.Create(reflect.ValueOf(&testFactoryT1{}))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestFactoryWithWrongArgumentType(t *testing.T) {
	ctor := MustNewFactory(func(
		obj1 *testFactoryT1,
		obj2 *testFactoryT2,
	) *testFactoryT3 {
		return &testFactoryT3{obj1, obj2}
	})
	t.Logf("%+v %+v", ctor.Type(), ctor.Parameters())
	_, _, err := ctor.Create(reflect.ValueOf(&testFactoryT1{}), reflect.ValueOf(0))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

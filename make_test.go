package kinitx

import (
	"testing"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
)

type testMakeT struct{}

func (*testMakeT) Close() error {
	return nil
}

func TestMakeConstructorWithConstructor(t *testing.T) {
	ctor := MustMakeConstructor(func() int { return 0 })
	if _, ok := ctor.(*Constructor); !ok {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithOpener(t *testing.T) {
	ctor := MustMakeConstructor(func() *testMakeT { return &testMakeT{} })
	if _, ok := ctor.(*Opener); !ok {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithConstructorOfCloser(t *testing.T) {
	ctor := MustMakeConstructor(func() (*testMakeT, kdone.Destructor, error) {
		return &testMakeT{}, kdone.Noop, nil
	})
	if _, ok := ctor.(*Constructor); !ok {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithStruct(t *testing.T) {
	ctor := MustMakeConstructor(testMakeT{})
	if _, ok := ctor.(*Initializer); !ok {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithStructPointer(t *testing.T) {
	ctor := MustMakeConstructor((*testMakeT)(nil))
	if _, ok := ctor.(*Initializer); !ok {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithWrongX(t *testing.T) {
	_, err := MakeConstructor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithWrongFunc(t *testing.T) {
	_, err := MakeConstructor(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestMakeConstructorWithWrongPointer(t *testing.T) {
	_, err := MakeConstructor((*int)(nil))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

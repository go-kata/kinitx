package kinitx

import (
	"testing"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
)

type testCloser struct{}

func (*testCloser) Close() error {
	return nil
}

func TestNewProperConstructorWithConstructor(t *testing.T) {
	ctor, err := newProperConstructor(func() int { return 0 })
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := ctor.(*Constructor); !ok {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithOpener1(t *testing.T) {
	ctor, err := newProperConstructor(func() *testCloser { return &testCloser{} })
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := ctor.(*Opener); !ok {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithOpener2(t *testing.T) {
	ctor, err := newProperConstructor(func() (*testCloser, error) { return &testCloser{}, nil })
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := ctor.(*Opener); !ok {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithConstructorOfCloser(t *testing.T) {
	ctor, err := newProperConstructor(func() (*testCloser, kdone.Destructor, error) {
		return &testCloser{}, kdone.Noop, nil
	})
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := ctor.(*Constructor); !ok {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithStruct(t *testing.T) {
	ctor, err := newProperConstructor(testCloser{})
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := ctor.(*Initializer); !ok {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithStructPointer(t *testing.T) {
	ctor, err := newProperConstructor((*testCloser)(nil))
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := ctor.(*Initializer); !ok {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithWrongType(t *testing.T) {
	_, err := newProperConstructor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithWrongFunc(t *testing.T) {
	_, err := newProperConstructor(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestNewProperConstructorWithWrongPointer(t *testing.T) {
	_, err := newProperConstructor((*int)(nil))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

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

func TestNewProperConstructor__Constructor(t *testing.T) {
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

func TestNewProperConstructor__Opener(t *testing.T) {
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

func TestNewProperConstructor__ErrorProneOpener(t *testing.T) {
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

func TestNewProperConstructor__ConstructorOfCloser(t *testing.T) {
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

func TestNewProperConstructor__Struct(t *testing.T) {
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

func TestNewProperConstructor__StructPointer(t *testing.T) {
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

func TestNewProperConstructor__Nil(t *testing.T) {
	_, err := newProperConstructor(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestNewProperConstructor__WrongType(t *testing.T) {
	_, err := newProperConstructor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestNewProperConstructor__WrongFunc(t *testing.T) {
	_, err := newProperConstructor(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestNewProperConstructor__WrongPointer(t *testing.T) {
	_, err := newProperConstructor((*int)(nil))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

func TestNewProperFunctor__Functor(t *testing.T) {
	fun, err := newProperFunctor(func() {})
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := fun.(*Functor); !ok {
		t.Fail()
		return
	}
}

func TestNewProperFunctor__Injector(t *testing.T) {
	fun, err := newProperFunctor(1)
	if err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if _, ok := fun.(*Injector); !ok {
		t.Fail()
		return
	}
}

func TestNewProperFunctor__Nil(t *testing.T) {
	_, err := newProperFunctor(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.EViolation {
		t.Fail()
		return
	}
}

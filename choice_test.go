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

func TestChooseConstructorWithConstructor(t *testing.T) {
	ctor, err := chooseConstructor(func() int { return 0 })
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

func TestChooseConstructorWithOpener(t *testing.T) {
	ctor, err := chooseConstructor(func() *testCloser { return &testCloser{} })
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

func TestChooseConstructorWithConstructorOfCloser(t *testing.T) {
	ctor, err := chooseConstructor(func() (*testCloser, kdone.Destructor, error) {
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

func TestChooseConstructorWithStruct(t *testing.T) {
	ctor, err := chooseConstructor(testCloser{})
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

func TestChooseConstructorWithStructPointer(t *testing.T) {
	ctor, err := chooseConstructor((*testCloser)(nil))
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

func TestChooseConstructorWithWrongX(t *testing.T) {
	_, err := chooseConstructor(0)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestChooseConstructorWithWrongFunc(t *testing.T) {
	_, err := chooseConstructor(func() {})
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

func TestChooseConstructorWithWrongPointer(t *testing.T) {
	_, err := chooseConstructor((*int)(nil))
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

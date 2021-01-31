package kinitx

import (
	"reflect"
	"testing"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
)

func TestLiteral(t *testing.T) {
	x := 1
	boot := MustNewLiteral(x)
	arena := kinit.NewArena()
	defer arena.MustFinalize()
	if err := boot.Bootstrap(arena); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
	if obj, ok := arena.Get(reflect.TypeOf(x)); !ok || obj.Interface() != x {
		t.Fail()
		return
	}
}

func TestLiteralWithNilX(t *testing.T) {
	_, err := NewLiteral(nil)
	t.Logf("%+v", err)
	if kerror.ClassOf(err) != kerror.ERuntime {
		t.Fail()
		return
	}
}

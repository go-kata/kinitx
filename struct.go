package kinitx

import (
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
)

// Struct represents a constructor based on a struct.
type Struct struct {
	// t specifies the type of an object that is created by this fusion.
	t reflect.Type
	// assignableFieldTypes specifies types of assignable struct fields.
	assignableFieldTypes []reflect.Type
	// assignableFieldIndexes specifies indexes of assignable struct fields.
	assignableFieldIndexes []int
}

// NewStruct returns a new constructor.
//
// The argument x must be a struct or a struct pointer.
func NewStruct(x interface{}) (*Struct, error) {
	if x == nil {
		return nil, kerror.New(kerror.ERuntime, "struct or struct pointer expected, nil given")
	}
	t := reflect.TypeOf(x)
	var st reflect.Type
	switch t.Kind() {
	default:
		return nil, kerror.Newf(kerror.ERuntime, "struct or struct pointer expected, %s given", t)
	case reflect.Struct:
		st = t
	case reflect.Ptr:
		st = t.Elem()
		if st.Kind() != reflect.Struct {
			return nil, kerror.Newf(kerror.ERuntime, "struct or struct pointer expected, %s given", t)
		}
	}
	c := &Struct{
		t: t,
	}
	for i, n := 0, st.NumField(); i < n; i++ {
		sf := st.Field(i)
		if sf.PkgPath != "" {
			continue
		}
		c.assignableFieldTypes = append(c.assignableFieldTypes, sf.Type)
		c.assignableFieldIndexes = append(c.assignableFieldIndexes, i)
	}
	return c, nil
}

// MustNewStruct is a variant of the NewStruct that panics on error.
func MustNewStruct(x interface{}) *Struct {
	c, err := NewStruct(x)
	if err != nil {
		panic(err)
	}
	return c
}

// Type implements the kinit.Constructor interface.
func (c *Struct) Type() reflect.Type {
	if c == nil {
		return nil
	}
	return c.t
}

// Parameters implements the kinit.Constructor interface.
func (c *Struct) Parameters() []reflect.Type {
	if c == nil {
		return nil
	}
	types := make([]reflect.Type, len(c.assignableFieldTypes))
	copy(types, c.assignableFieldTypes)
	return types
}

// Create implements the kinit.Constructor interface.
func (c *Struct) Create(a ...reflect.Value) (reflect.Value, kdone.Destructor, error) {
	if c == nil {
		return reflect.Value{}, kdone.Noop, nil
	}
	if len(a) != len(c.assignableFieldTypes) {
		return reflect.Value{}, kdone.Noop, kerror.Newf(kerror.ERuntime,
			"%s constructor expects %d argument(s), %d given",
			c.t, len(c.assignableFieldTypes), len(a))
	}
	for i, v := range a {
		if v.Type() != c.assignableFieldTypes[i] {
			return reflect.Value{}, kdone.Noop, kerror.Newf(kerror.ERuntime,
				"%s constructor expects argument %d to be of %s type, %s given",
				c.t, i+1, c.assignableFieldTypes[i], v.Type())
		}
	}
	var sp, obj reflect.Value
	if c.t.Kind() == reflect.Ptr {
		sp = reflect.New(c.t.Elem())
		obj = sp
	} else {
		sp = reflect.New(c.t)
		obj = sp.Elem()
	}
	sv := sp.Elem()
	for i, v := range a {
		sv.Field(c.assignableFieldIndexes[i]).Set(v)
	}
	return obj, kdone.Noop, nil
}

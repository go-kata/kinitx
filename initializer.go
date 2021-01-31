package kinitx

import (
	"reflect"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kerror"
)

// Initializer represents a constructor based on a struct.
type Initializer struct {
	// t specifies the type of an object that is created by this initializer.
	t reflect.Type
	// assignableFieldTypes specifies types of assignable struct fields.
	assignableFieldTypes []reflect.Type
	// assignableFieldIndexes specifies indexes of assignable struct fields.
	assignableFieldIndexes []int
}

// NewInitializer returns a new initializer.
//
// The argument x must be a struct or a struct pointer.
func NewInitializer(x interface{}) (*Initializer, error) {
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
	c := &Initializer{
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

// MustNewInitializer is a variant of the NewInitializer that panics on error.
func MustNewInitializer(x interface{}) *Initializer {
	c, err := NewInitializer(x)
	if err != nil {
		panic(err)
	}
	return c
}

// Type implements the kinit.Constructor interface.
func (c *Initializer) Type() reflect.Type {
	if c == nil {
		return nil
	}
	return c.t
}

// Parameters implements the kinit.Constructor interface.
func (c *Initializer) Parameters() []reflect.Type {
	if c == nil {
		return nil
	}
	types := make([]reflect.Type, len(c.assignableFieldTypes))
	copy(types, c.assignableFieldTypes)
	return types
}

// Create implements the kinit.Constructor interface.
func (c *Initializer) Create(a ...reflect.Value) (reflect.Value, kdone.Destructor, error) {
	if c == nil {
		return reflect.Value{}, kdone.Noop, nil
	}
	if len(a) != len(c.assignableFieldTypes) {
		return reflect.Value{}, kdone.Noop, kerror.Newf(kerror.ERuntime,
			"%s initializer expects %d argument(s), %d given",
			c.t, len(c.assignableFieldTypes), len(a))
	}
	for i, v := range a {
		if v.Type() != c.assignableFieldTypes[i] {
			return reflect.Value{}, kdone.Noop, kerror.Newf(kerror.ERuntime,
				"%s initializer expects argument %d to be of %s type, %s given",
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

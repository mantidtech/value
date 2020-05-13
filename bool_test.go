// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptrToBool(a bool) *Bool {
	x := new(Bool)
	*x = Bool(a)
	return x
}

func boolPtr(a bool) *bool {
	x := new(bool)
	*x = a
	return x
}

func TestBool_Type(t *testing.T) {
	a := new(Bool)
	assert.Equal(t, "bool", a.Type())
}

func TestBool_Zero(t *testing.T) {
	a := new(bool)
	b := new(Bool)
	assert.Equal(t, *a, b.Zero())
}

func TestBool_Clone(t *testing.T) {
	a := new(Bool)
	b := a.Clone()
	assert.Equal(t, bool(*a), b.Get())
	*a = Bool(true)
	assert.NotEqual(t, bool(*a), b.Get())
}

type TestBoolGetTest struct {
	name string
	a    *Bool
	want interface{}
}

func TestBool_Get(t *testing.T) {
	for _, st := range testBoolGetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBoolSetTest struct {
	name    string
	s       string
	a       *Bool
	wantErr bool
}

func TestBool_Set(t *testing.T) {
	for _, st := range testBoolSetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.a.Set(tt.s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

type TestBoolStringTest struct {
	name string
	a    *Bool
	want string
}

func TestBool_String(t *testing.T) {
	for _, st := range testBoolStringTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBindBoolTest struct {
	name string
	x    *bool
	want *Bool
}

func TestBindBool(t *testing.T) {
	for _, st := range testBindBoolTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BindBool(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestFactoryBind_Bool(t *testing.T) {
	a := bool(true)
	got := Bind(&a)
	err := got.Set("false")
	require.NoError(t, err)
	assert.NotEqual(t, bool(true), got.Get())
	assert.Equal(t, bool(false), got.Get())
}
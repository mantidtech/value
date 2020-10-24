// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptrToInt(a int) *Int {
	x := new(Int)
	*x = Int(a)
	return x
}

func intPtr(a int) *int {
	x := new(int)
	*x = a
	return x
}

func TestInt_Type(t *testing.T) {
	a := new(Int)
	assert.Equal(t, "int", a.Type())
}

func TestInt_Zero(t *testing.T) {
	a := new(int)
	b := new(Int)
	assert.Equal(t, *a, b.Zero())
}

func TestInt_Clone(t *testing.T) {
	a := new(Int)
	b := a.Clone()
	assert.Equal(t, int(*a), b.Get())
	*a = Int(14)
	assert.NotEqual(t, int(*a), b.Get())
}

type TestIntGetTest struct {
	name string
	a    *Int
	want interface{}
}

func TestInt_Get(t *testing.T) {
	for _, st := range testIntGetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestIntSetTest struct {
	name    string
	s       string
	a       *Int
	wantErr bool
}

func TestInt_Set(t *testing.T) {
	for _, st := range testIntSetTestData {
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

type TestIntStringTest struct {
	name string
	a    *Int
	want string
}

func TestInt_String(t *testing.T) {
	for _, st := range testIntStringTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBindIntTest struct {
	name string
	x    *int
	want *Int
}

func TestBindInt(t *testing.T) {
	for _, st := range testBindIntTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BindInt(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFactoryBind_Int(t *testing.T) {
	a := int(14)
	got := Bind(&a)
	err := got.Set("2147483647")
	require.NoError(t, err)
	assert.NotEqual(t, int(14), got.Get())
	assert.Equal(t, int((2<<30)-1), got.Get())
}

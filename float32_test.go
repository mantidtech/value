// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptrToFloat32(a float32) *Float32 {
	x := new(Float32)
	*x = Float32(a)
	return x
}

func float32Ptr(a float32) *float32 {
	x := new(float32)
	*x = a
	return x
}

func TestFloat32_Type(t *testing.T) {
	a := new(Float32)
	assert.Equal(t, "float32", a.Type())
}

func TestFloat32_Zero(t *testing.T) {
	a := new(float32)
	b := new(Float32)
	assert.Equal(t, *a, b.Zero())
}

func TestFloat32_Clone(t *testing.T) {
	a := new(Float32)
	b := a.Clone()
	assert.Equal(t, float32(*a), b.Get())
	*a = Float32(4.12890625)
	assert.NotEqual(t, float32(*a), b.Get())
}

type TestFloat32GetTest struct {
	name string
	a    *Float32
	want interface{}
}

func TestFloat32_Get(t *testing.T) {
	for _, st := range testFloat32GetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestFloat32SetTest struct {
	name    string
	s       string
	a       *Float32
	wantErr bool
}

func TestFloat32_Set(t *testing.T) {
	for _, st := range testFloat32SetTestData {
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

type TestFloat32StringTest struct {
	name string
	a    *Float32
	want string
}

func TestFloat32_String(t *testing.T) {
	for _, st := range testFloat32StringTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBindFloat32Test struct {
	name string
	x    *float32
	want *Float32
}

func TestBindFloat32(t *testing.T) {
	for _, st := range testBindFloat32TestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BindFloat32(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestFactoryBind_Float32(t *testing.T) {
	a := float32(4.12890625)
	got := Bind(&a)
	err := got.Set("-61724.0")
	require.NoError(t, err)
	assert.NotEqual(t, float32(4.12890625), got.Get())
	assert.Equal(t, float32(-61724.0), got.Get())
}

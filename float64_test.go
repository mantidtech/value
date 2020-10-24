// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptrToFloat64(a float64) *Float64 {
	x := new(Float64)
	*x = Float64(a)
	return x
}

func float64Ptr(a float64) *float64 {
	x := new(float64)
	*x = a
	return x
}

func TestFloat64_Type(t *testing.T) {
	a := new(Float64)
	assert.Equal(t, "float64", a.Type())
}

func TestFloat64_Zero(t *testing.T) {
	a := new(float64)
	b := new(Float64)
	assert.Equal(t, *a, b.Zero())
}

func TestFloat64_Clone(t *testing.T) {
	a := new(Float64)
	b := a.Clone()
	assert.Equal(t, float64(*a), b.Get())
	*a = Float64(5.5e32)
	assert.NotEqual(t, float64(*a), b.Get())
}

type TestFloat64GetTest struct {
	name string
	a    *Float64
	want interface{}
}

func TestFloat64_Get(t *testing.T) {
	for _, st := range testFloat64GetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestFloat64SetTest struct {
	name    string
	s       string
	a       *Float64
	wantErr bool
}

func TestFloat64_Set(t *testing.T) {
	for _, st := range testFloat64SetTestData {
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

type TestFloat64StringTest struct {
	name string
	a    *Float64
	want string
}

func TestFloat64_String(t *testing.T) {
	for _, st := range testFloat64StringTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBindFloat64Test struct {
	name string
	x    *float64
	want *Float64
}

func TestBindFloat64(t *testing.T) {
	for _, st := range testBindFloat64TestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BindFloat64(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFactoryBind_Float64(t *testing.T) {
	a := float64(5.5e32)
	got := Bind(&a)
	err := got.Set("-800.525")
	require.NoError(t, err)
	assert.NotEqual(t, float64(5.5e32), got.Get())
	assert.Equal(t, float64(-800.525), got.Get())
}

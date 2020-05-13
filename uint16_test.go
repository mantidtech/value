// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptrToUint16(a uint16) *Uint16 {
	x := new(Uint16)
	*x = Uint16(a)
	return x
}

func uint16Ptr(a uint16) *uint16 {
	x := new(uint16)
	*x = a
	return x
}

func TestUint16_Type(t *testing.T) {
	a := new(Uint16)
	assert.Equal(t, "uint16", a.Type())
}

func TestUint16_Zero(t *testing.T) {
	a := new(uint16)
	b := new(Uint16)
	assert.Equal(t, *a, b.Zero())
}

func TestUint16_Clone(t *testing.T) {
	a := new(Uint16)
	b := a.Clone()
	assert.Equal(t, uint16(*a), b.Get())
	*a = Uint16(50000)
	assert.NotEqual(t, uint16(*a), b.Get())
}

type TestUint16GetTest struct {
	name string
	a    *Uint16
	want interface{}
}

func TestUint16_Get(t *testing.T) {
	for _, st := range testUint16GetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestUint16SetTest struct {
	name    string
	s       string
	a       *Uint16
	wantErr bool
}

func TestUint16_Set(t *testing.T) {
	for _, st := range testUint16SetTestData {
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

type TestUint16StringTest struct {
	name string
	a    *Uint16
	want string
}

func TestUint16_String(t *testing.T) {
	for _, st := range testUint16StringTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBindUint16Test struct {
	name string
	x    *uint16
	want *Uint16
}

func TestBindUint16(t *testing.T) {
	for _, st := range testBindUint16TestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BindUint16(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestFactoryBind_Uint16(t *testing.T) {
	a := uint16(50000)
	got := Bind(&a)
	err := got.Set("603")
	require.NoError(t, err)
	assert.NotEqual(t, uint16(50000), got.Get())
	assert.Equal(t, uint16(603), got.Get())
}

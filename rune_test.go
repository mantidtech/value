// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptrToRune(a rune) *Rune {
	x := new(Rune)
	*x = Rune(a)
	return x
}

func runePtr(a rune) *rune {
	x := new(rune)
	*x = a
	return x
}

func TestRune_Type(t *testing.T) {
	a := new(Rune)
	assert.Equal(t, "rune", a.Type())
}

func TestRune_Zero(t *testing.T) {
	a := new(rune)
	b := new(Rune)
	assert.Equal(t, *a, b.Zero())
}

func TestRune_Clone(t *testing.T) {
	a := new(Rune)
	b := a.Clone()
	assert.Equal(t, rune(*a), b.Get())
	*a = Rune('J')
	assert.NotEqual(t, rune(*a), b.Get())
}

type TestRuneGetTest struct {
	name string
	a    *Rune
	want interface{}
}

func TestRune_Get(t *testing.T) {
	for _, st := range testRuneGetTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestRuneSetTest struct {
	name    string
	s       string
	a       *Rune
	wantErr bool
}

func TestRune_Set(t *testing.T) {
	for _, st := range testRuneSetTestData {
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

type TestRuneStringTest struct {
	name string
	a    *Rune
	want string
}

func TestRune_String(t *testing.T) {
	for _, st := range testRuneStringTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestBindRuneTest struct {
	name string
	x    *rune
	want *Rune
}

func TestBindRune(t *testing.T) {
	for _, st := range testBindRuneTestData {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BindRune(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

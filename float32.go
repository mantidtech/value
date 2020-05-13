// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"strconv"
)

// Float32 implements the value interface for the float32 type
type Float32 float32

// Type returns the underlying type of the value as a string
func (f *Float32) Type() string {
	return "float32"
}

// Get returns the underlying value as an interface{}
func (f *Float32) Get() interface{} {
	return float32(*f)
}

// Set a value from a string
func (f *Float32) Set(val string) error {
	x, err := strconv.ParseFloat(val, 32)
	r := float32(x)
	if err == nil {
		*f = Float32(r)
	}
	return err
}

// String returns the value as a string
func (f *Float32) String() string {
	return strconv.FormatFloat(float64(*f), 'f', -1, 32)
}

// Clone duplicates the Float32 as a copy of the instance
func (f *Float32) Clone() Value {
	ptr := new(Float32)
	*ptr = *f
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (f *Float32) Zero() interface{} {
	ptr := new(float32)
	return *ptr
}

// BindFloat32 returns a new Float32 with the underlying float32 used for state.
func BindFloat32(f *float32) *Float32 {
	return (*Float32)(f)
}
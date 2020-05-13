// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"strconv"
)

// Float64 implements the value interface for the float64 type
type Float64 float64

// Type returns the underlying type of the value as a string
func (f *Float64) Type() string {
	return "float64"
}

// Get returns the underlying value as an interface{}
func (f *Float64) Get() interface{} {
	return float64(*f)
}

// Set a value from a string
func (f *Float64) Set(val string) error {
	r, err := strconv.ParseFloat(val, 64)
	if err == nil {
		*f = Float64(r)
	}
	return err
}

// String returns the value as a string
func (f *Float64) String() string {
	return strconv.FormatFloat(float64(*f), 'f', -1, 64)
}

// Clone duplicates the Float64 as a copy of the instance
func (f *Float64) Clone() Value {
	ptr := new(Float64)
	*ptr = *f
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (f *Float64) Zero() interface{} {
	ptr := new(float64)
	return *ptr
}

// BindFloat64 returns a new Float64 with the underlying float64 used for state.
func BindFloat64(f *float64) *Float64 {
	return (*Float64)(f)
}

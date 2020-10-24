// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"strconv"
)

// Int16 implements the value interface for the int16 type
type Int16 int16

// Type returns the underlying type of the value as a string
func (i *Int16) Type() string {
	return "int16"
}

// Get returns the underlying value as an interface{}
func (i *Int16) Get() interface{} {
	return int16(*i)
}

// Set a value from a string
func (i *Int16) Set(val string) error {
	r, err := strconv.ParseInt(val, 0, 16)
	if err == nil {
		*i = Int16(r)
	}
	return err
}

// String returns the value as a string
func (i *Int16) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

// Clone duplicates the Int16 as a copy of the instance
func (i *Int16) Clone() Value {
	ptr := new(Int16)
	*ptr = *i
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (i *Int16) Zero() interface{} {
	ptr := new(int16)
	return *ptr
}

// BindInt16 returns a new Int16 with the underlying int16 used for state.
func BindInt16(i *int16) *Int16 {
	return (*Int16)(i)
}

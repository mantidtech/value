// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"strconv"
)

// Int32 implements the value interface for the int32 type
type Int32 int32

// Type returns the underlying type of the value as a string
func (i *Int32) Type() string {
	return "int32"
}

// Get returns the underlying value as an interface{}
func (i *Int32) Get() interface{} {
	return int32(*i)
}

// Set a value from a string
func (i *Int32) Set(val string) error {
	r, err := strconv.ParseInt(val, 0, 32)
	if err == nil {
		*i = Int32(r)
	}
	return err
}

// String returns the value as a string
func (i *Int32) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

// Clone duplicates the Int32 as a copy of the instance
func (i *Int32) Clone() Value {
	ptr := new(Int32)
	*ptr = *i
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (i *Int32) Zero() interface{} {
	ptr := new(int32)
	return *ptr
}

// BindInt32 returns a new Int32 with the underlying int32 used for state.
func BindInt32(i *int32) *Int32 {
	return (*Int32)(i)
}

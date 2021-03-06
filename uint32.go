// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"strconv"
)

// Uint32 implements the value interface for the uint32 type
type Uint32 uint32

// Type returns the underlying type of the value as a string
func (u *Uint32) Type() string {
	return "uint32"
}

// Get returns the underlying value as an interface{}
func (u *Uint32) Get() interface{} {
	return uint32(*u)
}

// Set a value from a string
func (u *Uint32) Set(val string) error {
	r, err := strconv.ParseUint(val, 0, 32)
	if err == nil {
		*u = Uint32(r)
	}
	return err
}

// String returns the value as a string
func (u *Uint32) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Clone duplicates the Uint32 as a copy of the instance
func (u *Uint32) Clone() Value {
	ptr := new(Uint32)
	*ptr = *u
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (u *Uint32) Zero() interface{} {
	ptr := new(uint32)
	return *ptr
}

// BindUint32 returns a new Uint32 with the underlying uint32 used for state.
func BindUint32(u *uint32) *Uint32 {
	return (*Uint32)(u)
}

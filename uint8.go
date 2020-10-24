// Code generated by valuegenerator. DO NOT EDIT

package value

import (
	"strconv"
)

// Uint8 implements the value interface for the uint8 type
type Uint8 uint8

// Type returns the underlying type of the value as a string
func (u *Uint8) Type() string {
	return "uint8"
}

// Get returns the underlying value as an interface{}
func (u *Uint8) Get() interface{} {
	return uint8(*u)
}

// Set a value from a string
func (u *Uint8) Set(val string) error {
	r, err := strconv.ParseUint(val, 0, 8)
	if err == nil {
		*u = Uint8(r)
	}
	return err
}

// String returns the value as a string
func (u *Uint8) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Clone duplicates the Uint8 as a copy of the instance
func (u *Uint8) Clone() Value {
	ptr := new(Uint8)
	*ptr = *u
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (u *Uint8) Zero() interface{} {
	ptr := new(uint8)
	return *ptr
}

// BindUint8 returns a new Uint8 with the underlying uint8 used for state.
func BindUint8(u *uint8) *Uint8 {
	return (*Uint8)(u)
}

// Code generated by valuegenerator. DO NOT EDIT

package value

// String implements the value interface for the string type
type String string

// Type returns the underlying type of the value as a string
func (s *String) Type() string {
	return "string"
}

// Get returns the underlying value as an interface{}
func (s *String) Get() interface{} {
	return string(*s)
}

// Set a value from a string
func (s *String) Set(val string) error {
	*s = String(val)
	return nil
}

// String returns the value as a string
func (s *String) String() string {
	return string(*s)
}

// Clone duplicates the String as a copy of the instance
func (s *String) Clone() Value {
	ptr := new(String)
	*ptr = *s
	return ptr
}

// Zero returns the zero value of the underlying primitive
func (s *String) Zero() interface{} {
	ptr := new(string)
	return *ptr
}

// BindString returns a new String with the underlying string used for state.
func BindString(s *string) *String {
	return (*String)(s)
}

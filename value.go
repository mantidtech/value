package value

// Value specifies the interface required for the defined types
type Value interface {
	Type() string
	Get() interface{}
	Set(string) error
	String() string
	Clone() Value
	Zero() interface{}
}

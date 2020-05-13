package value

// Value specifies the interface required for types that are to be used for flags
type Value interface {
	Type() string
	Get() interface{}
	Set(string) error
	String() string
	Clone() Value
	Zero() interface{}
}

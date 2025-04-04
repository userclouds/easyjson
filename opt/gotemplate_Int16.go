// generated by gotemplate

package opt

import (
	"fmt"

	"github.com/userclouds/easyjson/jlexer"
	"github.com/userclouds/easyjson/jwriter"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Int16 struct {
	V       int16
	Defined bool
}

// Creates an optional type with a given value.
func OInt16(v int16) Int16 {
	return Int16{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Int16) Get(deflt int16) int16 {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int16) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Int16(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int16) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int16{}
	} else {
		v.V = l.Int16()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int16) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int16) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Int16) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Int16) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}

package tests

import (
	"github.com/userclouds/easyjson"
	"github.com/userclouds/easyjson/jlexer"
	"github.com/userclouds/easyjson/jwriter"
)

//easyjson:json
type NestedMarshaler struct {
	Value easyjson.MarshalerUnmarshaler
	Value2 int
}

type StructWithMarshaler struct {
	Value int
}

func (s *StructWithMarshaler) UnmarshalEasyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}


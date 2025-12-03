// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Azure string       // Always "Azure"
type Comparative string // Always "comparative"
type Deepgram string    // Always "Deepgram"
type Simple string      // Always "simple"

func (c Azure) Default() Azure             { return "Azure" }
func (c Comparative) Default() Comparative { return "comparative" }
func (c Deepgram) Default() Deepgram       { return "Deepgram" }
func (c Simple) Default() Simple           { return "simple" }

func (c Azure) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Comparative) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Deepgram) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Simple) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}

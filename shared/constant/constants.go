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

type Android string           // Always "android"
type BookAppointment string   // Always "book_appointment"
type CheckAvailability string // Always "check_availability"
type Comparative string       // Always "comparative"
type Deepgram string          // Always "Deepgram"
type Function string          // Always "function"
type Handoff string           // Always "handoff"
type Ios string               // Always "ios"
type MediaName string         // Always "media_name"
type MediaURL string          // Always "media_url"
type PredefinedMedia string   // Always "predefined_media"
type Refer string             // Always "refer"
type Retrieval string         // Always "retrieval"
type SendDtmf string          // Always "send_dtmf"
type Simple string            // Always "simple"

func (c Android) Default() Android                     { return "android" }
func (c BookAppointment) Default() BookAppointment     { return "book_appointment" }
func (c CheckAvailability) Default() CheckAvailability { return "check_availability" }
func (c Comparative) Default() Comparative             { return "comparative" }
func (c Deepgram) Default() Deepgram                   { return "Deepgram" }
func (c Function) Default() Function                   { return "function" }
func (c Handoff) Default() Handoff                     { return "handoff" }
func (c Ios) Default() Ios                             { return "ios" }
func (c MediaName) Default() MediaName                 { return "media_name" }
func (c MediaURL) Default() MediaURL                   { return "media_url" }
func (c PredefinedMedia) Default() PredefinedMedia     { return "predefined_media" }
func (c Refer) Default() Refer                         { return "refer" }
func (c Retrieval) Default() Retrieval                 { return "retrieval" }
func (c SendDtmf) Default() SendDtmf                   { return "send_dtmf" }
func (c Simple) Default() Simple                       { return "simple" }

func (c Android) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c BookAppointment) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c CheckAvailability) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Comparative) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Deepgram) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Handoff) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Ios) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c MediaName) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c MediaURL) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c PredefinedMedia) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Refer) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Retrieval) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c SendDtmf) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Simple) MarshalJSON() ([]byte, error)            { return marshalString(c) }

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

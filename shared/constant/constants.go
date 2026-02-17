// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
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
type Assistant string         // Always "assistant"
type BookAppointment string   // Always "book_appointment"
type CheckAvailability string // Always "check_availability"
type Comparative string       // Always "comparative"
type Deepgram string          // Always "Deepgram"
type DeepgramNova2 string     // Always "deepgram/nova-2"
type DeepgramNova3 string     // Always "deepgram/nova-3"
type Developer string         // Always "developer"
type Function string          // Always "function"
type Handoff string           // Always "handoff"
type Ios string               // Always "ios"
type MediaName string         // Always "media_name"
type MediaURL string          // Always "media_url"
type Minimax string           // Always "minimax"
type PredefinedMedia string   // Always "predefined_media"
type Refer string             // Always "refer"
type Retrieval string         // Always "retrieval"
type SendDtmf string          // Always "send_dtmf"
type SendMessage string       // Always "send_message"
type Simple string            // Always "simple"
type SkipTurn string          // Always "skip_turn"
type System string            // Always "system"
type Tool string              // Always "tool"
type Transfer string          // Always "transfer"
type User string              // Always "user"

func (c Android) Default() Android                     { return "android" }
func (c Assistant) Default() Assistant                 { return "assistant" }
func (c BookAppointment) Default() BookAppointment     { return "book_appointment" }
func (c CheckAvailability) Default() CheckAvailability { return "check_availability" }
func (c Comparative) Default() Comparative             { return "comparative" }
func (c Deepgram) Default() Deepgram                   { return "Deepgram" }
func (c DeepgramNova2) Default() DeepgramNova2         { return "deepgram/nova-2" }
func (c DeepgramNova3) Default() DeepgramNova3         { return "deepgram/nova-3" }
func (c Developer) Default() Developer                 { return "developer" }
func (c Function) Default() Function                   { return "function" }
func (c Handoff) Default() Handoff                     { return "handoff" }
func (c Ios) Default() Ios                             { return "ios" }
func (c MediaName) Default() MediaName                 { return "media_name" }
func (c MediaURL) Default() MediaURL                   { return "media_url" }
func (c Minimax) Default() Minimax                     { return "minimax" }
func (c PredefinedMedia) Default() PredefinedMedia     { return "predefined_media" }
func (c Refer) Default() Refer                         { return "refer" }
func (c Retrieval) Default() Retrieval                 { return "retrieval" }
func (c SendDtmf) Default() SendDtmf                   { return "send_dtmf" }
func (c SendMessage) Default() SendMessage             { return "send_message" }
func (c Simple) Default() Simple                       { return "simple" }
func (c SkipTurn) Default() SkipTurn                   { return "skip_turn" }
func (c System) Default() System                       { return "system" }
func (c Tool) Default() Tool                           { return "tool" }
func (c Transfer) Default() Transfer                   { return "transfer" }
func (c User) Default() User                           { return "user" }

func (c Android) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c BookAppointment) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c CheckAvailability) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Comparative) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Deepgram) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c DeepgramNova2) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c DeepgramNova3) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Developer) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Handoff) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Ios) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c MediaName) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c MediaURL) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Minimax) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c PredefinedMedia) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Refer) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Retrieval) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c SendDtmf) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c SendMessage) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Simple) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c SkipTurn) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Transfer) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)              { return marshalString(c) }

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

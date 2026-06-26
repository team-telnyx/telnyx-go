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

type Android string         // Always "android"
type Arithmetic string      // Always "arithmetic"
type Assistant string       // Always "assistant"
type BoolLiteral string     // Always "bool_literal"
type BoolOp string          // Always "bool_op"
type Comparative string     // Always "comparative"
type Comparison string      // Always "comparison"
type Default string         // Always "default"
type Error string           // Always "error"
type Expression string      // Always "expression"
type Function string        // Always "function"
type Handoff string         // Always "handoff"
type Invite string          // Always "invite"
type Ios string             // Always "ios"
type Llm string             // Always "llm"
type MediaName string       // Always "media_name"
type MediaURL string        // Always "media_url"
type Node string            // Always "node"
type NumberLiteral string   // Always "number_literal"
type PredefinedMedia string // Always "predefined_media"
type Refer string           // Always "refer"
type Retrieval string       // Always "retrieval"
type SendDtmf string        // Always "send_dtmf"
type SendMessage string     // Always "send_message"
type Simple string          // Always "simple"
type SkipTurn string        // Always "skip_turn"
type StringLiteral string   // Always "string_literal"
type Transcript string      // Always "transcript"
type Transfer string        // Always "transfer"
type Variable string        // Always "variable"

func (c Android) Default() Android                 { return "android" }
func (c Arithmetic) Default() Arithmetic           { return "arithmetic" }
func (c Assistant) Default() Assistant             { return "assistant" }
func (c BoolLiteral) Default() BoolLiteral         { return "bool_literal" }
func (c BoolOp) Default() BoolOp                   { return "bool_op" }
func (c Comparative) Default() Comparative         { return "comparative" }
func (c Comparison) Default() Comparison           { return "comparison" }
func (c Default) Default() Default                 { return "default" }
func (c Error) Default() Error                     { return "error" }
func (c Expression) Default() Expression           { return "expression" }
func (c Function) Default() Function               { return "function" }
func (c Handoff) Default() Handoff                 { return "handoff" }
func (c Invite) Default() Invite                   { return "invite" }
func (c Ios) Default() Ios                         { return "ios" }
func (c Llm) Default() Llm                         { return "llm" }
func (c MediaName) Default() MediaName             { return "media_name" }
func (c MediaURL) Default() MediaURL               { return "media_url" }
func (c Node) Default() Node                       { return "node" }
func (c NumberLiteral) Default() NumberLiteral     { return "number_literal" }
func (c PredefinedMedia) Default() PredefinedMedia { return "predefined_media" }
func (c Refer) Default() Refer                     { return "refer" }
func (c Retrieval) Default() Retrieval             { return "retrieval" }
func (c SendDtmf) Default() SendDtmf               { return "send_dtmf" }
func (c SendMessage) Default() SendMessage         { return "send_message" }
func (c Simple) Default() Simple                   { return "simple" }
func (c SkipTurn) Default() SkipTurn               { return "skip_turn" }
func (c StringLiteral) Default() StringLiteral     { return "string_literal" }
func (c Transcript) Default() Transcript           { return "transcript" }
func (c Transfer) Default() Transfer               { return "transfer" }
func (c Variable) Default() Variable               { return "variable" }

func (c Android) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Arithmetic) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c BoolLiteral) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c BoolOp) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Comparative) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Comparison) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Default) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Expression) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Handoff) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Invite) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Ios) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Llm) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c MediaName) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c MediaURL) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Node) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c NumberLiteral) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c PredefinedMedia) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Refer) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Retrieval) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c SendDtmf) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c SendMessage) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Simple) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c SkipTurn) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c StringLiteral) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Transcript) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Transfer) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Variable) MarshalJSON() ([]byte, error)        { return marshalString(c) }

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

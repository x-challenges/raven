package json

import (
	jsonstd "encoding/json"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// aliases
var (
	NewEncoder          = json.NewEncoder
	NewDecoder          = json.NewDecoder
	Marshal             = json.Marshal
	MarshalToString     = json.MarshalToString
	Unmarshal           = json.Unmarshal
	UnmarshalFromString = json.UnmarshalFromString
)

type (
	RawMessage = jsonstd.RawMessage
)

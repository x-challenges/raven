package json

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// aliases
var (
	NewEncoder          = json.NewEncoder
	Marshal             = json.Marshal
	MarshalToString     = json.MarshalToString
	Unmarshal           = json.Unmarshal
	UnmarshalFromString = json.UnmarshalFromString
)

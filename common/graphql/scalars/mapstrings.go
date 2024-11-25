package scalars

import (
	"io"

	"github.com/x-challenges/raven/common/json"
)

type MapStrings map[string][]string

func (o *MapStrings) UnmarshalGQL(v interface{}) error {
	return json.Unmarshal(v.([]byte), o)
}

func (o MapStrings) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(o) //nolint
}

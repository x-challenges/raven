package scalars

import (
	"io"

	"github.com/x-challenges/raven/common/json"
)

type Object map[string]interface{}

func (o *Object) UnmarshalGQL(v interface{}) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, o)
	case map[string]interface{}:
		*o = vv
	}
	return nil
}

func (o Object) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(o) //nolint
}

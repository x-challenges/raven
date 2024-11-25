package scalars

import (
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"

	"github.com/x-challenges/raven/common/json"
)

func MarshalMapString(m map[string]string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		json.NewEncoder(w).Encode(m) //nolint
	})
}

func UnmarshalMapString(m interface{}) (map[string]string, error) {
	tmpMap, ok := m.(map[string]interface{})
	if !ok {
		return nil, errors.New("value should be a map[string]string")
	}

	result := make(map[string]string, len(tmpMap))

	for k, v := range tmpMap {
		result[k] = fmt.Sprintf("%v", v)
	}
	return result, nil
}

package scalars

import (
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

func MarshalDecimal(t decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.String())) // nolint
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	if tmpStr, ok := v.(string); ok {
		return decimal.NewFromString(tmpStr)
	}
	return decimal.Zero, errors.New("value shoud be a decimal")
}

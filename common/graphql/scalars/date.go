package scalars

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

const (
	dateFormatLayout = "2006-01-02"
)

func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(dateFormatLayout))) //nolint
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(dateFormatLayout, tmpStr)
	}
	return time.Time{}, errors.New("value should be a date string `2006-01-02`")
}

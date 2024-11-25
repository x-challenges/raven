package serializers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type JSONSerializer struct {
	json        jsoniter.API
	prettyPrint bool
}

func NewJSONSerialzer(prettyPrint bool) *JSONSerializer {
	return &JSONSerializer{
		prettyPrint: prettyPrint,
		json:        jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

func (j *JSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := j.json.NewEncoder(c.Response())

	if j.prettyPrint && indent != "" {
		enc.SetIndent("", indent)
	}

	return enc.Encode(i)
}

func (j *JSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := j.json.NewDecoder(c.Request().Body).Decode(i)

	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("unmarshal type error: expected=%v, got=%v, field=%v, offset=%v",
				ute.Type, ute.Value, ute.Field, ute.Offset),
		).SetInternal(err)
	}

	var se *json.SyntaxError
	if errors.As(err, &se) {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("syntax error: offset=%v, error=%v",
				se.Offset, se.Error(),
			),
		).SetInternal(err)
	}
	return err
}

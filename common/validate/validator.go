package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/x-challenges/raven/common/errors"
)

var (
	Struct = decorate(validate.Struct)
	Var    = validate.Var
)

var (
	validate = validator.New(
		validator.WithRequiredStructEnabled(),
	)
	uni   = ut.New(en.New(), en.New())
	trans ut.Translator
)

func init() {
	// register function for extract field name for validator from json struct tag
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name != "" && name != "-" {
			return name
		}

		return ""
	})

	// register EN ranslators
	trans, _ = uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
}

// transform validation error to our custom error
func decorate(fn func(s interface{}) error) func(s interface{}) error {
	return func(s interface{}) error {
		var (
			fields = errors.Fields{}
			err    error
		)

		// skip if error empty
		if err = fn(s); err == nil {
			return nil
		}

		for _, e := range err.(validator.ValidationErrors) {
			fields = append(fields, errors.String(e.Field(), e.Translate(trans)))
		}

		return errors.WithFields(errors.ErrInvalidRequest, fields...)
	}
}

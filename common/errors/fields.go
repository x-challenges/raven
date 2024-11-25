package errors

import "time"

type FieldVisibility int

const (
	Public FieldVisibility = iota
	Private
)

func GetFields(err error, withPrivate bool) Fields {
	data := make([]Field, 0)

	for err != nil {
		if e, ok := err.(*Error); ok {
			for _, f := range e.Fields {
				if f.Visibility == Private && !withPrivate {
					continue
				}
				data = append(data, f)
			}
		}

		err = Unwrap(err)
	}
	return data
}

// IsPrivate field visibility functional parameter.
func IsPrivate() func(*Field) {
	return func(f *Field) {
		f.Visibility = Private
	}
}

// IsPublic field visibility functional parameter.
func IsPublic() func(*Field) {
	return func(f *Field) {
		f.Visibility = Public
	}
}

func applyFieldOptions(f *Field, opts ...func(*Field)) {
	for _, o := range opts {
		o(f)
	}
}

// Field type.
type Field struct {
	Key        string          `json:"key"`
	Value      interface{}     `json:"value"`
	Visibility FieldVisibility `json:"-"`
}

func NewField(key string, value interface{}) *Field {
	return &Field{
		Key:        key,
		Value:      value,
		Visibility: Private,
	}
}

// String implements Stringer interface.
func (f *Field) String() string {
	return f.Key
}

// Fields type.
type Fields []Field

// Values returns all fields value as a map.
func (f Fields) Values() map[string]interface{} {
	var (
		result = make(map[string]interface{}, len(f))
	)

	for _, field := range f {
		result[field.Key] = field.Value
	}
	return result
}

// String create new field with String type.
func String(key string, value string, opts ...func(*Field)) Field {
	field := NewField(key, value)

	applyFieldOptions(field, opts...)
	return *field
}

// Int create new field with Integer type.
func Int(key string, value int, opts ...func(*Field)) Field {
	field := NewField(key, value)

	applyFieldOptions(field, opts...)
	return *field
}

// Boolean create new field with Boolean type.
func Boolean(key string, value bool, opts ...func(*Field)) Field {
	field := NewField(key, value)

	applyFieldOptions(field, opts...)
	return *field
}

// Object create new field with Any type.
func Object(key string, value any, opts ...func(*Field)) Field {
	field := NewField(key, value)

	applyFieldOptions(field, opts...)
	return *field
}

// Duration
func Duration(key string, value time.Duration, opts ...func(*Field)) Field {
	field := NewField(key, value)

	applyFieldOptions(field, opts...)
	return *field
}

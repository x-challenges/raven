package config

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func bindenvs(viper *viper.Viper, iface interface{}, parts ...string) error {
	ifv := reflect.ValueOf(iface)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}

	for i := 0; i < ifv.NumField(); i++ {
		v := ifv.Field(i)
		t := ifv.Type().Field(i)

		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		if tv == ",squash" {
			if err := bindenvs(viper, v.Interface(), parts...); err != nil {
				return err
			}
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			if err := bindenvs(viper, v.Interface(), append(parts, tv)...); err != nil {
				return err
			}
		default:
			if err := viper.BindEnv(strings.Join(append(parts, tv), ".")); err != nil {
				return err
			}
		}
	}
	return nil
}

package envvar

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var DefaultFormatter = NewFormatter("", "", 0)

// GetEnv get enviroment varible's value into struct
func GetEnv(ptr interface{}, f Formatter) error {

	m := make(map[string]interface{})

	// rv for reflect Value
	rvPtr := reflect.ValueOf(ptr)
	if rvPtr.Kind() != reflect.Ptr {
		msg := fmt.Sprintf("Want a Struct Prt, Got a %s", rvPtr.Kind())
		return errors.New(msg)
	}

	// sValue := reflect.Indirect(reflect.ValueOf(v))
	rv := reflect.Indirect(rvPtr)
	if rv.Kind() != reflect.Struct {
		msg := fmt.Sprintf("Want a Struct, Got a %s", rv.Kind())
		return errors.New(msg)
	}

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		// spew.Dump(sFiled)

		tag, ok := sFiled.Tag.Lookup("env")
		if !ok {
			continue
		}

		name := strings.Split(tag, ",")[0]
		if name == "-" {
			continue
		}

		envName := format(name, f)

		envValue, ok := os.LookupEnv(envName)
		if !ok {
			continue
		}

		switch sFiled.Type.Kind() {
		case reflect.String:
			m[name] = envValue
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[name] = MustParseInteger(envValue)
		case reflect.Bool:
			m[name] = MustParseBool(envValue)
		}
	}
	// spew.Dump(m)

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, ptr)
	if err != nil {
		return err
	}

	return nil
}

// SetEnv set struct value to envirotment
func SetEnv(v interface{}, f Formatter) {

	rv := reflect.Indirect(reflect.ValueOf(v))

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		if tag, ok := sFiled.Tag.Lookup("env"); ok {
			th := strings.Split(tag, ",")[0]
			if th == "-" {
				continue
			}

			sValue := rv.Field(i)
			// envName := fmt.Sprintf("%s_%s", prefix, strings.ToUpper(th))
			envName := format(th, f)

			switch sValue.Kind() {
			case reflect.String:
				os.Setenv(envName, sValue.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
				// os.Setenv(envName, string(sValue.Int()))
				os.Setenv(envName, strconv.FormatInt(sValue.Int(), 10))
			case reflect.Bool:
				os.Setenv(envName, strconv.FormatBool(sValue.Bool()))
			}
		}
	}
}

func MustParseInteger(s string) (n int64) {
	if s == "" {
		return 0
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return
}

func MustParseBool(s string) (ok bool) {
	if s == "" {
		return false
	}
	ok, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return
}

func format(s string, f Formatter) string {
	r := fmt.Sprintf("%s_%s_%s",
		f.prefix,
		f.handler(s),
		f.suffix)

	return strings.Trim(r, "_")
}

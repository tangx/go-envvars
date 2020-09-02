package envcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Drain(v interface{}, prefix string) error {
	m := map[string]interface{}{}

	return drain(v, prefix, m)
}

// drain convert struct to os env
func drain(v interface{}, prefix string, m map[string]interface{}) (err error) {

	rv := reflect.Indirect(reflect.ValueOf(v))

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)

		envTag, ok := sFiled.Tag.Lookup("env")

		if !ok {
			continue
		}

		envTagName := strings.Split(envTag, ",")[0]
		if envTagName == "-" {
			continue
		}

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, envTagName))

		sValue := rv.Field(i)

		switch sValue.Kind() {
		case reflect.String:
			m[key] = defaultString(sFiled, sValue.String())

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[key] = defaultInt(sFiled, sValue.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
			m[key] = defaultUint(sFiled, sValue.Uint())

		case reflect.Bool:
			m[key] = defaultBool(sFiled, sValue.Bool())

		case reflect.Struct:
			_ = drain(sValue.Interface(), key, m)
		}

	}

	b, err := YamlMarshal(m)
	if err != nil {
		return err
	}

	err = WriteToFile("config.yml", b)
	return
}

func defaultString(sFiled reflect.StructField, value string) string {
	if value != "" {
		return value
	}
	tag, ok := sFiled.Tag.Lookup("default")
	if !ok {
		return ""
	}
	return strings.Split(tag, ",")[0]
}

func defaultBool(sFiled reflect.StructField, value bool) bool {

	if value {
		return value
	}
	tag, ok := sFiled.Tag.Lookup("default")
	if !ok {
		return false
	}

	tmp := strings.Split(tag, ",")[0]
	_, err := strconv.ParseBool(tmp)
	return err == nil

}

func defaultInt(sFiled reflect.StructField, value int64) int64 {
	if value != 0 {
		return value
	}

	tag, ok := sFiled.Tag.Lookup("default")
	if !ok {
		return 0
	}
	tmp := strings.Split(tag, ",")[0]
	i, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func defaultUint(sFiled reflect.StructField, value uint64) uint64 {
	if value != 0 {
		return value
	}

	tag, ok := sFiled.Tag.Lookup("default")
	if !ok {
		return 0
	}
	tmp := strings.Split(tag, ",")[0]
	i, err := strconv.ParseUint(tmp, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

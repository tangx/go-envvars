package envcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var m = make(map[string]interface{})

// Unmarshal convert struct to os env
func Unmarshal(prefix string, v interface{}) (err error) {
	// spew.Dump(v)

	rv := reflect.Indirect(reflect.ValueOf(v))

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		// spew.Dump(sFiled)

		envTag, ok := sFiled.Tag.Lookup("env")
		// fmt.Println("envTag=", envTag)

		if !ok {
			continue
		}

		envTagName := strings.Split(envTag, ",")[0]
		if envTagName == "-" {
			continue
		}

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, envTagName))
		// m[key] = rv.Field(i).String()
		// fmt.Println("key=", key)

		sValue := rv.Field(i)

		switch sValue.Kind() {
		case reflect.String:
			val := sValue.String()
			m[key] = defaultString(sFiled, val)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// m[key] = strconv.FormatInt(sValue.Int(), 10)
			m[key] = int64(sValue.Int())

		case reflect.Bool:
			// m[key] = strconv.FormatBool(sValue.Bool())
			m[key] = sValue.Bool()

		case reflect.Struct:
			_ = Unmarshal(key, sValue.Interface())

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

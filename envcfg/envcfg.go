package envcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var m = make(map[string]string)

// Unmarshal convert struct to os env
func Unmarshal(prefix string, v interface{}) (err error) {
	// spew.Dump(v)

	rv := reflect.Indirect(reflect.ValueOf(v))

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		tag, ok := sFiled.Tag.Lookup("env")

		if !ok {
			continue
		}

		tagName := strings.Split(tag, ",")[0]
		if tagName == "-" {
			continue
		}

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, tagName))
		// m[key] = rv.Field(i).String()
		// fmt.Println("key=", key)

		sValue := rv.Field(i)

		switch sValue.Kind() {

		case reflect.String:
			m[key] = sValue.String()

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[key] = strconv.FormatInt(sValue.Int(), 10)

		case reflect.Bool:
			m[key] = strconv.FormatBool(sValue.Bool())

		case reflect.Struct:
			_ = Unmarshal(key, sValue.Interface())

		}

	}

	spew.Dump(m)

	return
}

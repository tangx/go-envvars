package envcfg

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// Pump value from env to struct
func Pump(v interface{}, prefix string) error {
	m := map[string]interface{}{}

	pump(v, prefix, m)

	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, v)
}

// pump value from env to struct
// v should be a ptr
func pump(v interface{}, prefix string, m map[string]interface{}) {

	rvPtr := reflect.ValueOf(v)
	// if rvPtr.Kind() == reflect.Ptr {
	// 	return
	// }

	rv := reflect.Indirect(rvPtr)
	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		sValue := rv.Field(i)

		tag, ok := sFiled.Tag.Lookup("env")
		var tagName string
		if ok {
			tagName = strings.Split(tag, ",")[0]
		} else {
			tagName = strings.ToLower(sFiled.Name)
		}

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, tagName))

		switch sValue.Kind() {
		case reflect.String:
			m[tagName] = os.Getenv(key)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

			typ := sValue.Type()
			// time.Duration
			if sValue.Kind() == reflect.Int64 && typ.PkgPath() == "time" && typ.Name() == "Duration" {
				m[tagName] = mustTimeDuration(os.Getenv(key))
			} else {
				m[tagName] = mustInt(os.Getenv(key))
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

			m[tagName] = mustUint(os.Getenv(key))

		case reflect.Bool:
			m[tagName] = mustBool(os.Getenv(key))

		case reflect.Struct:
			m2 := map[string]interface{}{}
			pump(sValue.Interface(), key, m2)

			m[tagName] = m2
		}
	}

}

func mustBool(str string) bool {
	boolean, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}

	return boolean
}

func mustInt(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func mustUint(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func mustTimeDuration(str string) time.Duration {
	dur, err := time.ParseDuration(str)
	if err != nil {
		panic(err)
	}
	return dur
}

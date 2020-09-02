package envcfg

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

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
		tag, ok := sFiled.Tag.Lookup("env")
		if !ok {
			continue
		}
		tagName := strings.Split(tag, ",")[0]

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, tagName))

		sValue := rv.Field(i)
		switch sValue.Kind() {
		case reflect.String:
			m[tagName] = os.Getenv(key)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[tagName] = shouldInt(os.Getenv(key))

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			m[tagName] = shouldUint(os.Getenv(key))

		case reflect.Bool:
			m[tagName] = shouldBool(os.Getenv(key))

		case reflect.Struct:
			// fmt.Printf("m[%s]", tagName)
			m2 := map[string]interface{}{}
			pump(sValue.Interface(), key, m2)

			m[tagName] = m2
		}
	}

}

// Pump
func PumpFileToEnv(file string) (err error) {
	m := map[string]interface{}{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(b, &m); err != nil {
		return
	}

	for k, v := range m {
		os.Setenv(k, shouldString(v))
	}

	return nil
}

func shouldString(v interface{}) string {
	rv := reflect.ValueOf(v)
	typ := rv.Type()

	// spew.Dump(rv)
	switch typ.Kind() {
	case reflect.String:
		return rv.String()
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	}

	return ""
}

func shouldBool(str string) bool {
	boolean, _ := strconv.ParseBool(str)

	return boolean
}

func shouldInt(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)

	return i
}

func shouldUint(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

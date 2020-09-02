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
// v should be a ptr
func Pump(ptr interface{}, prefix string, m map[string]interface{}) {

	// spew.Dump(ptr)

	rvPtr := reflect.ValueOf(ptr)
	if rvPtr.Kind() == reflect.Ptr {
		return
	}

	rv := reflect.Indirect(rvPtr)
	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		// spew.Dump(sFiled)
		tag, ok := sFiled.Tag.Lookup("env")
		if !ok {
			continue
		}
		tagName := strings.Split(tag, ",")[0]

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, tagName))
		// fmt.Printf("key( %s ) = value ( %s )\n", key, os.Getenv(key))

		sValue := rv.Field(i)
		switch sValue.Kind() {
		case reflect.String:
			m[tagName] = os.Getenv(key)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[tagName] = ShouldInt(os.Getenv(key))

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			m[tagName] = ShouldUint(os.Getenv(key))

		case reflect.Bool:
			m[tagName] = ShouldBool(os.Getenv(key))

		case reflect.Struct:
			// fmt.Printf("m[%s]", tagName)
			m2 := map[string]interface{}{}
			Pump(sValue.Interface(), key, m2)

			m[tagName] = m2
		}
	}
	// spew.Dump(m)

	// fmt.Printf("%s\n", m)

	// b, err := json.Marshal(m)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%s\n", b)
	// err = json.Unmarshal(b, ptr)
	// if err != nil {
	// 	panic(err)
	// }
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
		os.Setenv(k, convert(v))
	}

	return nil
}

func convert(v interface{}) string {
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

func ShouldBool(str string) bool {
	boolean, _ := strconv.ParseBool(str)

	return boolean
}

func ShouldInt(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)

	return i
}

func ShouldUint(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

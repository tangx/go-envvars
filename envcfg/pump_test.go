package envcfg

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

func TestPump(t *testing.T) {
	f := configFile
	_ = LoadConfigFileToEnv(f)

	var p = Person{}
	_ = Pump(&p, "USER")

	fmt.Println("=========")
	spew.Dump(p)
}

func LoadConfigFileToEnv(file string) (err error) {
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

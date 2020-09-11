package envcfg

import (
	"fmt"
	"reflect"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"
)

func Drain(v interface{}, prefix string) (err error) {

	m := &yamlv3.Node{
		Kind: 4,
	}

	_ = drain(v, prefix, m)
	// spew.Dump(m)

	// yamlv3.Marshal(m)
	b, err := YamlMarshal(m)
	if err != nil {
		panic(err)
	}
	_ = WriteToFile(configFile, b)
	return
}

// drain convert struct to os env
func drain(v interface{}, prefix string, m *yamlv3.Node) (err error) {

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

		commentTag := sFiled.Tag.Get("comment")
		if commentTag == "-" {
			commentTag = ""
		}

		valueTag := sFiled.Tag.Get("default")
		if valueTag == "-" {
			valueTag = ""
		}

		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, envTagName))
		sValue := rv.Field(i)

		switch sValue.Kind() {
		case reflect.String:
			contents := Yamlv3Marshal(key, valueTag, commentTag, "!!str")

			m.Content = append(m.Content, contents...)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if valueTag == "" {
				valueTag = "0"
			}
			m.Content = append(m.Content, Yamlv3Marshal(key, valueTag, commentTag, "!!int")...)

		case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
			if valueTag == "" {
				valueTag = "0"
			}
			m.Content = append(m.Content, Yamlv3Marshal(key, valueTag, commentTag, "!!int")...)

		case reflect.Bool:
			if valueTag == "" {
				valueTag = "false"
			}
			m.Content = append(m.Content, Yamlv3Marshal(key, valueTag, commentTag, "!!bool")...)

		case reflect.Struct:
			_ = drain(sValue.Interface(), key, m)
		}
	}

	return
}

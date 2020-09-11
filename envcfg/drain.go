package envcfg

import (
	"fmt"
	"reflect"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"
)

func Drain(v interface{}, prefix string, config string) (err error) {

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
	_ = WriteToFile(config, b)
	return
}

// drain convert struct to os env
func drain(v interface{}, prefix string, m *yamlv3.Node) (err error) {

	rv := reflect.Indirect(reflect.ValueOf(v))

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		sFiled := typ.Field(i)
		sValue := rv.Field(i)

		envTag, ok := sFiled.Tag.Lookup("env")
		var envTagName string
		if !ok {
			if sFiled.Type.Kind() == reflect.Struct {
				envTagName = sFiled.Name
				key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, envTagName))
				_ = drain(sValue.Interface(), key, m)
			}
			continue
		}
		envTagName = strings.Split(envTag, ",")[0]
		if envTagName == "-" || envTagName == "" {
			continue
		}
		key := strings.ToUpper(fmt.Sprintf("%s__%s", prefix, envTagName))

		commentTag := sFiled.Tag.Get("comment")
		if commentTag == "-" {
			commentTag = ""
		}

		valueTag := sFiled.Tag.Get("default")
		if valueTag == "-" {
			valueTag = ""
		}

		var yamlTag string
		switch sValue.Kind() {
		case reflect.String:
			yamlTag = "!!str"

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if valueTag == "" {
				valueTag = "0"
			}
			yamlTag = "!!int"

			// time.Duration
			if IsTimeDuration(sValue) {
				yamlTag = "!!str"
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
			if valueTag == "" {
				valueTag = "0"
			}

			yamlTag = "!!int"

		case reflect.Bool:
			if valueTag == "" {
				valueTag = "false"
			}
			yamlTag = "!!bool"

			// case reflect.Struct:
			// 	_ = drain(sValue.Interface(), key, m)
		}

		if yamlTag == "" {
			continue
		}
		m.Content = append(m.Content, combineContent(key, valueTag, commentTag, yamlTag)...)

	}

	return
}

func combineContent(key string, value interface{}, comment string, tag string) []*yamlv3.Node {

	k := &yamlv3.Node{
		Kind:        8,
		Value:       key,
		HeadComment: comment,
	}
	v := &yamlv3.Node{
		Kind:  8,
		Value: value.(string),
		Tag:   tag,
	}

	return []*yamlv3.Node{k, v}

}

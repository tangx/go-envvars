package envcfg

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"

	yamlv3 "gopkg.in/yaml.v3"
)

func YamlMarshal(v interface{}) ([]byte, error) {
	return yamlv3.Marshal(v)
}

func WriteTo(w io.Writer, b []byte) (err error) {

	buf := bytes.NewBuffer(b)
	_, err = buf.WriteTo(w)
	return
}

func WriteToFile(file string, b []byte) (err error) {
	_ = os.MkdirAll(filepath.Dir(file), os.ModePerm)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return WriteTo(f, b)
}

func IsTimeDuration(v reflect.Value) bool {
	typ := v.Type()
	return v.Kind() == reflect.Int64 && typ.PkgPath() == "time" && typ.Name() == "Duration"
}

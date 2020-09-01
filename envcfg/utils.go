package envcfg

import (
	"bytes"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func YamlMarshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func WriteTo(w io.Writer, b []byte) (err error) {

	buf := bytes.NewBuffer(b)
	_, err = buf.WriteTo(w)
	return
}

func WriteToFile(path string, b []byte) (err error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return WriteTo(f, b)
}

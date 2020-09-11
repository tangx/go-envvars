package envcfg

import (
	"bytes"
	"io"
	"log"
	"os"

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

func WriteToFile(path string, b []byte) (err error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return WriteTo(f, b)
}

func Yamlv3Marshal(key string, value interface{}, comment string, tag string) []*yamlv3.Node {

	k := &yamlv3.Node{
		Kind:        8,
		Value:       key,
		HeadComment: comment,
	}
	v := &yamlv3.Node{
		Kind:  8,
		Value: value.(string),
		// Tag: "!!str",
		Tag: tag,
	}

	return []*yamlv3.Node{k, v}

}

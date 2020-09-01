package envcfg

import (
	"testing"
)

func TestWriteTo(t *testing.T) {
	b := []byte(`hello world`)

	// WriteTo(os.Stdout, b)
	_ = WriteToFile("config.yml", b)
}

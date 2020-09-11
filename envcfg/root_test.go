package envcfg

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRoot(t *testing.T) {
	e := New("__snapshot__/config.yml", "USER")
	p := Person{}
	var config = &struct {
		Person Person
	}{
		Person: p,
	}

	_ = e.Drain(config)

	_ = LoadConfigFileToEnv(e.Config)
	_ = e.Pump(config)

	spew.Dump(config)
}

package envcfg

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPump(t *testing.T) {
	f := configFile
	_ = PumpFileToEnv(f)

	var p = Person{}
	_ = Pump(&p, "USER")

	fmt.Println("=========")
	spew.Dump(p)
}

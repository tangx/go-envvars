package envcfg

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPump(t *testing.T) {
	f := `config.yml`
	_ = PumpFileToEnv(f)

	println(os.Getenv("USER__ADDRESS__NUMBER"))
	println(os.Getenv("USER__GENDER"))
	println(os.Getenv("USER__NAME"))

	// p := Person{}
	// p := new(Person)
	var p = Person{
		// Name: "wusangui",
	}
	m := map[string]interface{}{}
	Pump(p, "USER", m)
	// fmt.Println("over")

	spew.Dump(m)
	b, _ := json.Marshal(m)
	fmt.Printf("%s", b)

	_ = json.Unmarshal(b, &p)
	spew.Dump(p)
}

func println(s string) {
	fmt.Println(s)
}

package envcfg

import (
	"testing"
)

func TestParse(t *testing.T) {

	// p := Person{"zhangsan", 10, true, Address{"sichuan", "chengdu", 10}}

	p := new(Person)
	_ = Unmarshal("USER", p)

}

type Person struct {
	Name    string `env:"name"`
	Age     int32  `env:"age" envDefault:"18"`
	Gender  bool   `env:"gender" envDefault:"false"`
	Address `env:"address"`
}

type Address struct {
	City   string `env:"city"`
	Street string `env:"street"`
	Number int64  `env:"number"`
}

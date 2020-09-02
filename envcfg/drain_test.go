package envcfg

import (
	"testing"
)

func TestDrain(t *testing.T) {

	// p := Person{"zhangsan", 10, true, Address{"sichuan", "chengdu", 10}}

	p := Person{
		Name: "zhangsan",
	}
	_ = Drain("USER", p)

}

type Person struct {
	Name    string `env:"name,omitempty" default:"zhuageliang"`
	Age     uint   `env:"age,omitempty" default:"18"`
	Gender  bool   `env:"gender,omitempty" default:"true"`
	Address `env:"address"`
}

type Address struct {
	City   string `env:"city" default:"sichuan"`
	Street string `env:"street" `
	Number int64  `env:"number" default:"100"`
}
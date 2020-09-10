package envcfg

import (
	"testing"
)

func TestDrain(t *testing.T) {

	// p := Person{"zhangsan", 10, true, Address{"sichuan", "chengdu", 10}}

	p := Person{
		Name: "zhangsan",
	}
	_ = Drain(p, "USER")

}

type Person struct {
	Name    string `env:"name,omitempty" default:"zhuageliang" yaml:"name,omitempty" summary:"user name"`
	Age     uint   `env:"age,omitempty" default:"18" yaml:"age,omitempty" summary:"user age"`
	Gender  bool   `env:"gender,omitempty" default:"true" yaml:"gender,omitempty" summary:"user gender"`
	Address `env:"address,omitempty" yaml:"address,omitempty" summary:"user address"`
}

type Address struct {
	City   string `env:"city,omitempty" default:"sichuan" yaml:"city,omitempty" summary:"-"`
	Street string `env:"street,omitempty" yaml:"street,omitempty" summary:"apartment address street"`
	Number int64  `env:"number,omitempty" default:"99999182" yaml:"number,omitempty" summary:"apartment address number"`
}

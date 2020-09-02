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
	Name    string `env:"name,omitempty" default:"zhuageliang" yaml:"name,omitempty" json:"name,omitempty"`
	Age     uint   `env:"age,omitempty" default:"18" yaml:"age,omitempty" json:"age,omitempty"`
	Gender  bool   `env:"gender,omitempty" default:"true" yaml:"gender,omitempty" json:"gender,omitempty"`
	Address `env:"address,omitempty" yaml:"address,omitempty" json:"address,omitempty"`
}

type Address struct {
	City   string `env:"city,omitempty" default:"sichuan" yaml:"city,omitempty" json:"city,omitempty"`
	Street string `env:"street,omitempty" yaml:"street,omitempty" json:"street,omitempty"`
	Number int64  `env:"number,omitempty" default:"100" yaml:"number,omitempty" json:"number,omitempty"`
}

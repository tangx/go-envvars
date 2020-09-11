package envcfg

import (
	"testing"
	"time"
)

const (
	configFile = `__snapshot__/default.yml`
)

func TestDrain(t *testing.T) {

	// p := Person{"zhangsan", 10, true, Address{"sichuan", "chengdu", 10}}

	p := Person{
		Name: "zhangsan",
	}
	_ = Drain(p, "USER", configFile)

}

type Person struct {
	Address
	Food

	Name    string        `env:"name,omitempty" default:"zhuageliang" yaml:"name,omitempty" comment:"user name"`
	Age     uint          `env:"age,omitempty" default:"18" yaml:"age,omitempty" comment:"user age"`
	Gender  bool          `env:"gender,omitempty" default:"true" yaml:"gender,omitempty" comment:"user gender"`
	Timeout time.Duration `env:"timeout" yaml:"timeout" comment:"timeout to work" default:"5m"`
}

type Address struct {
	City   string `env:"city,omitempty" default:"sichuan" yaml:"city,omitempty" comment:"-"`
	Street string `env:"street,omitempty" yaml:"street,omitempty" comment:"apartment address street"`
	Number int64  `env:"number,omitempty" default:"99999182" yaml:"number,omitempty" comment:"apartment address number"`
}

type Food struct {
	Name  string
	Price int `env:"-" default:"30" comment:"food price, ($)"`
}

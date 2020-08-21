package envvar

import (
	"os"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/log"
	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

type Person struct {
	Name         string `env:"-,omitempty"`
	Age          int    `env:"age,omitempty"`
	Gender       bool   `env:"gender,omitempty"`
	EmailAddress string `env:"email_address,omitempty" json:"email_address,omitempty"`
}

var p = Person{"zhangsan", 10, true, "zhangsan@goole.com"}

func Test_marshal(t *testing.T) {

	// fmtter := NewDefaultFormatter().SetPrefix("USER").ToUpper()
	fmtter := NewFormatter("USER", "", 0)
	err := GetEnv(&p, fmtter)
	if err != nil {
		log.Error(err)
	}
	// spew.Dump(p)

	convey.Convey("env mashal test", t, func() {
		convey.So(p.Name, ShouldEqual, "zhangsan")
		convey.So(p.Age, ShouldEqual, 20)
		convey.So(p.Gender, ShouldBeFalse)
		convey.So(p.EmailAddress, ShouldEqual, "wangwu@qq.com")
	})
}

func Test_Set(t *testing.T) {
	f := NewDefaultFormatter().SetPrefix("USER").ToUpper()
	SetEnv(p, f)

	convey.Convey("Set env variable", t, func() {
		convey.So(os.Getenv("USER_NAME"), ShouldEqual, "wangwu")
		convey.So(os.Getenv("USER_AGE"), ShouldEqual, "10")
		convey.So(os.Getenv("USER_GENDER"), ShouldEqual, "true")
		convey.So(os.Getenv("USER_EMAIL_ADDRESS"), ShouldEqual, "zhangsan@goole.com")
	})
}

func init() {
	os.Setenv("USER_NAME", "wangwu")
	os.Setenv("USER_AGE", "20")
	os.Setenv("USER_GENDER", "false")
	os.Setenv("USER_EMAIL_ADDRESS", "wangwu@qq.com")
}

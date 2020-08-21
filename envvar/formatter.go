package envvar

import "strings"

const (
	ToUpper = iota
	ToLower
	ToTitle
)

type Formatter struct {
	prefix  string
	suffix  string
	handler func(string) string
}

func NewFormatter(pre, post string, formatter int) Formatter {
	return Formatter{
		prefix:  pre,
		suffix:  post,
		handler: trans(formatter),
	}
}

func NewDefaultFormatter() Formatter {
	return Formatter{"", "", strings.ToUpper}
}

func (f Formatter) SetPrefix(s string) Formatter {
	f.prefix = s
	return f
}
func (f Formatter) SetSuffix(s string) Formatter {
	f.suffix = s
	return f
}
func (f Formatter) ToUpper() Formatter {
	f.handler = strings.ToUpper
	return f
}
func (f Formatter) ToLower() Formatter {
	f.handler = strings.ToLower
	return f
}
func (f Formatter) ToTitle() Formatter {
	f.handler = strings.ToTitle
	return f
}

func trans(i int) (handler func(string) string) {
	switch i {
	case 1:
		return strings.ToLower
	case 2:
		return strings.ToTitle
	}

	return strings.ToUpper
}

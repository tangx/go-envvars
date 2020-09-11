package envcfg

type EnvCfg struct {
	Config string
	Prefix string
}

func New(config string, prefix string) *EnvCfg {
	return &EnvCfg{
		config,
		prefix,
	}
}

func (e *EnvCfg) Drain(v interface{}) error {
	return Drain(v, e.Prefix, e.Config)
}

func (e *EnvCfg) Pump(v interface{}) error {
	return Pump(v, e.Prefix)
}

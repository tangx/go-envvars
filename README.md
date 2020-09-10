# envcfg library

## todo

+ env struct
  +[x] nested env: `sValue:=rv.Filed(i); sValue.Interface()`

+ env field
  - [x] env
  - [x] default
  - [ ] summary

+ env support type
  + time.Duration
  + [x] int,int8,int16,int32,int64
  + [x] string
  + [x] bool

+ env action
  + [x] drain: `convert` struct to os env
  + [x] setenv: `load` os env to struct
  + `env` to `config.yml`


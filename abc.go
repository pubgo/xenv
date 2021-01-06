package xenv

import (
	"os"
	"strconv"
)

type Env string

func (env Env) New(name string) Env { return Env(env.key() + "_" + name) }
func (env Env) Enable()             { env.Set(strconv.FormatBool(true)) }
func (env Env) key() string         { return string(env) }
func (env Env) Unset() {
	envs.Delete(env.key())
	_ = os.Unsetenv(env.key())
}
func (env Env) Set(dt string) {
	envs.Store(env.key(), dt)
	_ = os.Setenv(env.key(), dt)
}
func (env Env) Get() string {
	val, ok := envs.Load(env.key())
	if ok {
		return val.(string)
	}
	return ""
}
func (env Env) Bool() bool     { b, _ := strconv.ParseBool(env.Get()); return b }
func (env Env) Int() int64     { b, _ := strconv.ParseInt(env.Get(), 10, 64); return b }
func (env Env) Float() float64 { b, _ := strconv.ParseFloat(env.Get(), 64); return b }

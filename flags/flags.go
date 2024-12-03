package flags

import (
	"errors"
	"fmt"

	"github.com/jasonhancock/go-helpers"
	"github.com/spf13/pflag"
)

type addFunc func(fs *pflag.FlagSet, f *flag)
type checkFunc func(f *flag) error

var flagTypes = map[string]flagInfo{}

type flagInfo struct {
	add   addFunc
	check checkFunc
}

func registerType(name string, add addFunc, check checkFunc) {
	flagTypes[name] = flagInfo{
		add:   add,
		check: check,
	}
}

// Option is used to customize a flag.
type Option func(*flag)

// Env sets the name of an environment variable to check for the value.
func Env(name string) Option {
	return func(o *flag) {
		o.envVar = name
	}
}

// Default sets the default value of a flag.
func Default(value any) Option {
	return func(o *flag) {
		o.defaultValue = value
	}
}

// Required marks the specified flag as being required.
func Required() Option {
	return func(o *flag) {
		o.required = true
	}
}

// NotRequired marks the specified flag as not being required.
func NotRequired() Option {
	return func(o *flag) {
		o.required = false
	}
}

func Validate(fn ValidationFunc) Option {
	return func(o *flag) {
		o.required = false
	}
}

type ValidationFunc func(val any) error

type flag struct {
	dest         any
	name         string
	envVar       string
	defaultValue any
	usage        string
	required     bool

	validationFn ValidationFunc
}

func (f *flag) Usage() string {
	// TODO: maybe throw in our default value?
	if f.envVar == "" {
		return f.usage
	}

	return helpers.EnvDesc(f.usage, f.envVar)
}

// New sets up a new flag.
func New(p any, name, usage string, opts ...Option) *flag {
	f := flag{
		dest:  p,
		name:  name,
		usage: usage,
	}

	for _, opt := range opts {
		opt(&f)
	}

	return &f
}

type FlagSet struct {
	flags []*flag
}

func (s *FlagSet) Add(fs *pflag.FlagSet, flags ...*flag) {
	for i := range flags {
		t := fmt.Sprintf("%T", flags[i].dest)
		fi, ok := flagTypes[t]
		if !ok {
			panic(fmt.Sprintf("unsupported type %s", t))
		}
		fi.add(fs, flags[i])
	}

	s.flags = append(s.flags, flags...)
}

func (s *FlagSet) Check() error {
	var errs []error

	for _, f := range s.flags {
		if !f.required {
			continue
		}

		t := fmt.Sprintf("%T", f.dest)
		fi, ok := flagTypes[t]
		if !ok {
			panic(fmt.Sprintf("unsupported type %s", t))
		}

		if err := fi.check(f); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

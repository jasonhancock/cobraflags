package flags

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jasonhancock/go-env"
	"github.com/jasonhancock/go-helpers"
	"github.com/spf13/pflag"
)

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

type flag struct {
	dest         any
	name         string
	envVar       string
	defaultValue any
	usage        string
	required     bool
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
		switch t := flags[i].dest.(type) {
		case *time.Duration:
			var defaultVal time.Duration
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(time.Duration)
				if !ok {
					panic(fmt.Sprintf("%s is a time.Duration, but the default value is not a time.Duration", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Duration(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Duration(flags[i].envVar, 0)
				}
			}

			fs.DurationVar(flags[i].dest.(*time.Duration), flags[i].name, defaultVal, flags[i].Usage())
		case *string:
			var defaultVal string
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(string)
				if !ok {
					panic(fmt.Sprintf("%s is a string, but the default value is not a string", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.String(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = os.Getenv(flags[i].envVar)
				}
			}

			fs.StringVar(flags[i].dest.(*string), flags[i].name, defaultVal, flags[i].Usage())

		default:
			panic(fmt.Sprintf("unsupported type %T", t))
		}
	}

	s.flags = append(s.flags, flags...)
}

func (s *FlagSet) Check() error {
	var errs []error

	for _, f := range s.flags {
		if !f.required {
			continue
		}

		switch t := f.dest.(type) {
		case *string:
			val, ok := f.dest.(*string)
			if !ok {
				panic(fmt.Sprintf("%q not a string", f.name))
			}

			if *val == "" {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}
		case *time.Duration:
			val, ok := f.dest.(*time.Duration)
			if !ok {
				panic(fmt.Sprintf("%q not a time.Duration", f.name))
			}

			if *val == 0 {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}

		default:
			panic(fmt.Sprintf("unsupported type %T", t))
		}
	}

	return errors.Join(errs...)
}

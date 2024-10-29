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
		case *bool:
			var defaultVal bool
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(bool)
				if !ok {
					panic(fmt.Sprintf("%s is a bool, but the default value is not a bool", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Bool(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Bool(flags[i].envVar, false)
				}
			}
			fs.BoolVar(flags[i].dest.(*bool), flags[i].name, defaultVal, flags[i].Usage())
		case *int:
			var defaultVal int
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(int)
				if !ok {
					panic(fmt.Sprintf("%s is an int, but the default value is not an int", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Int(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Int(flags[i].envVar, 0)
				}
			}
			fs.IntVar(flags[i].dest.(*int), flags[i].name, defaultVal, flags[i].Usage())
		case *int64:
			var defaultVal int64
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(int64)
				if !ok {
					panic(fmt.Sprintf("%s is an int64, but the default value is not an int64", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Int64(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Int64(flags[i].envVar, 0)
				}
			}

			fs.Int64Var(flags[i].dest.(*int64), flags[i].name, defaultVal, flags[i].Usage())
		case *uint8:
			var defaultVal uint8
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(uint8)
				if !ok {
					panic(fmt.Sprintf("%s is an uint8, but the default value is not an uint8", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Uint8(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Uint8(flags[i].envVar, 0)
				}
			}
			fs.Uint8Var(flags[i].dest.(*uint8), flags[i].name, defaultVal, flags[i].Usage())
		case *uint16:
			var defaultVal uint16
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(uint16)
				if !ok {
					panic(fmt.Sprintf("%s is an uint16, but the default value is not an uint16", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Uint16(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Uint16(flags[i].envVar, 0)
				}
			}
			fs.Uint16Var(flags[i].dest.(*uint16), flags[i].name, defaultVal, flags[i].Usage())
		case *uint32:
			var defaultVal uint32
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(uint32)
				if !ok {
					panic(fmt.Sprintf("%s is an uint32, but the default value is not an uint32", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Uint32(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Uint32(flags[i].envVar, 0)
				}
			}
			fs.Uint32Var(flags[i].dest.(*uint32), flags[i].name, defaultVal, flags[i].Usage())
		case *uint64:
			var defaultVal uint64
			if flags[i].defaultValue != nil {
				var ok bool
				defaultVal, ok = flags[i].defaultValue.(uint64)
				if !ok {
					panic(fmt.Sprintf("%s is an uint64, but the default value is not an uint64", flags[i].name))
				}

				if flags[i].envVar != "" {
					defaultVal = env.Uint64(flags[i].envVar, defaultVal)
				}
			} else {
				if flags[i].envVar != "" {
					defaultVal = env.Uint64(flags[i].envVar, 0)
				}
			}
			fs.Uint64Var(flags[i].dest.(*uint64), flags[i].name, defaultVal, flags[i].Usage())
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
		case *bool:
			_, ok := f.dest.(*bool)
			if !ok {
				panic(fmt.Sprintf("%q not a bool", f.name))
			}

			// Doesn't really make sense to check for false here.
		case *int:
			val, ok := f.dest.(*int)
			if !ok {
				panic(fmt.Sprintf("%q not an int", f.name))
			}
			if *val == 0 {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}
		case *int64:
			val, ok := f.dest.(*int64)
			if !ok {
				panic(fmt.Sprintf("%q not an int64", f.name))
			}
			if *val == 0 {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}
		case *uint8:
			val, ok := f.dest.(*uint8)
			if !ok {
				panic(fmt.Sprintf("%q not an uint8", f.name))
			}
			if *val == 0 {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}
		case *uint16:
			val, ok := f.dest.(*uint16)
			if !ok {
				panic(fmt.Sprintf("%q not an uint16", f.name))
			}
			if *val == 0 {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}
		case *uint32:
			val, ok := f.dest.(*uint32)
			if !ok {
				panic(fmt.Sprintf("%q not an uint32", f.name))
			}
			if *val == 0 {
				errs = append(errs, fmt.Errorf("required value %q not specified", f.name))
			}
		case *uint64:
			val, ok := f.dest.(*uint64)
			if !ok {
				panic(fmt.Sprintf("%q not an uint64", f.name))
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

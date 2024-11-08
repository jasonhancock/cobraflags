package flags

import (
	"fmt"

	"github.com/jasonhancock/go-env"
	"github.com/spf13/pflag"
	"golang.org/x/exp/constraints"
)

func init() {
	registerType("*float32", float32Add, numberCheck[float32])
	registerType("*float64", float64Add, numberCheck[float64])
}

func float32Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal float32
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(float32)
		if !ok {
			panic(fmt.Sprintf("%s is a float32, but the default value is not a float32", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Float32(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Float32(f.envVar, 0)
		}
	}
	fs.Float32Var(f.dest.(*float32), f.name, defaultVal, f.Usage())
}

func float64Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal float64
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(float64)
		if !ok {
			panic(fmt.Sprintf("%s is a float64, but the default value is not a float64", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Float64(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Float64(f.envVar, 0)
		}
	}
	fs.Float64Var(f.dest.(*float64), f.name, defaultVal, f.Usage())
}

func numberCheck[T constraints.Float | constraints.Integer](f *flag) error {
	var zero T
	val, ok := f.dest.(*T)
	if !ok {
		return fmt.Errorf("%q not a %T", f.name, zero)
	}
	if *val == T(0) {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil

}

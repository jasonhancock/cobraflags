package flags

import (
	"fmt"

	"github.com/jasonhancock/go-env"
	"github.com/spf13/pflag"
)

func init() {
	registerType("*bool", boolAdd, boolCheck)
}

func boolAdd(fs *pflag.FlagSet, f *flag) {
	var defaultVal bool
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(bool)
		if !ok {
			panic(fmt.Sprintf("%s is a bool, but the default value is not a bool", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Bool(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Bool(f.envVar, false)
		}
	}
	fs.BoolVar(f.dest.(*bool), f.name, defaultVal, f.Usage())
}

func boolCheck(f *flag) error {
	_, ok := f.dest.(*bool)
	if !ok {
		return fmt.Errorf("%q not a bool", f.name)
	}

	// Doesn't really make sense to check for false here.

	return nil
}

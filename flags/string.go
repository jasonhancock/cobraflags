package flags

import (
	"fmt"
	"os"

	"github.com/jasonhancock/go-env"
	"github.com/spf13/pflag"
)

func init() {
	registerType("*string", stringAdd, stringCheck)
}

func stringAdd(fs *pflag.FlagSet, f *flag) {
	var defaultVal string
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(string)
		if !ok {
			panic(fmt.Sprintf("%s is a string, but the default value is not a string", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.String(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = os.Getenv(f.envVar)
		}
	}
	fs.StringVar(f.dest.(*string), f.name, defaultVal, f.Usage())
}

func stringCheck(f *flag) error {
	val, ok := f.dest.(*string)
	if !ok {
		return fmt.Errorf("%q not a string", f.name)
	}
	if *val == "" {
		return fmt.Errorf("required value %q not specified", f.name)
	}

	return nil
}

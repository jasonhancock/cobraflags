package flags

import (
	"fmt"
	"time"

	"github.com/jasonhancock/go-env"
	"github.com/spf13/pflag"
)

func init() {
	registerType("*time.Duration", durationAdd, durationCheck)
}

func durationAdd(fs *pflag.FlagSet, f *flag) {
	var defaultVal time.Duration
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(time.Duration)
		if !ok {
			panic(fmt.Sprintf("%s is a time.Duration, but the default value is not a time.Duration", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Duration(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Duration(f.envVar, 0)
		}
	}

	fs.DurationVar(f.dest.(*time.Duration), f.name, defaultVal, f.Usage())
}

func durationCheck(f *flag) error {
	val, ok := f.dest.(*time.Duration)
	if !ok {
		return fmt.Errorf("%q not a time.Duration", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}

	return nil
}

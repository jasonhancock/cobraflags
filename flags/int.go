package flags

import (
	"fmt"

	"github.com/jasonhancock/go-env"
	"github.com/spf13/pflag"
)

func init() {
	registerType("*int", intAdd, intCheck)
	registerType("*int8", int8Add, int8Check)
	registerType("*int16", int16Add, int16Check)
	registerType("*int32", int32Add, int32Check)
	registerType("*int64", int64Add, int64Check)

	registerType("*uint", uintAdd, uintCheck)
	registerType("*uint8", uint8Add, uint8Check)
	registerType("*uint16", uint16Add, uint16Check)
	registerType("*uint32", uint32Add, uint32Check)
	registerType("*uint64", uint64Add, uint64Check)
}

func intAdd(fs *pflag.FlagSet, f *flag) {
	var defaultVal int
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(int)
		if !ok {
			panic(fmt.Sprintf("%s is an int, but the default value is not an int", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Int(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Int(f.envVar, 0)
		}
	}
	fs.IntVar(f.dest.(*int), f.name, defaultVal, f.Usage())
}

func int8Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal int8
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(int8)
		if !ok {
			panic(fmt.Sprintf("%s is an int8, but the default value is not an int8", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Int8(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Int8(f.envVar, 0)
		}
	}
	fs.Int8Var(f.dest.(*int8), f.name, defaultVal, f.Usage())
}

func int16Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal int16
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(int16)
		if !ok {
			panic(fmt.Sprintf("%s is an int16, but the default value is not an int16", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Int16(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Int16(f.envVar, 0)
		}
	}
	fs.Int16Var(f.dest.(*int16), f.name, defaultVal, f.Usage())
}

func int32Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal int32
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(int32)
		if !ok {
			panic(fmt.Sprintf("%s is an int32, but the default value is not an int32", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Int32(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Int32(f.envVar, 0)
		}
	}
	fs.Int32Var(f.dest.(*int32), f.name, defaultVal, f.Usage())
}

func int64Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal int64
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(int64)
		if !ok {
			panic(fmt.Sprintf("%s is an int64, but the default value is not an int64", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Int64(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Int64(f.envVar, 0)
		}
	}

	fs.Int64Var(f.dest.(*int64), f.name, defaultVal, f.Usage())
}

func uintAdd(fs *pflag.FlagSet, f *flag) {
	var defaultVal uint
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(uint)
		if !ok {
			panic(fmt.Sprintf("%s is an uint, but the default value is not an uint", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Uint(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Uint(f.envVar, 0)
		}
	}
	fs.UintVar(f.dest.(*uint), f.name, defaultVal, f.Usage())
}

func uint8Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal uint8
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(uint8)
		if !ok {
			panic(fmt.Sprintf("%s is an uint8, but the default value is not an uint8", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Uint8(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Uint8(f.envVar, 0)
		}
	}
	fs.Uint8Var(f.dest.(*uint8), f.name, defaultVal, f.Usage())
}

func uint16Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal uint16
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(uint16)
		if !ok {
			panic(fmt.Sprintf("%s is an uint16, but the default value is not an uint16", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Uint16(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Uint16(f.envVar, 0)
		}
	}
	fs.Uint16Var(f.dest.(*uint16), f.name, defaultVal, f.Usage())
}

func uint32Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal uint32
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(uint32)
		if !ok {
			panic(fmt.Sprintf("%s is an uint32, but the default value is not an uint32", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Uint32(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Uint32(f.envVar, 0)
		}
	}
	fs.Uint32Var(f.dest.(*uint32), f.name, defaultVal, f.Usage())
}

func uint64Add(fs *pflag.FlagSet, f *flag) {
	var defaultVal uint64
	if f.defaultValue != nil {
		var ok bool
		defaultVal, ok = f.defaultValue.(uint64)
		if !ok {
			panic(fmt.Sprintf("%s is an uint64, but the default value is not an uint64", f.name))
		}

		if f.envVar != "" {
			defaultVal = env.Uint64(f.envVar, defaultVal)
		}
	} else {
		if f.envVar != "" {
			defaultVal = env.Uint64(f.envVar, 0)
		}
	}
	fs.Uint64Var(f.dest.(*uint64), f.name, defaultVal, f.Usage())
}

func intCheck(f *flag) error {
	val, ok := f.dest.(*int)
	if !ok {
		return fmt.Errorf("%q not an int", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil
}

func int8Check(f *flag) error {
	val, ok := f.dest.(*int8)
	if !ok {
		return fmt.Errorf("%q not an int8", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}

	return nil
}

func int16Check(f *flag) error {
	val, ok := f.dest.(*int16)
	if !ok {
		return fmt.Errorf("%q not an int16", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}

	return nil
}

func int32Check(f *flag) error {
	val, ok := f.dest.(*int32)
	if !ok {
		return fmt.Errorf("%q not an int32", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}

	return nil
}

func int64Check(f *flag) error {
	val, ok := f.dest.(*int64)
	if !ok {
		return fmt.Errorf("%q not an int64", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}

	return nil
}

func uintCheck(f *flag) error {
	val, ok := f.dest.(*uint)
	if !ok {
		return fmt.Errorf("%q not an uint", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil
}

func uint8Check(f *flag) error {
	val, ok := f.dest.(*uint8)
	if !ok {
		return fmt.Errorf("%q not an uint8", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil
}

func uint16Check(f *flag) error {
	val, ok := f.dest.(*uint16)
	if !ok {
		return fmt.Errorf("%q not an uint16", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil
}

func uint32Check(f *flag) error {
	val, ok := f.dest.(*uint32)
	if !ok {
		return fmt.Errorf("%q not an uint32", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil
}

func uint64Check(f *flag) error {
	val, ok := f.dest.(*uint64)
	if !ok {
		return fmt.Errorf("%q not an uint64", f.name)
	}
	if *val == 0 {
		return fmt.Errorf("required value %q not specified", f.name)
	}
	return nil
}

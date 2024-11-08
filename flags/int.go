package flags

import (
	"fmt"

	"github.com/jasonhancock/go-env"
	"github.com/spf13/pflag"
)

func init() {
	registerType("*int", intAdd, numberCheck[int])
	registerType("*int8", int8Add, numberCheck[int8])
	registerType("*int16", int16Add, numberCheck[int16])
	registerType("*int32", int32Add, numberCheck[int32])
	registerType("*int64", int64Add, numberCheck[int64])

	registerType("*uint", uintAdd, numberCheck[uint])
	registerType("*uint8", uint8Add, numberCheck[uint8])
	registerType("*uint16", uint16Add, numberCheck[uint16])
	registerType("*uint32", uint32Add, numberCheck[uint32])
	registerType("*uint64", uint64Add, numberCheck[uint64])
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

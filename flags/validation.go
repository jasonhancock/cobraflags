package flags

import (
	"cmp"
	"errors"
	"fmt"
)

// Enum provides a validation function where the flag value must be in the set
// of provided values.
func Enum[T comparable](vals ...T) ValidationFunc {
	allowed := make(map[T]struct{}, len(vals))
	for _, v := range vals {
		allowed[v] = struct{}{}
	}
	return func(val any) error {
		str, ok := val.(*T)
		if !ok {
			var zero T
			return fmt.Errorf("value was not a %T", zero)
		}
		if _, ok := allowed[*str]; !ok {
			return errors.New("value not in allowed values")
		}

		return nil
	}
}

// Minimum provides a validation function where the flag value must be greater
// than or equal to the provided minimum value.
func Minimum[T cmp.Ordered](min T) ValidationFunc {
	return func(val any) error {
		str, ok := val.(*T)
		if !ok {
			var zero T
			return fmt.Errorf("value was not a %T", zero)
		}
		if *str < min {
			return errLessThanMin
		}

		return nil
	}
}

// Maximum provides a validation function where the flag value must be less than
// or equal to the proovided max value.
func Maximum[T cmp.Ordered](max T) ValidationFunc {
	return func(val any) error {
		str, ok := val.(*T)
		if !ok {
			var zero T
			return fmt.Errorf("value was not a %T", zero)
		}
		if *str > max {
			return errGreaterThanMax

		}

		return nil
	}
}

var errLessThanMin = errors.New("value is less than minimum value")
var errGreaterThanMax = errors.New("value is greater than maximum value")

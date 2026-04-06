package flags

import (
	"cmp"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/jasonhancock/go-helpers"
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

var errLessThanMin = errors.New("value is less than minimum value")
var errGreaterThanMax = errors.New("value is greater than maximum value")

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

// URL provides a validation function where the flag value must be parseable as a URL.
func URL(opts ...URLValidationOption) ValidationFunc {
	var o urlValidationOptions
	for _, opt := range opts {
		opt(&o)
	}

	return func(val any) error {
		str, ok := val.(*string)
		if !ok {
			return errors.New("value was not a string")
		}

		u, err := url.Parse(*str)
		if err != nil {
			return fmt.Errorf("parsing %q as a URL: %w", *str, err)
		}

		if len(o.validSchemes) > 0 {
			if _, ok := o.validSchemes[u.Scheme]; !ok {
				return fmt.Errorf("%q is not in list of valid schemes: %s", u.Scheme, strings.Join(helpers.KeysSorted(o.validSchemes), "|"))
			}
		}

		return nil
	}
}

type urlValidationOptions struct {
	validSchemes map[string]struct{}
}

// URLValidationOption is used to validate URLs.
type URLValidationOption func(*urlValidationOptions)

// RequireScheme sets a list of acceptable schemes.
func RequireScheme(validSchemes ...string) URLValidationOption {
	return func(o *urlValidationOptions) {
		if o.validSchemes == nil {
			o.validSchemes = make(map[string]struct{})
		}
		for _, v := range validSchemes {
			o.validSchemes[v] = struct{}{}
		}
	}
}

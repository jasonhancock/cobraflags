package flags

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnum(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		v := "green"
		validate := Enum("red", "green", "blue")

		err := validate(&v)
		require.NoError(t, err)
	})

	t.Run("invalid value", func(t *testing.T) {
		v := "yellow"
		validate := Enum("red", "green", "blue")

		err := validate(&v)
		require.Error(t, err)
		assert.EqualError(t, err, "value not in allowed values")
	})

	t.Run("wrong type", func(t *testing.T) {
		v := 42
		validate := Enum("red", "green")

		err := validate(&v)
		require.Error(t, err)
		assert.ErrorContains(t, err, "value was not a")
	})
}

func TestMinimum(t *testing.T) {
	t.Run("value equal to minimum", func(t *testing.T) {
		v := 10
		validate := Minimum(10)

		err := validate(&v)
		require.NoError(t, err)
	})

	t.Run("value greater than minimum", func(t *testing.T) {
		v := 15
		validate := Minimum(10)

		err := validate(&v)
		require.NoError(t, err)
	})

	t.Run("value less than minimum", func(t *testing.T) {
		v := 5
		validate := Minimum(10)

		err := validate(&v)
		require.Error(t, err)
		assert.True(t, errors.Is(err, errLessThanMin))
	})

	t.Run("wrong type", func(t *testing.T) {
		v := "not-an-int"
		validate := Minimum(10)

		err := validate(&v)
		require.Error(t, err)
		assert.ErrorContains(t, err, "value was not a")
	})
}

func TestMaximum(t *testing.T) {
	t.Run("value equal to maximum", func(t *testing.T) {
		v := 10
		validate := Maximum(10)

		err := validate(&v)
		require.NoError(t, err)
	})

	t.Run("value less than maximum", func(t *testing.T) {
		v := 5
		validate := Maximum(10)

		err := validate(&v)
		require.NoError(t, err)
	})

	t.Run("value greater than maximum", func(t *testing.T) {
		v := 15
		validate := Maximum(10)

		err := validate(&v)
		require.Error(t, err)
		assert.True(t, errors.Is(err, errGreaterThanMax))
	})

	t.Run("wrong type", func(t *testing.T) {
		v := "not-an-int"
		validate := Maximum(10)

		err := validate(&v)
		require.Error(t, err)
		assert.ErrorContains(t, err, "value was not a")
	})
}

func TestURL(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		tests := []struct {
			input string
			err   error
		}{
			{"https://example.com", nil},
			{"12345", nil},
			{"", nil},
			{"http://example.com/%zz", errors.New("invalid URL escape")},
			{"http://example.com:ab", errors.New("invalid port")},
			{"http://exa mple.com/helloworld", errors.New("invalid character")},
			{"://example.com", errors.New("missing protocol scheme")},
			{"htt p://example.com", errors.New("first path segment in URL cannot contain colon")},
		}

		fn := URL()

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				result := fn(&tt.input)
				if tt.err == nil {
					require.NoError(t, result)
					return
				}
				require.Error(t, result)
				require.Contains(t, result.Error(), tt.err.Error())
			})
		}
	})

	t.Run("scheme", func(t *testing.T) {
		tests := []struct {
			input string
			err   error
		}{
			{"https://example.com", nil},
			{"file:///foo", nil},
			{"http://example.com", errors.New(`"http" is not in list of valid schemes: file|https`)},
		}

		fn := URL(RequireScheme("https", "file"))

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				result := fn(&tt.input)
				if tt.err == nil {
					require.NoError(t, result)
					return
				}
				require.Error(t, result)
				require.Contains(t, result.Error(), tt.err.Error())
			})
		}
	})
}

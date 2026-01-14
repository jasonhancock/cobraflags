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

package postgresql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		prefix       string
		expectedFlag string
		expectedEnv  string
	}{
		{"", "db", "DB"},
		{"PreFix", "prefix-db", "PREFIX_DB"},
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			o := options{prefix: tt.prefix}

			require.Equal(t, o.flagName("db"), tt.expectedFlag)
			require.Equal(t, o.envName("DB"), tt.expectedEnv)
		})
	}
}

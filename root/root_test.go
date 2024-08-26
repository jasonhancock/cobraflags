package root

import (
	"testing"

	ver "github.com/jasonhancock/cobra-version"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestUserAgent(t *testing.T) {
	cmd := &cobra.Command{
		Use:  "foo",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	r := New(
		"myapp",
		WithVersion(ver.New("1.2.3", "abc123", "2024-01-02")),
		WithCommand(cmd),
	)

	require.Equal(t, "myapp-foo / 1.2.3", r.UserAgent(cmd))
}

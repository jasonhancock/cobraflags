package root

import (
	"bytes"
	"context"
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

type ctxKey struct{}

func TestLoggerContextResolutionFuncs(t *testing.T) {
	r := New("myapp", LoggerEnabled(true))

	var buf bytes.Buffer
	l := r.Logger(
		&buf,
		WithKeyVals("key", "value"),
		WithContextResolutionFunc(func(ctx context.Context) []any {
			if v, ok := ctx.Value(ctxKey{}).(string); ok {
				return []any{"request_id", v}
			}
			return nil
		}),
		WithContextResolutionFunc(func(ctx context.Context) []any {
			return []any{"second", "yes"}
		}),
	)

	l.InfoContext(context.WithValue(context.Background(), ctxKey{}, "abc123"), "hello")
	out := buf.String()
	require.Contains(t, out, "key=value")
	require.Contains(t, out, "request_id=abc123")
	require.Contains(t, out, "second=yes")

	// Sub-loggers inherit the resolvers.
	buf.Reset()
	l.New("sub").InfoContext(context.WithValue(context.Background(), ctxKey{}, "def456"), "hello")
	require.Contains(t, buf.String(), "request_id=def456")
}

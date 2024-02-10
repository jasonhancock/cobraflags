package root_test

import (
	"os"
	"sync"

	ver "github.com/jasonhancock/cobra-version"
	cr "github.com/jasonhancock/cobraflags/root"
	"github.com/spf13/cobra"
)

// These variables are populated by goreleaser when the binary is built.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func Example() {
	r := cr.New(
		"myapp",
		cr.WithVersion(ver.New(version, commit, date)),
		cr.LoggerEnabled(true),
	)

	r.AddCommand(NewCmd(r))

	r.Execute()
}

func NewCmd(root *cr.Command) *cobra.Command {
	return &cobra.Command{
		Use:          "server",
		Short:        "Starts the server.",
		SilenceUsage: true,
		Args:         cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var wg sync.WaitGroup

			wg.Add(1)
			go func() {
				defer wg.Done()
				root.Logger(os.Stdout).Info("hello world")
				<-cmd.Context().Done()
			}()

			wg.Wait()

			return nil
		},
	}
}

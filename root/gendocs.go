package root

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// GenDocs generates markdown documentation from your command definitions into
// the specified output directory.
func GenDocs(r *Command) *cobra.Command {
	return &cobra.Command{
		Use:          "gendocs directory",
		Short:        "Generates CLI Documentation.",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doc.GenMarkdownTree(r.root, args[0])
		},
	}
}

package root

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	ver "github.com/jasonhancock/cobra-version"
	"github.com/spf13/cobra"
)

type Command struct {
	root *cobra.Command
}

func New(use string, opts ...Option) *Command {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	var c Command

	if o.cmd == nil {
		c.root = &cobra.Command{
			Use:   use,
			Short: o.short,
			Long:  o.long,
		}
	} else {
		c.root = o.cmd
	}

	if o.version != nil {
		c.root.AddCommand(ver.NewCmd(*o.version))
	}

	c.root.AddCommand(o.commands...)

	return &c
}

func (c *Command) Execute() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := c.root.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func (c *Command) AddCommand(cmds ...*cobra.Command) {
	c.root.AddCommand(cmds...)
}

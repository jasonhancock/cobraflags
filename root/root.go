package root

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	clog "github.com/jasonhancock/cobra-logger"
	ver "github.com/jasonhancock/cobra-version"
	"github.com/jasonhancock/go-logger"
	"github.com/spf13/cobra"
)

type Command struct {
	root         *cobra.Command
	loggerConfig *clog.Config
	logger       *logger.L
}

func New(use string, opts ...Option) *Command {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	var c Command

	if o.cmd == nil {
		c.root = &cobra.Command{
			Use:           use,
			Short:         o.short,
			Long:          o.long,
			SilenceErrors: true,
		}
	} else {
		c.root = o.cmd
	}

	if o.version != nil {
		c.root.AddCommand(ver.NewCmd(*o.version))
	}

	c.root.AddCommand(o.commands...)

	if o.loggerEnabled {
		c.loggerConfig = clog.NewConfigPflags(
			strings.Fields(use)[0],
			c.root.PersistentFlags(),
		)
	}

	return &c
}

func (c *Command) Execute() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := c.root.ExecuteContext(ctx); err != nil {
		if c.loggerConfig != nil {
			// the logger has been enabled
			if c.logger == nil {
				c.logger = c.loggerConfig.Logger(os.Stderr)
			}
			c.logger.LogError("execution error", err)
		} else {
			fmt.Println(err)
		}
		if ec, ok := err.(ExitCoder); ok {
			os.Exit(ec.ExitCode())
		}
		os.Exit(1)
	}
}

func (c *Command) AddCommand(cmds ...*cobra.Command) {
	c.root.AddCommand(cmds...)
}

func (c *Command) Logger(dest io.Writer, keyvals ...interface{}) *logger.L {
	c.logger = c.loggerConfig.Logger(dest, keyvals...)
	return c.logger
}

// ExitCoder allows for customization of the exit code when an error is
// encountered.
type ExitCoder interface {
	ExitCode() int
}

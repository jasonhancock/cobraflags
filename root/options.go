package root

import (
	ver "github.com/jasonhancock/cobra-version"
	"github.com/spf13/cobra"
)

type options struct {
	cmd           *cobra.Command
	short         string
	long          string
	version       *ver.Info
	commands      []*cobra.Command
	loggerEnabled bool
}

// Option is used to customize the command.
type Option func(*options)

// WithBaseCommand allows you to completely swap out the root command for your own.
func WithBaseCommand(cmd *cobra.Command) Option {
	return func(o *options) {
		o.cmd = cmd
	}
}

// WithShort sets the command's short description.
func WithShort(str string) Option {
	return func(o *options) {
		o.short = str
	}
}

// WithLong sets the command's long description.
func WithLong(str string) Option {
	return func(o *options) {
		o.long = str
	}
}

// WithVersion passes in a vesion structure to use and wires up the version command.
func WithVersion(info *ver.Info) Option {
	return func(o *options) {
		o.version = info
	}
}

// WithCommand adds commands.
func WithCommand(cmd ...*cobra.Command) Option {
	return func(o *options) {
		o.commands = append(o.commands, cmd...)
	}
}

// LoggerEnabled will add peristent flags from cobra-logger to the root command
// and enable a helper method for constructing a logger.
func LoggerEnabled(enabled bool) Option {
	return func(o *options) {
		o.loggerEnabled = enabled
	}
}

type loggerOptions struct {
	name    string
	keyvals []any
}

// LoggerOption is used to customize the logger.
type LoggerOption func(*loggerOptions)

// WithKeyVals adds key/value pairs to the logger.
func WithKeyVals(keyvals ...any) LoggerOption {
	return func(o *loggerOptions) {
		o.keyvals = append(o.keyvals, keyvals...)
	}
}

// WithName sets the logger name.
func WithName(name string) LoggerOption {
	return func(o *loggerOptions) {
		o.name = name
	}
}

package postgresql

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/jasonhancock/cobraflags/flags"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/pflag"
)

type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int

	SSLMode     string
	SSLRootCert string
	SSLKey      string
	SSLCert     string
	SSLInline   bool

	flags.FlagSet
}

func NewConfig(flagSet *pflag.FlagSet, opts ...Option) *Config {
	var o options
	var c Config

	for _, opt := range opts {
		opt(&o)
	}

	c.Add(
		flagSet,

		flags.New(
			&c.Host,
			o.flagName("db-host"),
			"Database hostname or IP address",
			flags.Env(o.envName("DB_HOST")),
			flags.Default("127.0.0.1"),
			flags.Required(),
		),

		flags.New(
			&c.User,
			o.flagName("db-user"),
			"Datatabase username",
			flags.Env(o.envName("DB_USER")),
			flags.Required(),
		),

		flags.New(
			&c.Password,
			o.flagName("db-pass"),
			"Database password",
			flags.Env(o.envName("DB_PASSWORD")),
		),

		flags.New(
			&c.Name,
			o.flagName("db-name"),
			"Database name",
			flags.Env(o.envName("DB_NAME")),
			flags.Required(),
		),

		flags.New(
			&c.Port,
			o.flagName("db-port"),
			"Database port",
			flags.Env(o.envName("DB_PORT")),
			flags.Default(5432),
			flags.Required(),
		),

		flags.New(
			&c.SSLMode,
			o.flagName("db-ssl-mode"),
			"Database SSL mode",
			flags.Env(o.envName("DB_SSL_MODE")),
			flags.Default("disable"),
			flags.Required(),
		),
	)

	return &c
}

// DSN will return the DSN string.
func (cfg *Config) DSN() (string, error) {
	if err := cfg.Check(); err != nil {
		return "", err
	}

	hosts := strings.Split(cfg.Host, ",")
	for i := range hosts {
		hosts[i] = fmt.Sprintf("%s:%d", hosts[i], cfg.Port)
	}

	var auth string
	if cfg.User != "" {
		auth = cfg.User
		if cfg.Password != "" {
			auth += ":" + cfg.Password
		}
		auth += "@"
	}

	data := make(url.Values)
	data.Add("sslmode", cfg.SSLMode)

	if cfg.SSLMode != "disable" {
		// By setting the SSLInline parameter to true we can pass the certificates
		// directly in the connection string instead of writing them to disk.
		if cfg.SSLInline {
			data.Add("sslinline", "true")
		}

		if cfg.SSLMode == "verify-full" || cfg.SSLMode == "require" {
			data.Add("sslrootcert", cfg.SSLRootCert)
		}

		data.Add("sslkey", cfg.SSLKey)
		data.Add("sslcert", cfg.SSLCert)
	}

	return "postgresql://" + auth + strings.Join(hosts, ",") + "/" + cfg.Name + "?" + data.Encode(), nil
}

// Connect attempts to connect to the database.
func (cfg *Config) Connect() (*sqlx.DB, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}

	return sqlx.Connect("postgres", dsn)
}

type options struct {
	prefix string
}

func (o options) flagName(name string) string {
	if o.prefix == "" {
		return name
	}

	return strings.ToLower(o.prefix) + "-" + name
}

func (o options) envName(name string) string {
	if o.prefix == "" {
		return name
	}

	return strings.ToUpper(o.prefix) + "_" + name
}

// Option is used to customize
type Option func(*options)

// WithPrefix sets the prefix name to use for environment variables and flags.
// Useful if your app has to connect to multiple databases.
func WithPrefix(prefix string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

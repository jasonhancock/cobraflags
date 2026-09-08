package postgresql

import (
	"fmt"
	"net/url"
	"strings"
	"time"

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

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration

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
			"Database hostname or IP address.",
			flags.Env(o.envName("DB_HOST")),
			flags.Default("127.0.0.1"),
			flags.Required(),
		),

		flags.New(
			&c.User,
			o.flagName("db-user"),
			"Datatabase username.",
			flags.Env(o.envName("DB_USER")),
			flags.Required(),
		),

		flags.New(
			&c.Password,
			o.flagName("db-pass"),
			"Database password.",
			flags.Env(o.envName("DB_PASSWORD")),
		),

		flags.New(
			&c.Name,
			o.flagName("db-name"),
			"Database name.",
			flags.Env(o.envName("DB_NAME")),
			flags.Required(),
		),

		flags.New(
			&c.Port,
			o.flagName("db-port"),
			"Database port.",
			flags.Env(o.envName("DB_PORT")),
			flags.Default(5432),
			flags.Required(),
		),

		flags.New(
			&c.SSLMode,
			o.flagName("db-ssl-mode"),
			"Database SSL mode.",
			flags.Env(o.envName("DB_SSL_MODE")),
			flags.Default("disable"),
			flags.Required(),
		),

		flags.New(
			&c.SSLCert,
			o.flagName("db-tls-cert"),
			"TLS client certificate.",
			flags.Env(o.envName("DB_TLS_CERT")),
		),

		flags.New(
			&c.SSLKey,
			o.flagName("db-tls-key"),
			"TLS client private key.",
			flags.Env(o.envName("DB_TLS_KEY")),
		),

		flags.New(
			&c.SSLRootCert,
			o.flagName("db-tls-ca-cert"),
			"TLS CA Certificate.",
			flags.Env(o.envName("DB_TLS_CA_CERT")),
		),

		flags.New(
			&c.MaxOpenConns,
			o.flagName("db-max-open-conns"),
			"Max open connections. 0 means unlimited.",
			flags.Env(o.envName("DB_MAX_OPEN_CONNS")),
			flags.Default(0),
		),

		flags.New(
			&c.MaxIdleConns,
			o.flagName("db-max-idle-conns"),
			"Max idle connections retained. 0 retains none; no unlimited value exists, set equal to max-open-conns.",
			flags.Env(o.envName("DB_MAX_IDLE_CONNS")),
			// database/sql's own default. Must be stated explicitly: passing 0 to
			// SetMaxIdleConns means "retain none", not "use the default".
			flags.Default(2),
		),

		flags.New(
			&c.ConnMaxIdleTime,
			o.flagName("db-conn-max-idle-time"),
			"Max time a connection may sit idle before being closed. 0 means no limit.",
			flags.Env(o.envName("DB_CONN_MAX_IDLE_TIME")),
		),

		flags.New(
			&c.ConnMaxLifetime,
			o.flagName("db-conn-max-lifetime"),
			"Max time a connection may be reused. 0 means no limit.",
			flags.Env(o.envName("DB_CONN_MAX_LIFETIME")),
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

	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(cfg.MaxOpenConns)
	conn.SetMaxIdleConns(cfg.MaxIdleConns)
	conn.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	conn.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return conn, nil
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

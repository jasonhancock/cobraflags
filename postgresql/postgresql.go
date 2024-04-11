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

func NewConfig(flagSet *pflag.FlagSet) *Config {
	var c Config

	c.Add(
		flagSet,

		flags.New(
			&c.Host,
			"db-host",
			"Database hostname or IP address",
			flags.Env("DB_HOST"),
			flags.Default("127.0.0.1"),
			flags.Required(),
		),

		flags.New(
			&c.User,
			"db-user",
			"Datatabase username",
			flags.Env("DB_USER"),
			flags.Required(),
		),

		flags.New(
			&c.Password,
			"db-pass",
			"Database password",
			flags.Env("DB_PASSWORD"),
		),

		flags.New(
			&c.Name,
			"db-name",
			"Database name",
			flags.Env("DB_NAME"),
			flags.Required(),
		),

		flags.New(
			&c.Port,
			"db-port",
			"Database port",
			flags.Env("DB_PORT"),
			flags.Default(5432),
			flags.Required(),
		),

		flags.New(
			&c.SSLMode,
			"db-ssl-mode",
			"Database SSL mode",
			flags.Env("DB_SSL_MODE"),
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

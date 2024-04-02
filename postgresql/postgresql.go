package postgresql

import (
	"strconv"
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
			flags.Required(),
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

// Connect attempts to connect to the database.
func (cfg *Config) Connect() (*sqlx.DB, error) {
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	opts := []string{
		"host=" + cfg.Host,
		"port=" + strconv.Itoa(cfg.Port),
		"user=" + cfg.User,
		"dbname=" + cfg.Name,
		"sslmode=" + cfg.SSLMode,
	}

	if cfg.SSLMode != "disable" {
		if cfg.SSLMode == "verify-full" || cfg.SSLMode == "require" {
			opts = append(opts, "sslrootcert="+cfg.SSLRootCert)
		}

		// TODO: we can apparently provide the cert/key as raw strings by setting the
		// sslinline=true parameter. This might be useful if/when loading the cert from
		// something like Vault. See https://github.com/lib/pq/blob/master/ssl.go#L96
		opts = append(opts, "sslkey="+cfg.SSLKey)
		opts = append(opts, "sslcert="+cfg.SSLCert)
	}

	return sqlx.Connect("postgres", strings.Join(opts, " "))
}

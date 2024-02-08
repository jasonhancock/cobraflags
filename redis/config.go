package logger

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jasonhancock/go-env"
	"github.com/jasonhancock/go-helpers"
	"github.com/spf13/cobra"
)

// DefaultIdleTimeout is the default used for the idle timeout parameter.
const DefaultIdleTimeout = 1 * time.Second

const EnvAddr = "REDIS_ADDR"
const EnvDB = "REDIS_DB"
const EnvIdleTimeout = "REDIS_IDLE_TIMEOUT"

type Config struct {
	Addr        string
	DB          int
	IdleTimeout time.Duration
}

func NewConfig(cmd *cobra.Command) *Config {
	var c Config

	cmd.Flags().StringVar(
		&c.Addr,
		"redis-addr",
		env.String(EnvAddr, "127.0.0.1:6379"),
		helpers.EnvDesc("Redis address and port to connect to. Default: 127.0.0.1:6379.", EnvAddr),
	)

	cmd.Flags().IntVar(
		&c.DB,
		"redis-db",
		env.Int(EnvDB, 0),
		helpers.EnvDesc("Redis database number to use. Default: 0.", EnvDB),
	)

	cmd.Flags().DurationVar(
		&c.IdleTimeout,
		"redis-idle-timeout",
		env.Duration(EnvIdleTimeout, DefaultIdleTimeout),
		helpers.EnvDesc("Redis idle timeout to use.", EnvIdleTimeout),
	)

	return &c
}

// Pool gets the Pool.
func (cfg *Config) Pool() *redis.Pool {
	return &redis.Pool{
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", cfg.DB); err != nil {
				cErr := c.Close()
				if cErr != nil {
					err = errors.Join(err, cErr)
				}
				return nil, err
			}
			return c, nil
		},
	}
}

package clickhouse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDSN(t *testing.T) {
	tests := []struct {
		desc     string
		conf     Config
		expected string
	}{
		{
			"with db, with user, with password",
			Config{
				Host:     "localhost",
				User:     "username",
				Password: "password",
				Name:     "mydb",
				Port:     9000,
			},
			"clickhouse://localhost:9000/mydb?password=password&username=username",
		},
		{
			"with user, with password",
			Config{
				Host:     "localhost",
				User:     "username",
				Password: "password",
				Port:     9000,
			},
			"clickhouse://localhost:9000?password=password&username=username",
		},
		{
			"no user, with password",
			Config{
				Host:     "localhost",
				Password: "password",
				Name:     "mydb",
				Port:     9000,
			},
			"clickhouse://localhost:9000/mydb?password=password",
		},
		{
			"no password",
			Config{
				Host: "localhost",
				Name: "mydb",
				Port: 9000,
			},
			"clickhouse://localhost:9000/mydb?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			dsn, err := tt.conf.DSN()
			require.NoError(t, err)
			require.Equal(t, tt.expected, dsn)
		})
	}
}

package config

import (
	"time"

	"github.com/wb-go/wbf/config"
)

type Config struct {
	Server   server
	Postgres postgres
}

type server struct {
	Addr string
}

type postgres struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func New() *Config {
	c := config.New()
	c.EnableEnv("")
	_ = c.LoadEnvFiles(".env", ".env.example")

	cfg := &Config{
		Server: server{
			Addr: c.GetString("server.addr"),
		},
		Postgres: postgres{
			URL:             c.GetString("postgres.url"),
			MaxOpenConns:    c.GetInt("postgres.max_open_conns"),
			MaxIdleConns:    c.GetInt("postgres.max_idle_conns"),
			ConnMaxLifetime: c.GetDuration("postgres.conn_max_lifetime"),
		},
	}

	return cfg
}

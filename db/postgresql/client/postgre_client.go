package client

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/url"
)

type PostgresConfig struct {
	User       string
	Password   string
	Host       string
	Port       string
	DB         string
	QueryItems map[string]interface{}
}

func NewPostgresClient(ctx context.Context, cfg PostgresConfig) (*pgx.Conn, error) {
	values := url.Values{}

	for k, v := range cfg.QueryItems {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
		values.Encode())

	fmt.Println(connStr)

	conn, err := pgx.Connect(ctx, connStr)
	defer conn.Close(ctx)

	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

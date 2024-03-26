package db

import (
	"context"
	"fmt"

	"github.com/iancharters/gostart/internal/db/dao"
	"github.com/jmoiron/sqlx"
)

// Database configuration object
type Config struct {
	DatabaseURL string
}

// A Client is responsible for communicating with the database.
type Client struct {
	config Config
	conn   *sqlx.DB
	*dao.Queries
}

// NewClient returns a new Client from the provided conf.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	conn, err := sqlx.ConnectContext(ctx, "postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	// Auto run migrations on start
	if err := Migrate(ctx, cfg.DatabaseURL); err != nil {
		return nil, fmt.Errorf("run database migrations: %w", err)
	}

	return &Client{
		cfg,
		conn,
		dao.New(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

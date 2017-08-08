package driver

import (
	"database/sql/driver"
	"io"
	"net/http"
	"time"
)

const (
	tokenExpiryDelta    = 10 * time.Second
	tokenExpiryDuration = 60 * time.Minute
)

// Conn represents a connection to a database and implements driver.Conn.
type Conn struct {
	client    *http.Client
	adwordsID string
	opts      *Opts
}

// Close marks this connection as no longer in use.
func (c *Conn) Close() error {
	// Resets client
	c.client = nil
	return nil
}

// Begin is dedicated to start a transaction and awql does not support it.
func (c *Conn) Begin() (driver.Tx, error) {
	return nil, driver.ErrSkip
}

// Prepare returns a prepared statement, bound to this connection.
func (c *Conn) Prepare(q string) (driver.Stmt, error) {
	if q == "" {
		// No query to prepare.
		return nil, io.EOF
	}
	return &Stmt{Db: c, SrcQuery: q}, nil
}

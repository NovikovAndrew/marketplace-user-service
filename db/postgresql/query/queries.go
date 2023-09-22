package query

import "github.com/jackc/pgx/v5"

type Queries struct {
	conn *pgx.Conn
}

func NewQueries(conn *pgx.Conn) Querier {
	return &Queries{
		conn: conn,
	}
}

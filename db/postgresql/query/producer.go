package query

import "context"

type RegisterProducerParams struct {
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

const registerProducer = `
INSERT INTO producer.producers
(phone, email, hashed_password)
VALUES ($1, $2, $3)
RETURNING *
`

func (q *Queries) RegisterProducer(ctx context.Context, args RegisterProducerParams) (Producer, error) {
	row := q.conn.QueryRow(ctx, registerProducer, args.Phone, args.Email, args.HashedPassword)
	var p Producer
	err := row.Scan(
		&p.ProducerID,
		&p.Phone,
		&p.Email,
		&p.HashedPassword,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

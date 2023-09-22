package query

import "context"

type RegisterCourierParams struct {
	ProducerID     int64  `json:"producer_id"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

const registerCourier = `
INSERT INTO courier.couriers
(producer_id, phone, email, hashed_password)
VALUES ($1, $2, $3, $4)
RETURNING *
`

func (q *Queries) RegisterCourier(ctx context.Context, args RegisterCourierParams) (Courier, error) {
	row := q.conn.QueryRow(ctx, registerCourier, args.ProducerID, args.Phone, args.Email, args.HashedPassword)
	var c Courier
	err := row.Scan(
		&c.CourierID,
		&c.Phone,
		&c.Email,
		&c.HashedPassword,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	return c, err
}

package query

import "context"

type RegisterCustomerParams struct {
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

const registerCustomer = `
INSERT INTO customer.customers 
(phone, email, hashed_password)
VALUES ($1, $2, $3)
RETURNING *;
`

func (q *Queries) RegisterCustomer(ctx context.Context, args RegisterCustomerParams) (Customer, error) {
	row := q.conn.QueryRow(ctx, registerCustomer, args.Phone, args.Email, args.HashedPassword)
	var c Customer
	err := row.Scan(
		&c.CustomerID,
		&c.Phone,
		&c.Email,
		&c.HashedPassword,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	return c, err
}

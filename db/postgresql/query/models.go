package query

import "time"

type Customer struct {
	CustomerID     int64     `json:"customer_id"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Producer struct {
	ProducerID     int64     `json:"producer_id"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Courier struct {
	CourierID      int64     `json:"courier_id"`
	ProducerID     int64     `json:"producer_id"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

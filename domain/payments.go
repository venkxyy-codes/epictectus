package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "initiated"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type Payment struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PgPaymentId     string             `json:"pg_payment_id" bson:"pg_payment_id"`
	PaymentProvider PaymentProvider    `json:"payment_provider" bson:"payment_provider"`
	Amount          int64              `json:"amount" bson:"amount"`
	Currency        string             `json:"currency" bson:"currency"`
	Status          PaymentStatus      `json:"status" bson:"status"`
	UserId          int64              `json:"user_id" bson:"user_id"`
}

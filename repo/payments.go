package repo

import (
	"epictectus/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type paymentsRepository struct {
	collection *mongo.Collection
}

type PaymentsRepository interface {
}

func NewPaymentsRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("payments"),
	}
}

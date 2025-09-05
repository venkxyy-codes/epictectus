package repo

import (
	"context"
	"epictectus/config"
	"epictectus/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type userRepository struct {
	collection *mongo.Collection
}

type UserRepository interface {
	AddNewUser(ctx *gin.Context, user *domain.User) error
	GetUserByUserId(ctx *gin.Context, userId int64) (*domain.User, error)
	GetAllUsers(ctx *gin.Context) ([]domain.User, error)
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		collection: db.Database(config.GetConfig().DbConfig.DBName).Collection("users"),
	}
}

func (r *userRepository) AddNewUser(ctx *gin.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) GetUserByUserId(ctx *gin.Context, userId int64) (*domain.User, error) {
	filter := bson.M{"user_id": userId}
	var result domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	return &result, err
}

func (r *userRepository) GetAllUsers(ctx *gin.Context) ([]domain.User, error) {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "user_id", Value: -1}})

	cur, err := r.collection.Find(cctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("find users: %w", err)
	}
	defer func() { _ = cur.Close(cctx) }()

	var users []domain.User
	if err := cur.All(cctx, &users); err != nil {
		return nil, fmt.Errorf("cursor all: %w", err)
	}
	return users, nil
}

package repo

import (
	"epictectus/config"
	"epictectus/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	filter := bson.D{}
	opts := options.Find().SetSort(map[string]interface{}{"user_id": -1})
	cursor, err := r.collection.Find(ctx, filter, opts)
	var users []domain.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}

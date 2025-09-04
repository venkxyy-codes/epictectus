package appcontext

import (
	"context"
	"epictectus/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type appContext struct {
	mongoDbClient *mongo.Client
	redisClient   *redis.Client
}

var appCtx *appContext

func Init() {
	appCtx = &appContext{}
	appCtx.mongoDbClient = newDbClient()
}

func newDbClient() *mongo.Client {
	dbConfig := config.GetConfig().DbConfig
	appName := config.GetConfig().AppName
	if len(appName) == 0 {
		appName = "default"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()

	connectionString := dbConfig.GetConnectionString()

	client, err := mongo.
		Connect(ctx, options.Client().
			ApplyURI(connectionString).
			SetMaxPoolSize(30).
			SetMinPoolSize(10).
			SetMaxConnecting(10).
			SetRetryWrites(false))
	if err != nil {
		log.Fatal(err, "err-connection-could-not-be-created")
		return nil
	}

	// Ping the primary
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err, "err-connection-ping-failed")
		return nil
	}

	log.Print("info-successfully-connected-to-db, ", "connection_string: ", connectionString)
	return client
}

func GetDBClient() *mongo.Client {
	return appCtx.mongoDbClient
}

func GetRedisClient() *redis.Client {
	return appCtx.redisClient
}

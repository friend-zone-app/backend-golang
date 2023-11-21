package database

import (
	"context"
	"log"
	"parties-app/backend/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client          *mongo.Client
	UserCollection  *mongo.Collection
	PostCollection  *mongo.Collection
	EventCollection *mongo.Collection
	Context         context.Context
)

func init() {
	Context, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)

	Client, connectErr := mongo.Connect(Context, options.Client().ApplyURI(config.MONGO_URL))
	if connectErr != nil {
		panic(connectErr)
	}

	defer cancelContext()

	log.Println("Connected to database!")

	UserCollection = Client.Database("FriendZone").Collection("users")
	PostCollection = Client.Database("FriendZone").Collection("posts")
	EventCollection = Client.Database("FriendZone").Collection("events")

	eventLocationModel := mongo.IndexModel{
		Keys: bson.D{{Key: "location", Value: "2dsphere"}},
	}
	_, err := EventCollection.Indexes().CreateOne(context.TODO(), eventLocationModel)
	if err != nil {
		panic(err)
	}

	userLocationModel := mongo.IndexModel{
		Keys: bson.D{{Key: "setting.location.eventAutomation.location", Value: "2dsphere"}},
	}

	_, err = UserCollection.Indexes().CreateOne(context.TODO(), userLocationModel)
	if err != nil {
		panic(err)
	}
}

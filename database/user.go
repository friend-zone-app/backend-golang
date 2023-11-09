package database

import (
	"parties-app/backend/graph/customTypes"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByID(id primitive.ObjectID) (*customTypes.User, bool, error) {
	var user *customTypes.User

	err := UserCollection.FindOne(Context, bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func GetUserByEmail(email string) (*customTypes.User, bool, error) {
	var user *customTypes.User

	err := UserCollection.FindOne(Context, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func GetUserByUsername(username string) (*customTypes.User, bool, error) {
	var user *customTypes.User

	err := UserCollection.FindOne(Context, bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func CreateUser(user customTypes.User) (bool, error) {
	_, err := UserCollection.InsertOne(Context, user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetManyUserID(ids []primitive.ObjectID) (*[]customTypes.SensoredUser, bool, error) {
	var users []customTypes.SensoredUser
	cursor, err := UserCollection.Find(Context, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &users); err != nil {
		return nil, false, err
	}
	return &users, true, nil
}

func UpdateRefreshToken(userId primitive.ObjectID, token string) error {
	_, err := UserCollection.UpdateOne(Context, bson.D{{Key: "_id", Value: userId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "credentials.refreshToken", Value: token}, {Key: "credentials.lastRefreshed", Value: time.Now()}}}})
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(userId primitive.ObjectID, newUser customTypes.UpdateUserArgs) error {
	_, err := UserCollection.UpdateOne(Context, bson.D{{Key: "_id", Value: userId}}, bson.D{{Key: "$set", Value: newUser}})
	if err != nil {
		return err
	}

	return nil
}

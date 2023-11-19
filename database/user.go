package database

import (
	"crypto/rand"
	"math"
	"math/big"
	"parties-app/backend/graph/customTypes"
	"strings"
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

	if len(user.Events) >= int(math.Pow(float64(user.Level), 2)) {
		err := AddLevelUser(id)
		if err != nil {
			return nil, false, err
		}
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

func GetManyUserID(ids []primitive.ObjectID) (*[]customTypes.User, bool, error) {
	var users []customTypes.User
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

func AddEventToUser(userId primitive.ObjectID, eventId primitive.ObjectID) error {
	_, err := UserCollection.UpdateOne(Context, bson.D{{Key: "_id", Value: userId}}, bson.D{{Key: "$push", Value: bson.D{{Key: "events", Value: eventId}}}})
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

func AddLevelUser(userID primitive.ObjectID) error {
	_, err := UserCollection.UpdateOne(Context, bson.D{{Key: "_id", Value: userID}}, bson.M{"$inc": bson.M{"level": 1}})
	if err != nil {
		return err
	}

	return nil
}

func CreateARandomUsername(username string) (string, error) {
	newUsername := strings.ReplaceAll(username, " ", "")

	for {
		randomIdentifier, err := generateRandomIdentifier(3)
		if err != nil {
			return "", err
		}
		randomSeperator, err := getRandomSeparator()
		if err != nil {
			return "", err
		}
		proposedUsername := newUsername + randomSeperator + randomIdentifier
		count, err := UserCollection.CountDocuments(Context, bson.M{"username": proposedUsername})
		if err != nil {
			return "", err
		}
		if count == 0 {
			return proposedUsername, nil
		}
	}
}

func generateRandomIdentifier(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	_, err := rand.Read(result)
	if err != nil {
		return "", err
	}
	for i := range result {
		result[i] = charset[int(result[i])%len(charset)]
	}
	return string(result), nil
}

func getRandomSeparator() (string, error) {
	separators := []string{".", "-", "_"}

	// Generate a random index using crypto/rand
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(separators))))
	if err != nil {
		return "", err
	}

	return separators[randomIndex.Int64()], err
}

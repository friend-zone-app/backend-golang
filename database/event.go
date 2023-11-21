package database

import (
	"parties-app/backend/graph/customTypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEventByID(id primitive.ObjectID) (*customTypes.Event, bool, error) {
	var event *customTypes.Event

	err := EventCollection.FindOne(Context, bson.D{{Key: "_id", Value: id}}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return event, true, nil
}

func GetUserEvent(id primitive.ObjectID) (*customTypes.Event, bool, error) {
	var event *customTypes.Event

	err := EventCollection.FindOne(Context, bson.D{{Key: "author", Value: id}}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return event, true, nil
}

func GetUserEvents(id primitive.ObjectID) ([]*customTypes.Event, error) {
	var event []*customTypes.Event
	cursor, err := EventCollection.Find(Context, bson.D{{Key: "author", Value: id}})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(Context, &event); err != nil {
		return nil, err
	}
	return event, nil
}

func GetLocalEvent(lat float64, long float64) ([]*customTypes.Event, error) {
	events := make([]*customTypes.Event, 5)

	location := bson.D{
		{Key: "type", Value: "Point"},
		{Key: "coordinates", Value: []float64{lat, long}},
	}

	filter := bson.D{
		{
			Key: "location",
			Value: bson.D{
				{
					Key: "$near",
					Value: bson.M{
						"$geometry":    location,
						"$maxDistance": 1000,
					},
				},
			},
		},
	}

	cursor, findErr := EventCollection.Find(Context, filter)
	if findErr != nil {
		return nil, findErr
	}

	if err := cursor.All(Context, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func CreateEvent(event customTypes.Event) (bool, error) {
	data := bson.M{
		"_id":         event.ID,
		"author":      event.Author,
		"title":       event.Title,
		"description": event.Description,
		"location": bson.M{
			"type":        "Point",
			"coordinates": []float64{event.Location.Coordinates[0], event.Location.Coordinates[1]}, // Replace with actual longitude and latitude values
		},
		"startsAt":  event.StartsAt,
		"endsAt":    event.EndsAt,
		"createdAt": event.CreatedAt,
		"inviters":  event.Inviters,
		"type":      event.Type,
		"privacy": bson.M{
			"WhoCanJoin": event.Privacy.WhoCanJoin,
			"WhoCanSee":  event.Privacy.WhoCanSee,
		},
	}

	_, err := EventCollection.InsertOne(Context, data)
	if err != nil {
		return false, err
	}

	err = AddEventToUser(event.Author, event.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetManyEventID(ids []primitive.ObjectID) (*[]customTypes.Event, bool, error) {
	var event []customTypes.Event
	cursor, err := EventCollection.Find(Context, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &event); err != nil {
		return nil, false, err
	}
	return &event, true, nil
}

func UpdateEvent(userId primitive.ObjectID, newEvent customTypes.UpdateEventInput) error {
	_, err := EventCollection.UpdateOne(Context, bson.D{{Key: "_id", Value: userId}}, bson.D{{Key: "$set", Value: newEvent}})
	if err != nil {
		return err
	}

	return nil
}

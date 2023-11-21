package directives

import (
	"context"
	"errors"
	"net/http"
	"parties-app/backend/authentication"
	"parties-app/backend/database"
	"parties-app/backend/errorHandler"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

func AtTheEvent(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	contextUserId := ForContext(ctx)
	if len(contextUserId) < 5 {
		errorHandler.HandleError(ctx, http.StatusNotAcceptable, "The authorization header is not found! (ForContext)")
		return nil, errors.New("the authorization header is not found (ForContext)")
	}
	userId, err := authentication.ConvertUserIDStringToObjectID(contextUserId)
	if err != nil {
		errorHandler.HandleError(ctx, http.StatusNotAcceptable, "The user id provided is invalid! (ConvertUserIDStringToObjectID)")
		return nil, err
	}

	idParam := ctx.Value("ID")
	if idParam != nil {
		eventIdString, ok := idParam.(string)
		if !ok {
			errorHandler.HandleError(ctx, http.StatusNotAcceptable, "ID parameter is not a string")
			return nil, errors.New("ID parameter is not a string")
		}

		eventId, err := authentication.ConvertUserIDStringToObjectID(eventIdString)
		if err != nil {
			errorHandler.HandleError(ctx, http.StatusNotAcceptable, "The event id provided is invalid! (ConvertUserIDStringToObjectID)")
			return nil, err
		}

		count, err := database.EventCollection.CountDocuments(database.Context, bson.D{{Key: "_id", Value: eventId}, {Key: "inviters.inviteTo", Value: userId}, {Key: "inviters.status", Value: true}})
		if count < 1 {
			errorHandler.HandleError(ctx, http.StatusUnauthorized, "The user not part of that event (EventCollection.CountDocuments)")
			return nil, err
		}

		return next(ctx)

	} else {
		errorHandler.HandleError(ctx, http.StatusNotAcceptable, "ID parameter not found in the context")
		return nil, errors.New("ID parameter not found in the context")
	}
}

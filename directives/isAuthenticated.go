package directives

import (
	"context"
	"net/http"
	auth "parties-app/backend/authentication"
	"parties-app/backend/errorHandler"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

var userCtxKey = &contextKey{"accessToken"}

type contextKey struct {
	name string
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	request := graphql.GetOperationContext(ctx)
	headers := request.Headers

	header := headers.Get("Authorization")
	token, err := auth.ValidateHeaders(header)
	if err != nil {
		return nil, errorHandler.ValidateErrorMessage(http.StatusBadRequest, *err)
	}

	claims, valid, parseErr := auth.ParseAccessToken(*token)
	if !valid && parseErr != nil {
		if claims != nil && claims.ExpiresAt.Unix() > time.Now().Unix() {
			errorHandler.HandleError(ctx, http.StatusRequestTimeout, errorHandler.RefreshToken)
		}
		return nil, errorHandler.ValidateErrorMessage(http.StatusUnauthorized, *parseErr)
	}
	newContext := context.WithValue(ctx, userCtxKey, claims.ID)

	return next(newContext)
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}

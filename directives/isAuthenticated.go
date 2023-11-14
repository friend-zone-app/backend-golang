package directives

import (
	"context"
	"fmt"
	"net/http"
	auth "parties-app/backend/authentication"
	"parties-app/backend/errorHandler"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

var userCtxKey = &contextKey{"accessToken"}

type contextKey struct {
	name string
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	var newContext context.Context

	request := graphql.GetOperationContext(ctx)
	headers := request.Headers

	header := headers.Get("Authorization")
	token, err := auth.ValidateHeaders(header)
	if err != nil {
		return nil, errorHandler.ValidateErrorMessage(http.StatusBadRequest, *err)
	}

	claims, valid, parseErr := auth.ParseAccessToken(*token)
	if parseErr != nil {
		fmt.Println("Directive", parseErr.Error(), strings.HasPrefix(parseErr.Error(), "token is expired"))
		if strings.HasPrefix(parseErr.Error(), "token is expired") {
			return nil, errorHandler.ValidateErrorMessage(http.StatusRequestTimeout, "The token is expired, please login again!")
		}
		return nil, errorHandler.ValidateErrorMessage(http.StatusUnauthorized, "Invalid token")
	}

	if !valid {
		return nil, errorHandler.ValidateErrorMessage(http.StatusUnauthorized, "Invalid token")
	} else {
		newContext = context.WithValue(ctx, userCtxKey, claims.ID)
	}

	return next(newContext)
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}

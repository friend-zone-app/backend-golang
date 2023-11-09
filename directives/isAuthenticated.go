package directives

import (
	"context"
	"net/http"
	auth "parties-app/backend/authentication"
	"parties-app/backend/errorHandler"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	headers := graphql.GetOperationContext(ctx).Headers

	header := headers.Get("Authorization")
	token, err := auth.ValidateHeaders(header)
	if err != nil {
		return nil, errorHandler.ValidateErrorMessage(http.StatusBadRequest, *err)
	}

	claims, valid, parseErr := auth.ParseAccessToken(*token)
	if !valid && parseErr != nil {
		if claims.ExpiresAt.Unix() > time.Now().Unix() {
			errorHandler.HandleError(ctx, http.StatusRequestTimeout, errorHandler.RefreshToken)
		}
		return nil, errorHandler.ValidateErrorMessage(http.StatusUnauthorized, *parseErr)
	}

	return next(ctx)
}

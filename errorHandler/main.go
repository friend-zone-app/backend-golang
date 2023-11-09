package errorHandler

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
)

var (
	InvalidMethod            = "Invalid authorization method. Please use Bearer instead."
	AuthorizationKeyNotFound = "Authorization header not found, please contact support!"
	InvalidToken             = "Invalid token, please check your credentials again!"
	FailedTokenValidation    = "Failed to validate token, please login again!"
	InvalidFormBody          = "Invalid form body, please contact support!"
	InternalServerError      = "Internal server error, please try again later!"
	PostNotFound             = "No posts found for that user!"
	UsernameExisted          = "Username already existed"
	InvalidID                = "The User ID provided is invalid or not found!"
	RefreshToken             = "Access token expired, please refresh the token!"
)

func HandleError(ctx context.Context, statusCode int, message string) {
	graphql.AddErrorf(ctx, fmt.Sprintf("[%v] %v", statusCode, message))
}

func ValidateErrorMessage(statusCode int, message string) error {
	return fmt.Errorf("[%v] %v", statusCode, message)
}

func Unauthorized(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
}

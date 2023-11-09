package errorHandler

import (
	"fmt"
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

func HandleError(statusCode int, message string) error {
	return fmt.Errorf("[%v] %v", statusCode, message)
}

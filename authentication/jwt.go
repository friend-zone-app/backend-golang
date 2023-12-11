package authentication

import (
	"fmt"
	"parties-app/backend/config"
	"parties-app/backend/errorHandler"
	"parties-app/backend/graph/customTypes"
	"parties-app/backend/types"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	RefreshKey = []byte(config.RF_JWT_SECRET)
	AccessKey  = []byte(config.JWT_SECRET)
)

func Sign(userId primitive.ObjectID) (*customTypes.Tokens, error) {

	acClaim := &types.SignedDetails{
		ID: ValidateUserID(userId.Hex()),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			Issuer:    config.IDENTIFY_KEY,
		},
	}

	rfClaim := &types.SignedDetails{
		ID: ValidateUserID(userId.Hex()),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(14 * 24 * time.Hour)),
			Issuer:    config.IDENTIFY_KEY,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, acClaim)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rfClaim)

	actString, err := accessToken.SignedString(AccessKey)
	if err != nil {
		return nil, err
	}
	rfString, err := refreshToken.SignedString(RefreshKey)
	if err != nil {
		return nil, err
	}

	tokens := customTypes.Tokens{
		AccessToken:  actString,
		RefreshToken: rfString,
	}

	return &tokens, nil
}

func ParseAccessToken(tokenStr string) (*types.SignedDetails, bool, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &types.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return AccessKey, nil
	})
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "token is expired") {
			fmt.Println(err)
		}
		return nil, false, err
	}

	claims := token.Claims.(*types.SignedDetails)

	return claims, token.Valid, nil
}

func ParseRefreshToken(tokenStr string) (*types.SignedDetails, bool, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &types.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})
	if err != nil {
		return nil, false, err
	}

	claims := token.Claims.(*types.SignedDetails)

	return claims, token.Valid, nil
}

func ValidateHeaders(header string) (*string, *string) {
	if !strings.HasPrefix(header, "Bearer") {
		return nil, &errorHandler.InvalidMethod
	}

	tokenString := strings.Split(header, " ")[1]
	return &tokenString, nil
}

func ConvertUserIDStringToObjectID(userIDString string) (*primitive.ObjectID, error) {
	if strings.HasPrefix(userIDString, "ObjectID") {
		userIDString = strings.Replace(userIDString, "(", "", 1)
		userIDString = strings.Replace(userIDString, ")", "", 1)
		idWithoutPrefix := strings.TrimPrefix(userIDString, "ObjectID")
		objectID, err := primitive.ObjectIDFromHex(idWithoutPrefix)
		if err != nil {
			fmt.Println("Error converting string to ObjectID:", err)
			return nil, err
		}
		fmt.Println("Converted ObjectID:", objectID)
		return &objectID, nil
	} else {
		objectID, err := primitive.ObjectIDFromHex(userIDString)
		if err != nil {
			fmt.Println("Error converting string to ObjectID:", err)
			return nil, err
		}
		fmt.Println("Converted ObjectID:", objectID)
		return &objectID, nil
	}
}

func ValidateUserID(userIDString string) string {
	if strings.HasPrefix(userIDString, "ObjectID") {
		idWithoutPrefix := strings.TrimPrefix(userIDString, "ObjectID")
		return idWithoutPrefix
	} else {
		return userIDString
	}
}

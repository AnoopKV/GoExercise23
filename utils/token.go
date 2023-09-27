package utils

import (
	"time"

	"github.com/AnoopKV/GoExercise23/entities"
	"github.com/golang-jwt/jwt"
)

// Generate token during login
func GenerateToken(email string, firstName string, lastName string, id string, userType string) (string, error) {
	claims := &entities.SignedDetails{
		Email:      email,
		First_name: firstName,
		Uid:        id,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(GetEnvVal("SECRET_KEY")))
	return token, err
}

// Validate passed token during api call
func ValidateToken(signedToken string) (*entities.SignedDetails, string) {
	if token, err := jwt.ParseWithClaims(
		signedToken,
		&entities.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(GetEnvVal("SECRET_KEY")), nil
		},
	); err != nil {
		return nil, err.Error()
	} else {
		claims, ok := token.Claims.(*entities.SignedDetails)
		if !ok {
			return nil, err.Error()
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return nil, err.Error()
		}
		return claims, ""
	}
}

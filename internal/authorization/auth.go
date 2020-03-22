package auth

import (
	mgerror "authorization-service/pkg/errors"
	"time"

	gologger "github.com/jamolpe/go-logger"

	"authorization-service/pkg/models"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type (
	// ValidationResult validation specification result
	ValidationResult struct {
		IsValid bool
		Expired bool
		Error   bool
	}
)

func createUserToken(email string) (string, error) {
	mySigningKey := []byte("mytopSecret")

	createdAtInt := int64(time.Second)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Subject:   email,
		Issuer:    "gosessioner",
		IssuedAt:  createdAtInt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", mgerror.NewError("authorization error")
	}
	return ss, nil
}

// CheckTokenIsValid checks if the token is valid
func CheckTokenIsValid(tokenString string) ValidationResult {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mytopSecret"), nil
	})

	if token.Valid {
		return ValidationResult{IsValid: true, Expired: false, Error: false}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return ValidationResult{IsValid: false, Expired: false, Error: true}
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			gologger.ERROR("TokenValidation: token expired")
			return ValidationResult{IsValid: false, Expired: true, Error: false}
		} else {
			gologger.ERROR("TokenValidation: couldn't handle this token:" + err.Error())
			return ValidationResult{IsValid: false, Expired: false, Error: true}
		}
	}
	return ValidationResult{IsValid: false, Expired: false, Error: true}
}

// Authorization : check if the user is authoriced
func Authorization(authUser *models.User, requestedUser *models.User) (*models.AuthUser, error) {
	logerUser := &models.AuthUser{}
	if authUser.Password == requestedUser.Password {
		token, err := createUserToken(authUser.Email)
		logerUser.User = authUser
		logerUser.Logged = true
		logerUser.Token = token
		return logerUser, err
	}
	return logerUser, nil
}
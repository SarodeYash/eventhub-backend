package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const Securekey = "Yash"

func GeerateJWTToken(email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //this is the method used to generate a token (Takes header which specidies which algo is used to generate token)
		"email":   email, //Takes Payload as an arguments
		"user_id": user_id,
		"Expo":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(Securekey)) //this method is used to integrate header.payload and signature with securekey and return a JWT token

}

func ValidateJWTToken(token string) (int64, error) { //this function used to parse a token from the memory and compare it with the token shared by client
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { //Takes token as an arg and a function literal to return a secure key and a error
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //the SigningMethod of both the tokens are same or not is being checked
		if !ok {
			return nil, errors.New("invalid token method")
		}
		return []byte(Securekey), nil
	})
	if err != nil { //if the function returns a error
		return 0, errors.New("invalid Token")
	}
	validToken := parsedToken.Valid //token.valid is used to check if the token is valid or not(true/false)
	if !validToken {
		return 0, err
	}
	claim, ok := parsedToken.Claims.(jwt.MapClaims) //if required We can also access the payload or claims data of the token in this way
	if !ok {
		return 0, errors.New("invalid Token Claim")
	}
	// Email:=claim["email"]
	User_id := int64(claim["user_id"].(float64)) //type assertion
	return User_id, nil

}

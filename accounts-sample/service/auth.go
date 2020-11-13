package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// VerifyJWT verifies the token and returns a map of claims
func VerifyJWT(t string) (*jwt.MapClaims, error) {
	segments := strings.Split(t, ".")

	if len(segments) != 3 {
		return nil, fmt.Errorf("[VerifyJWT]: invalid token, token must have three segments; found %d", len(segments))
	}

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[VerifyJWT]: unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("BANK_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		err := fmt.Errorf("invalid token cannot map claims")
		return nil, err
	}

	return &claims, nil
}

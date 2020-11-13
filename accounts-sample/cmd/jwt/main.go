package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	key := os.Getenv("BANK_KEY")

	teamNum := flag.Int64("team", 0, "enter the number for the team (required)")
	generateToken := flag.Bool("g", false, "g flag to generate a new token for the give team")
	checkToken := flag.Bool("c", false, "c flag to validate token for the give team")
	tokenString := flag.String("token", "", "token to validate")

	flag.Parse()

	if *teamNum <= 0 || *teamNum > 9 {
		log.Println("You must provide a team number 1-9 to generate (-c) or validate (-g) a token")
		os.Exit(1)
	}

	team := fmt.Sprintf("Team-%d", *teamNum)

	if *generateToken {
		jwtString, err := generateJWT(team, []byte(key))
		if err != nil {
			log.Fatal("Unable to generate JWT: ", err)
		}
		fmt.Printf("token:%s\n", jwtString)
	} else if *checkToken {
		_, err := validateJWT(*tokenString, team)
		if err != nil {
			log.Fatalf("Token not valid: %v\n", err)
		}
		log.Println("Token is Valid for ", team)
	} else {
		log.Println("You must pass either a -g (generate) or -c (validate)")
	}
}

func validateJWT(tokenString, team string) (*jwt.Token, error) {
	segments := strings.Split(tokenString, ".")

	if len(segments) != 3 {
		return nil, fmt.Errorf("[VerifyToken]: invalid token, token must have three segments; found %d", len(segments))
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[VerifyToken]: unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("BANK_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		err := fmt.Errorf("invalid token")
		return nil, err
	}

	checkTeam := claims["team"].(string)
	if checkTeam != team {
		return nil, fmt.Errorf("Not a valid token for %s", checkTeam)
	}

	return token, nil
}

func generateJWT(team string, key []byte) (string, error) {
	var err error
	var tokenString string

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["team"] = team
	claims["aud"] = os.Getenv("AUDIENCE")
	claims["exp"] = time.Now().Add(time.Hour * 24 * 60).Unix()

	tokenString, err = token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

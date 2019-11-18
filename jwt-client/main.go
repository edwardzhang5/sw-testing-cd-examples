package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	_ = godotenv.Load()

	fmt.Println("Generate JWT Token for OBS")

	obsKey := os.Getenv("OBS_KEY")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		ts, _ := GenerateJWT([]byte(obsKey))
		c.String(http.StatusOK, ts)
	})

	r.Run(":8989")
	// ts, err := GenerateJWT([]byte(obsKey))
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }

	// log.Println("TokenString: ", ts)
}

// GenerateJWT generates a jwt
func GenerateJWT(obsKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = "drbyronw-2019"
	claims["user"] = "Byron Williams"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(obsKey)
	if err != nil {
		log.Fatalf("Cannot gnerate token: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()
  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

  api := router.Group("/api")
  api.Use(cors.New(config))
	{
		api.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "SW Testing API")
		})
		api.GET("/email", EmailHandler)
		api.GET("/distance", DistanceHandler)
		api.GET("/retire", RetirementHandler)
		api.GET("/bmi", BMIHandler)
		api.GET("/tip", TipHandler)
	}

  router.Run(":9090")
}

// Check panic if error is present
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

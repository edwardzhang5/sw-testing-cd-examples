package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DBHandler persists the mongo client for use in Handlers
type DBHandler struct {
	Client *mongo.Client
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Unable to Connect")
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Could not Ping the DB")
		panic(err)
	}

	var dbh DBHandler
	dbh.Client = client

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := SetupRouter(&dbh)
	go r.Run(":5000")

	for {
		fmt.Println("Welcome to the HPC Platform")
		fmt.Println("Enter Your Selection Below (1, 2, 3, 4)")
		fmt.Println("1 - Body Mass Index")
		fmt.Println("2 - Retirement Calculator")
		fmt.Println("3 - Distance Formula")
		fmt.Println("4 - Email Verifier")
		fmt.Println("5 - Split Tip")
		fmt.Println("0 - Exit Program")
		fmt.Print("Enter Selection: ")

		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadByte()
		Check(err)

		switch choice {
		case '1':
			_ = BMIInterface(client)
		case '2':
			RetirementInterface(client)
		case '3':
			_ = DistanceInterface()
		case '4':
			EmailInterface()
		case '5':
			TipInterface()
		case '0':
			os.Exit(0)
		default:
			fmt.Println("Invalid input, Exiting...")
			os.Exit(0)
		}
	}
}

// SetupRouter gets the gin router for app and testing
func SetupRouter(dbh *DBHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "SW Testing")
	})
	r.GET("/bmi/:feet/:inches/:weight", dbh.BMIHandler)
	r.GET("/bmidata", dbh.BMIEndpoint)
	r.GET("/retiredata", dbh.RetireEndpoint)
	r.GET("/retire", dbh.RetireHandler)
	return r
}

// Check panic if error is present
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// BuildTimeStamp creates a readable TS from the time passed
func BuildTimeStamp(t time.Time) string {
	return t.Format(time.UnixDate) + fmt.Sprintf("-%d", t.Nanosecond())
}

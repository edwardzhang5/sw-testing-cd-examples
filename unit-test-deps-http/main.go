package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

// DBHandler persists the mongo client for use in Handlers
type DBHandler struct {
	Session *mgo.Session
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("Unable to Connect to Mongo instance")
		panic(err)
  }
  defer session.Close()
	err = session.Ping()
	if err != nil {
		log.Fatal("Could not Ping the DB")
		panic(err)
	}

	var dbh DBHandler
	dbh.Session = session

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
			_ = BMIInterface(session)
		case '2':
			RetirementInterface(session)
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
	r.GET("/retire", dbh.RetireHandler)
	r.GET("/retiredata", dbh.RetireEndpoint)
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

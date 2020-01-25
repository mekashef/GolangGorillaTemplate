package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// envHandler -- struct allocated to manage the client we recieve
type envHandler struct {
	PORT   string `string:"PORT"`
	DBHOST string `string:"DBHOST"`
	DBPORT string `string:"DBPORT"`
}

//ENVFILE -- constant string name for env file
var ENVFILE = "apiEnv.env"

//Handler -- used to hold
var Handler envHandler

func main() {
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	//Paths that can be handled
	router.GET("/", homePage)

	router.Run(":" + Handler.PORT)
}

func init() {
	fmt.Println("Getting Environment Variables...")

	//Load environmenatal variables
	err := godotenv.Load(ENVFILE)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Setting up information needed to run server...")

	err = setupHandler(&Handler)

	if err != nil {
		log.Fatal("Failed to setup server values")
	}

	fmt.Println("Server Running at Port: " + Handler.PORT)
}

// setupHandler -- Initialize handler variable to be populated
//
// Error handling still needs to be implmented here
func setupHandler(h *envHandler) error {
	// Populate from Env
	h.PORT = os.Getenv("serverPort")
	h.DBHOST = os.Getenv("databaseHost")
	h.DBPORT = os.Getenv("databasePort")

	return nil
}

//homePage -- base landing page to test restAPI is up
func homePage(c *gin.Context) {
	fmt.Println("Endpoint Hit: homePage")
}

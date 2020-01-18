package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

// dbHandler -- struct allocated to manage the client we recieve
type dbHandler struct {
	Client *mongo.Client `json:"Client"`
	PORT   string        `string:"PORT"`
	DBHOST string        `string:"DBHOST"`
	DBPORT string        `string:"DBPORT"`
}
var ENVFILE = "apiEnv.env"
//Handler -- used to hold our database connection to mongoDB and other useful info
var Handler dbHandler

func main() {
	router := mux.NewRouter()

	//Paths that can be handled
	router.HandleFunc("/", homePage).Methods("GET")

	//Start server and log fatal if it fails
	log.Fatal(http.ListenAndServe(
		":"+Handler.PORT,
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))
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
func setupHandler(h *dbHandler) error {
	// Populate from Env
	h.PORT = os.Getenv("serverPort")
	h.DBHOST = os.Getenv("databaseHost")
	h.DBPORT = os.Getenv("databasePort")

	return nil
}

//homePage -- base landing page to test restAPI is up
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to a Golang RestAPI by Mehdi!")
	fmt.Println("Endpoint Hit: homePage")
}
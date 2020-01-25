package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

	router.Run(":10000")
}

//homePage -- base landing page to test restAPI is up
func homePage(c *gin.Context) {
	fmt.Println("Endpoint Hit: homePage")
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to the HomePage!")
	fmt.Println("Endpoint: homePage")
}

func testEndPoint(context *gin.Context) {
	id := context.DefaultQuery("id", "id void")
	name := context.DefaultQuery("name", "name void")

	context.String(http.StatusOK, "test id=%s  name=%s", id, name)

	fmt.Println("Endpoint: test")
}

func handleRequests() {

	// Instantiate a new router
	router := gin.Default()

	// Add a handler on /
	router.GET("/", homePage)

	// Add a handler on /test
	router.GET("/test", testEndPoint)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	handleRequests()
}

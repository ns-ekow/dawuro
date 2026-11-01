package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []user{
	{ID: "001", Username: "ekow", Email: "ekowsackey@gmail.com", Password: "1234"},
	{ID: "002", Username: "nii", Email: "niisa@gmail.com", Password: "1234"},
	{ID: "003", Username: "saki", Email: "sakisa@gmail.com", Password: "1234"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUsersByID)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUsersByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

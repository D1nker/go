package main

import (
	"errors"

	"github.com/gin-gonic/gin"

	"net/http"
)

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{Id: "1", Title: "My First Book", Author: "Quenta", Quantity: 1},
	{Id: "2", Title: "My Second Book", Author: "Alex", Quantity: 2},
	{Id: "3", Title: "My Third Book", Author: "Manu", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func findBookById(id string) (*book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	book, err := findBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/books", getBooks)

	router.GET("/books/:id", getBook)

	router.Run("localhost:8080")
}

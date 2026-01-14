package handler

import (
	"fmt"
	"gin-rest/book"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type bookHanlder struct {
	bookService book.ServiceBook
}

func NewBookHandler(bookService book.ServiceBook) *bookHanlder {
	return &bookHanlder{bookService}
}

func (handler *bookHanlder) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "harisriyoni",
		"bio":  "akjkad",
	})
}

func (handler *bookHanlder) HelloHanlder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "adkjadjkn",
		"subtitle": "adkjadjkn",
	})
}

func (handler *bookHanlder) Bookshandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (handler *bookHanlder) QueryHelper(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (handler *bookHanlder) PostBookHandler(c *gin.Context) {
	var bookInput book.BookRequest
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}

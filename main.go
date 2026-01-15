package main

import (
	"gin-rest/book"
	"gin-rest/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gin-rest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	// bookFileRepositry := book.NewFileRepository()
	bookService := book.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	router.GET("/books", bookHandler.GetBooks)
	router.GET("/books/:id", bookHandler.GetBook)
	router.POST("/books", bookHandler.CreateBook)
	router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8080")
}

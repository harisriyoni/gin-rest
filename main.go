package main

import (
	"gin-rest/book"
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
	bookService := book.NewBookService(bookRepository)
	bookRequest := book.BookRequest{
		Title: "yayaya",
		Price: "9000",
		// Description: "asjkdnjkas",
		// Rating:      5,
	}

	bookService.Create(bookRequest)
	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Tittle:", book.Title)
	// }

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
}

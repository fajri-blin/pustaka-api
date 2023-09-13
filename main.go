package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/initiliazer"
	"pustaka-api/middleware"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initiliazer.LoadEnvVariables()
	db, err = initiliazer.ConnectToDatabase()
	err = initiliazer.SyncDatabase(db)
	if err != nil {
		log.Fatal("Connection Database Failed")
	}
}

func main() {

	//Adding Repository
	bookRepository := book.NewRepository(db)
	userRepository := user.NewRepository(db)

	//Adding Service
	bookService := book.NewService(bookRepository)
	userSerice := user.NewService(userRepository)

	//Adding Handler
	BookHandler := handler.NewBookHandler(bookService)
	UserHandler := handler.NewUserHandler(userSerice)

	//Routing
	router := gin.Default()

	//Routing Grouping

	routerV1 := router.Group("/v1", middleware.RequireAuth)
	routerUser := router.Group("/user", middleware.RequireAuth)

	routerV1.POST("/books", BookHandler.PostBooksHandler)
	routerV1.GET("/books", BookHandler.GetBooksHandler)
	routerV1.GET("/books/:id", BookHandler.GetBookByIdHandler)
	routerV1.PUT("/books/:id", BookHandler.UpdateBookHandler)
	routerV1.DELETE("/books/:id", BookHandler.DeleteBookHandler)

	routerUser.POST("/signup", UserHandler.SignUp)
	routerUser.POST("/login", UserHandler.Login)
	routerUser.GET("/books", BookHandler.GetBooksByUserHandler)

	router.Run()

}

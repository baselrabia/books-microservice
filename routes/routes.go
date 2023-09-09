package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/baselrabia/book-microservice/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Books endpoints
    r.POST("/books", handlers.CreateBook)
    r.PUT("/books/:id", handlers.UpdateBook)
    r.DELETE("/books", handlers.DeleteAllBooks)
    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBookByID)

    return r
}

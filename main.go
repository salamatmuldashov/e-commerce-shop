package main

import (
	"e-commerce-golang/database"
	"e-commerce-golang/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	rtr := gin.Default()

	auth := rtr.Group("/auth")
	{
		auth.POST("/sign-up", handlers.Registration)
		auth.POST("/sign-in", handlers.Authorisation)
	}

	api := rtr.Group("/api")
	{
		items := api.Group("/items")
		{
			items.GET("", handlers.GetItems)    // get all items
			items.GET("/sort", handlers.Filter) // Filtering items based on price, rating
			items.POST("", handlers.AddItem)    // Publishing item

			items.GET("/:id", handlers.GetItem)                            // get item by id
			items.GET("/:id/with_comments", handlers.GetItemsWithComments) // get item with comments
			items.DELETE("/:id", handlers.DeleteItem)                      // Purchasing item

			items.GET("/:id/comments", handlers.GetCommentaries) // get item commments
			items.POST("/:id/comments", handlers.SetComment)     //Commenting items + Giving rating for items

			////get all comments
			items.GET("/comments", handlers.GetComments) // just get all comments
		}
	}
	rtr.Run("localhost:8088")

}

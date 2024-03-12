package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanugi/rest-api-assignment2/controllers"
)

func RouteServer(c *controllers.Controller) *gin.Engine {

	router := gin.Default()

	router.POST("/orders", c.CreateOrder)
	router.GET("/orders", c.GetOrders)
	router.PUT("/orders/:orderId", c.UpdateOrder)
	router.DELETE("/orders/:orderId", c.DeleteOrder)

	return router
}

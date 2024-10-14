package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Rules
	server.GET("/rules") //ListRules
	server.GET("/rules/:ruleName") // ReadRule
	server.POST("/rules") // AddRule
	server.PUT("/rules") // UpdateRule
	server.DELETE("/rules/:ruleName") // DeleteRule

	// Sliding window logs
	server.GET("/sliding-window-logs/:requestID")
	server.POST("/sliding-window-logs/:requestID")
	server.PUT("/sliding-window-logs/:requestID")
	server.DELETE("/sliding-window-logs/:requestID")
}
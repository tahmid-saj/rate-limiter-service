package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Rules
	server.GET("/rules", getRules) //ListRules
	server.GET("/rules/:ruleName", readRule) // ReadRule
	server.POST("/rules", addRule) // AddRule
	server.PUT("/rules", updateRule) // UpdateRule
	server.DELETE("/rules/:ruleName", deleteRule) // DeleteRule

	// Sliding window logs
	server.GET("/sliding-window-logs/:requestID") //ReadRequest
	server.POST("/sliding-window-logs/:requestID") //SendRequest (rate limiting)
	server.PUT("/sliding-window-logs/:requestID") //UpdateRequest
	server.DELETE("/sliding-window-logs/:requestID") //DeleteRequest
}
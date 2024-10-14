package main

import (
	"os"
	"rate-limiter-service/routes"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// CreateTable
	// createdTable, err := dynamodb.CreateTable("rate-limiter-rules")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(createdTable)

	// ListTables
	// tableNames, err := dynamodb.ListTables()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(tableNames)

	// AddRule
	// rule := dynamodb.Rule{
	// 	RuleName: "chat-sigma-api-chat-message",
	// 	ParamName: "send-message",
	// 	Limit: 1,
	// 	WindowInterval: 1,
	// 	WindowTime: "minute",
	// }
	// addedRule, err := dynamodb.AddRule(rule, "rate-limiter-rules")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(addedRule)

	// ReadItem
	// readRule, err := dynamodb.ReadRule("chat-sigma-api-chat-message", "rate-limiter-rules")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(readRule)

	// UpdateRule
	// rule := dynamodb.Rule{
	// 	RuleName: "chat-sigma-api-chat-message",
	// 	ParamName: "send-message",
	// 	Limit: 2,
	// 	WindowInterval: 1,
	// 	WindowTime: "minute",
	// }
	// updadedRule, err := dynamodb.UpdateRule(rule, "rate-limiter-rules")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(updadedRule)

	// DeleteRule
	// deletedRule, err := dynamodb.DeleteRule("chat-sigma-api-chat-message", "rate-limiter-rules")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(deletedRule)

	// Sliding window ops
	// CreateSlidingWindowLogTable
	// createdTable, err := dynamodb.CreateSlidingWindowLogTable("rate-limiter-sliding-window-logs")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(createdTable)

	// AddRequest
	// request := dynamodb.SlidingWindowLogRequest{
	// 	RequestID: "1044eb44-c478-4caa-ad8a-0ec43a4a8495",
	// 	LogRequests: []dynamodb.LogRequests{
	// 		dynamodb.LogRequests{
	// 			Timestamp: time.Now(),
	// 			RuleName: "chat-sigma-api-chat-message",
	// 			ParamName: "send-message",
	// 		},
	// 	},
	// }
	// addedRequest, err := dynamodb.AddRequest(request, "rate-limiter-sliding-window-logs")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(addedRequest)

	// ReadRequest
	// readRequest, err := dynamodb.ReadRequest("1044eb44-c478-4caa-ad8a-0ec43a4a8495", "rate-limiter-sliding-window-logs")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(readRequest)

	// UpdateRequest
	// request := dynamodb.SlidingWindowLogRequest{
	// 	RequestID: "1044eb44-c478-4caa-ad8a-0ec43a4a8495",
	// 	LogRequests: []dynamodb.LogRequests{
	// 		dynamodb.LogRequests{
	// 			Timestamp: time.Now(),
	// 			RuleName: "chat-sigma-api-chat-message",
	// 			ParamName: "send-message",
	// 		},
	// 	},
	// }
	// updatedRequest, err := dynamodb.UpdateRequest(request, "rate-limiter-sliding-window-logs")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(updatedRequest)

	// DeleteRequest
	// deletedRequest, err := dynamodb.DeleteRequest("1044eb44-c478-4caa-ad8a-0ec43a4a8495", "rate-limiter-sliding-window-logs")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(deletedRequest)

	// SendRequest
	// requestOk, err := slidingwindow.SendRequest(
	// 	"1044eb44-c478-4caa-ad8a-0ec43a4a8495", 
	// 	"chat-sigma-api-chat-message", 
	// 	"send-message", 
	// 	"rate-limiter-sliding-window-logs",
	// 	"rate-limiter-rules",
	// )
	// if err != nil {
	// 	return
	// }
	// fmt.Print(requestOk)

	godotenv.Load()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(os.Getenv("PORT"))
}
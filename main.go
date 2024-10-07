package main

import (
	"fmt"
	"os"

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

	// UpdateItem
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

	// DeleteItem
	// deletedRule, err := dynamodb.DeleteRule("chat-sigma-api-chat-message", "rate-limiter-rules")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(deletedRule)

	godotenv.Load()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(os.Getenv("PORT"))
}
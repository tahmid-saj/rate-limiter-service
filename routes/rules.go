package routes

import (
	"net/http"
	"rate-limiter-service/models"

	"github.com/gin-gonic/gin"
)

func getRules(context *gin.Context) {
	res, err := models.ListRules()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not list rules"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func readRule(context *gin.Context) {
	ruleName := context.Param("ruleName")

	res, err := models.ReadRule(ruleName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not read rule"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func addRule(context *gin.Context) {
	var ruleInput models.RuleInput

	err := context.ShouldBindJSON(&ruleInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.AddRule(ruleInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not add rule"})
		return
	}

	context.JSON(http.StatusCreated, res)
}

func updateRule(context *gin.Context) {
	var ruleInput models.RuleInput

	err := context.ShouldBindJSON(&ruleInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.UpdateRule(ruleInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update rule"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func deleteRule(context *gin.Context) {
	ruleName := context.Param("ruleName")

	res, err := models.DeleteRule(ruleName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete rule"})
		return
	}

	context.JSON(http.StatusOK, res)
}
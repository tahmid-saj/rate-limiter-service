package routes

import (
	"net/http"
	"rate-limiter-service/models"

	"github.com/gin-gonic/gin"
)

func readRequest(context *gin.Context) {
	requestID := context.Param("requestID")

	var readRequestInput models.ReadRequestInput

	err := context.ShouldBindJSON(&readRequestInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.ReadRequest(requestID, readRequestInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not read request"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func sendRequest(context *gin.Context) {
	requestID := context.Param("requestID")

	var sendRequestInput models.SendRequestInput

	err := context.ShouldBindJSON(&sendRequestInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.SendRequest(requestID, sendRequestInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not send request"})
		return
	}

	context.JSON(http.StatusCreated, res)
}

func updateRequest(context *gin.Context) {
	requestID := context.Param("requestID")

	var updateRequestInput models.UpdateRequestInput

	err := context.ShouldBindJSON(&updateRequestInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.UpdateRequest(requestID, updateRequestInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update request"})
		return
	}

	context.JSON(http.StatusBadRequest, res)
}

func deleteRequest(context *gin.Context) {
	requestID := context.Param("requestID")

	var deleteRequestInput models.DeleteRequestInput

	err := context.ShouldBindJSON(&deleteRequestInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.DeleteRequest(requestID, deleteRequestInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update request"})
		return
	}

	context.JSON(http.StatusOK, res)
}
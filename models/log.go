package models

import (
	"rate-limiter-service/dynamodb"
	slidingwindow "rate-limiter-service/sliding-window"
)

type ReadRequestInput struct {
	TableName string `json:"tableName"`
}

type SendRequestInput struct {
	RuleName                   string `json:"ruleName"`
	ParamName                  string `json:"paramName"`
	SlidingWindowLogsTableName string `json:"slidingWindowLogsTableName"`
	RulesTableName             string `json:"rulesTableName"`
}

type UpdateRequestInput struct {
	SlidingWindowLogRequests []dynamodb.LogRequests `json:"slidingWindowLogRequest"`
	SlidingWindowLogsTableName string `json:"slidingWindowLogsTableName"`
}

type DeleteRequestInput struct {
	SlidingWindowLogsTableName string `json:"slidingWindowLogsTableName"`
}

type SendRequestResponse struct {
	IsRequestOk bool `json:"isRequestOk"`
}

func ReadRequest(requestID string, readRequestInput ReadRequestInput) (*Response, error) {
	readRequest, err := dynamodb.ReadRequest(requestID, readRequestInput.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: readRequest,
	}, nil
}

func SendRequest(requestID string, sendRequestInput SendRequestInput) (*Response, error) {
	requestOk, err := slidingwindow.SendRequest(
		requestID, 
		sendRequestInput.RuleName, 
		sendRequestInput.ParamName, 
		sendRequestInput.SlidingWindowLogsTableName, 
		sendRequestInput.RulesTableName,
	)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	if !requestOk {
		return &Response{
			Ok: true,
			Response: SendRequestResponse{
				IsRequestOk: false,
			},
		}, nil
	}

	return &Response{
		Ok: true,
		Response: SendRequestResponse{
			IsRequestOk: true,
		},
	}, nil
}

func UpdateRequest(requestID string, updateRequestInput UpdateRequestInput) (*Response, error) {
	updatedRequest, err := dynamodb.UpdateRequest(dynamodb.SlidingWindowLogRequest{
		RequestID: requestID,
		LogRequests: updateRequestInput.SlidingWindowLogRequests,
	}, updateRequestInput.SlidingWindowLogsTableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: updatedRequest,
	}, nil
}

func DeleteRequest(requestID string, deleteRequestInput DeleteRequestInput) (*Response, error) {
	deletedRequest, err := dynamodb.DeleteRequest(requestID, deleteRequestInput.SlidingWindowLogsTableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: deletedRequest,
	}, nil
}
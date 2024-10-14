package models

import "rate-limiter-service/dynamodb"

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
	SlidingWindowLogRequest dynamodb.SlidingWindowLogRequest `json:"slidingWindowLogRequest"`
	SlidingWindowLogsTableName string `json:"slidingWindowLogsTableName"`
}

type DeleteRequestInput struct {
	SlidingWindowLogsTableName string `json:"slidingWindowLogsTableName"`
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
	
}

func UpdateRequest(requestID string, updateRequestInput UpdateRequestInput) (*Response, error) {

}

func DeleteRequest(requestID string, deleteRequestInput DeleteRequestInput) (*Response, error) {

}
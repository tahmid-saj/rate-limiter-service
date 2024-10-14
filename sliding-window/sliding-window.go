package slidingwindow

import (
	"log"
	"rate-limiter-service/dynamodb"
	"rate-limiter-service/utils"
	"time"
)

func SendRequest(requestID, ruleName, paramName, slidingWindowLogsTableName, rulesTableName string) (bool, error) {
	// check if the requestID exists in dynamodb
	requests, err := dynamodb.ReadRequest(requestID, slidingWindowLogsTableName)
	if err != nil {
		// if the requestID doesn't exist, add it to the requests log in dynamodb and return true
		logRequest := dynamodb.SlidingWindowLogRequest{
			RequestID: requestID,
			LogRequests: []dynamodb.LogRequests{
				dynamodb.LogRequests{
					Timestamp: time.Now(),
					RuleName: ruleName,
					ParamName: paramName,
				},
			},
		}

		_, err := dynamodb.AddRequest(logRequest, slidingWindowLogsTableName)
		if err != nil {
			log.Println(err)
			return false, err
		}

		return true, nil
	}

	// if the requestID exists, check if the request should be rate limited or not
	rule, err := dynamodb.ReadRule(ruleName, rulesTableName)
	if err != nil {
		log.Println(err)
		return false, err
	}

	var requestsInSlidingWindow int
	for _, requestLog := range requests.LogRequests {
		isRequestBetween, err := utils.IsTimeBetween(
			requestLog.Timestamp, 
			time.Now().Add(time.Duration((-1 * rule.WindowInterval) * utils.RATE_LIMITER_WINDOW_TIME_MAPPINGS[rule.WindowTime])), 
			time.Now())
		if err != nil {
			log.Println(err)
			return false, err
		}
		
		if isRequestBetween {
			requestsInSlidingWindow += 1
		}

		// if the request should be rate limited, drop it
		if requestsInSlidingWindow >= rule.Limit {
			return false, nil
		}
	}

	// otherwise add the request to the requests log
	request := dynamodb.SlidingWindowLogRequest{
		RequestID: requestID,
		LogRequests: append(requests.LogRequests, dynamodb.LogRequests{
			Timestamp: time.Now(),
			RuleName: ruleName,
			ParamName: paramName,
		}),
	}
	_, err = dynamodb.AddRequest(request, slidingWindowLogsTableName)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
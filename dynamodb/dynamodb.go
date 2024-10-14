package dynamodb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Rule struct {
	RuleName string
	ParamName string
	Limit int
	WindowInterval int
	WindowTime string
}

type SlidingWindowLogRequest struct {
	RequestID string
	LogRequests []LogRequests
}

type LogRequests struct {
	Timestamp time.Time
	RuleName string
	ParamName string
}

func ListTables() ([]string, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}

	var tableNames []string
	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				log.Print(err)
				return nil, err
			}
			log.Print(err)
			return nil, err
		}

		for _, tableName := range result.TableNames {
			tableNames = append(tableNames, *tableName)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}

	return tableNames, nil
}

// Rule operations

func CreateRuleTable(tableName string) (*dynamodb.CreateTableOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create the table input with "Rule" as the primary key (HASH)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("RuleName"), // Define Rule attribute
				AttributeType: aws.String("S"),      // Rule is a string (S)
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("RuleName"), // Primary key (HASH)
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	// Create the table
	result, err := svc.CreateTable(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

func AddRule(item Rule, tableName string) (*dynamodb.PutItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	attributeValue, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

func AddItemsFromJSON(items []interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	var result *dynamodb.PutItemOutput
	for _, item := range items {
    attributeValue, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
			log.Print(err)
			return nil, err
    }

    // Create item in table Movies
    input := &dynamodb.PutItemInput{
			Item:      attributeValue,
			TableName: aws.String(tableName),
    }

    result, err = svc.PutItem(input)
    if err != nil {
			log.Print(err)
			return nil, err
    }
	}

	return result, nil
}

// Get table items from JSON file
func getItems(fileName string) interface{} {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Print(err)
		return err
	}

	var items interface{}
	json.Unmarshal(raw, &items)
	return items
}

func ReadRule(ruleName, tableName string) (*Rule, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
			"RuleName": {
				S: aws.String(ruleName),
			},
    },
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if result.Item == nil {
    msg := "Could not find '" + ruleName + "'"
    return nil, errors.New(msg)
	}
			
	var item *Rule

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return item, nil
}

func UpdateRule(updatedValue Rule, tableName string) (*dynamodb.UpdateItemOutput, error) {
	// Initialize a session that the SDK will use to load credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Prepare the update expression and the attribute values
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"RuleName": {
				S: aws.String(updatedValue.RuleName),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#PN": aws.String("ParamName"),
			"#L":  aws.String("Limit"),
			"#WT": aws.String("WindowTime"),
			"#WI": aws.String("WindowInterval"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pn": {
				S: aws.String(updatedValue.ParamName),
			},
			":l": {
				N: aws.String(strconv.Itoa(updatedValue.Limit)), // Convert int to string
			},
			":wt": {
				S: aws.String(updatedValue.WindowTime),
			},
			":wi": {
				N: aws.String(strconv.Itoa(updatedValue.WindowInterval)), // Convert int to string
			},
		},
		UpdateExpression: aws.String("SET #PN = :pn, #L = :l, #WT = :wt, #WI = :wi"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	// Execute the update
	result, err := svc.UpdateItem(input)
	if err != nil {
		log.Printf("Failed to update item: %v", err)
		return nil, err
	}

	return result, nil
}

func DeleteRule(ruleName, tableName string) (*dynamodb.DeleteItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
    Key: map[string]*dynamodb.AttributeValue{
			"RuleName": {
				S: aws.String(ruleName),
			},
    },
    TableName: aws.String(tableName),
	}

	result, err := svc.DeleteItem(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SlidingWindowLog operations

func CreateSlidingWindowLogTable(tableName string) (*dynamodb.CreateTableOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create the table input with "RequestID" as the primary key (HASH)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("RequestID"), // Define RequestID attribute
				AttributeType: aws.String("S"),         // RequestID is a string (S)
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("RequestID"), // Primary key (HASH)
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	// Create the table
	result, err := svc.CreateTable(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}


func AddRequest(item SlidingWindowLogRequest, tableName string) (*dynamodb.PutItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	attributeValue, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

func ReadRequest(requestID, tableName string) (*SlidingWindowLogRequest, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
			"RuleName": {
				S: aws.String(requestID),
			},
    },
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if result.Item == nil {
    msg := "Could not find '" + requestID + "'"
    return nil, errors.New(msg)
	}
			
	var item *SlidingWindowLogRequest

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return item, nil
}

func UpdateRequest(updatedValue SlidingWindowLogRequest, tableName string) (*dynamodb.UpdateItemOutput, error) {
	// Initialize a session that the SDK will use to load credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Prepare the update expression and the attribute values
	// LogRequests will be serialized as a list of maps
	logRequestsAttributeValue := &dynamodb.AttributeValue{
		L: make([]*dynamodb.AttributeValue, len(updatedValue.LogRequests)),
	}

	for i, logRequest := range updatedValue.LogRequests {
		logRequestsAttributeValue.L[i] = &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{
				"Timestamp": {
					S: aws.String(logRequest.Timestamp.Format(time.RFC3339)), // Format time as string
				},
				"RuleName": {
					S: aws.String(logRequest.RuleName),
				},
				"ParamName": {
					S: aws.String(logRequest.ParamName),
				},
			},
		}
	}

	// Prepare the UpdateItemInput
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"RequestID": {
				S: aws.String(updatedValue.RequestID), // Partition key
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#LR": aws.String("LogRequests"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":lr": logRequestsAttributeValue,
		},
		UpdateExpression: aws.String("SET #LR = :lr"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	// Execute the update
	result, err := svc.UpdateItem(input)
	if err != nil {
		log.Printf("Failed to update item: %v", err)
		return nil, err
	}

	return result, nil
}

func DeleteRequest(requestID, tableName string) (*dynamodb.DeleteItemOutput, error) {
	// Initialize a session that the SDK will use to load credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Prepare the DeleteItemInput with RequestID as the key
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"RequestID": {
				S: aws.String(requestID), // Partition key
			},
		},
		TableName: aws.String(tableName),
	}

	// Execute the delete
	result, err := svc.DeleteItem(input)
	if err != nil {
		log.Printf("Failed to delete item: %v", err)
		return nil, err
	}

	return result, nil
}

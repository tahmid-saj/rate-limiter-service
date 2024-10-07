package dynamodb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

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

func CreateTable(tableName string) (*dynamodb.CreateTableOutput, error) {
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
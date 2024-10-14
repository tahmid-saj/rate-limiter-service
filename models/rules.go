package models

import (
	"rate-limiter-service/dynamodb"
	"rate-limiter-service/utils"
)

type RuleInput struct {
	RuleName       string `json:"ruleName"`
	ParamName      string `json:"paramName"`
	Limit          int    `json:"limit"`
	WindowInterval int    `json:"windowInterval"`
	WindowTime     string `json:"windowTime"`
}

func ListRules() (*Response, error) {
	return nil, nil
}

func ReadRule(ruleName string) (*Response, error) {
	readRule, err := dynamodb.ReadRule(ruleName, utils.RATE_LIMITER_RULES_TABLE_NAME)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: readRule,
	}, nil
}

func AddRule(ruleInput RuleInput) (*Response, error) {
	rule := dynamodb.Rule{
		RuleName: ruleInput.RuleName,
		ParamName: ruleInput.ParamName,
		Limit: ruleInput.Limit,
		WindowInterval: ruleInput.WindowInterval,
		WindowTime: ruleInput.WindowTime,
	}

	addedRule, err := dynamodb.AddRule(rule, utils.RATE_LIMITER_RULES_TABLE_NAME)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: addedRule,
	}, nil
}

func UpdateRule(ruleInput RuleInput) (*Response, error) {
	rule := dynamodb.Rule{
		RuleName: ruleInput.RuleName,
		ParamName: ruleInput.ParamName,
		Limit: ruleInput.Limit,
		WindowInterval: ruleInput.WindowInterval,
		WindowTime: ruleInput.WindowTime,
	}

	updatedRule, err := dynamodb.UpdateRule(rule, utils.RATE_LIMITER_RULES_TABLE_NAME)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: updatedRule,
	}, nil
}

func DeleteRule(ruleName string) (*Response, error) {
	deletedRule, err := dynamodb.DeleteRule(ruleName, utils.RATE_LIMITER_RULES_TABLE_NAME)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: deletedRule,
	}, nil
}
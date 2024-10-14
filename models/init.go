package models

type Response struct {
	Ok       bool        `json:"ok"`
	Response interface{} `json:"response"`
}
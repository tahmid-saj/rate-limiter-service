package utils

import "time"

const RATE_LIMITER_RULES_TABLE_NAME = "rate-limiter-rules"

var RATE_LIMITER_WINDOW_TIME_MAPPINGS = map[string]int{
	"second": int(time.Second),
	"minute": int(time.Minute),
	"hour": int(time.Hour),
	"day": int(time.Hour * 24),
	"week": int(time.Hour * 24 * 7),
	"month": int(time.Hour * 24 * 7 * 30),
}
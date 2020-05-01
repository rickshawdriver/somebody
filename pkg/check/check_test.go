package check

import (
	"testing"
)

const (
	ENDPOINT = "http://localhost:1080"

	METHOD = "POST"
)

func TestChecker(t *testing.T) {
	checkers := Add("hello,world", ENDPOINT, "", METHOD, 5)
	checkers.RunCheck()
}

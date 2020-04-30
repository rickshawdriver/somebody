package check

import (
	"testing"
)

const ENDPOINT = "http://localhost:1080"

func TestChecker(t *testing.T) {
	checkers := Add("hello,world", ENDPOINT, "", 5)
	checkers.RunCheck()
}

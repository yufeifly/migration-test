package utils

import (
	"fmt"
	"testing"
)

func TestParseRange(t *testing.T) {
	input := "key{1:10000}"
	r, err := ParseRange(input)
	if err != nil {
		t.Errorf("parse range failed: %v", err)
	} else {
		fmt.Printf("range: %v", r)
	}
}

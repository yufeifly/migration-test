package utils

import (
	"fmt"
	"testing"
)

func TestParseAddress(t *testing.T) {
	addr, err := ParseAddress("192.176.1.0:1")
	if err != nil {
		t.Errorf("parse err: %v", err)
	}
	fmt.Printf("ip: %v, port:%v\n", addr.IP, addr.Port)
}

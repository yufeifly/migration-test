package utils

import (
	"fmt"
	"github.com/yufeifly/validator/cuserr"
	"regexp"
	"strings"
)

const notFound = -1

type Address struct {
	IP   string
	Port string
}

// ParseAddress 127.0.0.1:6789 -> ip, port. Both ip and port should exist, or return an error
func ParseAddress(raw string) (Address, error) {
	if raw == "" {
		return Address{}, cuserr.ErrEmptyAddress
	}
	addr := Address{}
	var ip, port string

	colonInd := strings.Index(raw, ":")
	if colonInd == notFound { // not found colon, means no port
		return Address{}, cuserr.ErrBadAddress
	}

	ip = raw[:colonInd]
	port = raw[colonInd+1:]

	matchedIP, err := regexp.MatchString("^((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", ip)
	if err != nil {
		return Address{}, err
	}
	if !matchedIP {
		return Address{}, cuserr.ErrBadAddress
	}
	addr.IP = ip

	portMatched, err := regexp.MatchString("^[1-9]\\d*$", port)
	if err != nil {
		return Address{}, err
	}
	if !portMatched {
		return Address{}, cuserr.ErrBadAddress
	}
	addr.Port = port

	return addr, nil
}

// BuildAddress build address, e.g. 127.0.0.1:8888
func BuildAddress(ip, port string) string {
	return fmt.Sprintf("%s:%s", ip, port)
}

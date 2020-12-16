package utils

import (
	"github.com/yufeifly/validator/cuserr"
	"strconv"
	"strings"
)

type Range struct {
	Name       string
	Start, End int
}

// ParseRange key{1:10000}
func ParseRange(r string) (Range, error) {
	if r == "" {
		return Range{}, cuserr.ErrBadParams
	}
	leftSquareBracket := strings.Index(r, "{")
	rightSquareBracket := strings.Index(r, "}")
	colon := strings.Index(r, ":")
	if leftSquareBracket == notFound || rightSquareBracket == notFound || colon == notFound ||
		!(leftSquareBracket < colon && colon < rightSquareBracket) {
		return Range{}, cuserr.ErrBadParams
	}
	name := r[:leftSquareBracket]
	start, err := strconv.Atoi(r[leftSquareBracket+1 : colon])
	if err != nil {
		return Range{}, err
	}
	end, err := strconv.Atoi(r[colon+1 : rightSquareBracket])
	if err != nil {
		return Range{}, nil
	}
	if start > end {
		return Range{}, cuserr.ErrBadParams
	}
	return Range{
		Name:  name,
		Start: start,
		End:   end,
	}, nil
}

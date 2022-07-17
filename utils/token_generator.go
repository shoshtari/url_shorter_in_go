package utils

import (
	"strconv"
)

func shortLinkGenOuter() func() string {
	start := 0
	return func() string {
		start++
		return strconv.Itoa(start)
		//return start
	}
}

var ShortLinkGen = shortLinkGenOuter()

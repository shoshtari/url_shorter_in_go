package utils

import (
	"errors"
	"strconv"
)

var prev = make(map[string]string)

func shortLinkGenOuter() func() string {
	start := 0
	return func() string {
		start++
		return strconv.Itoa(start)
		//return start
	}
}

var shortLinkGen = shortLinkGenOuter()

func TokenGenerator(fullLink string) (err error) {
	_, ok := prev[fullLink]
	if ok {
		err = errors.New("link already exists")
		return
	}
	prev[fullLink] = shortLinkGen()
	return
}

package main

import (
	"regexp"
	"strings"
)

func Parse(s string) ([]string, error) {
	comments, err := regexp.Compile("#.*")
	if err != nil {
		panic(err)
	}

	edgeWs, err := regexp.Compile("(^\\s+|\\s+$)+")
	if err != nil {
		return nil, err
	}

	whiteSpace, err := regexp.Compile("\\s+")
	if err != nil {
		return nil, err
	}

	repeatComma, err := regexp.Compile(",{2,}")
	if err != nil {
		return nil, err
	}

	userStrings, err := regexp.Compile("\"([[:graph:]]|\\s)+\"")
	if err != nil {
		return nil, err
	}

	placeHolder := "__PLACE_HOLDER__"
	phReg, err := regexp.Compile(placeHolder)
	if err != nil {
		return nil, err
	}

	us := userStrings.FindString(s)
	s = userStrings.ReplaceAllString(s, placeHolder)
	s = comments.ReplaceAllString(s, "")
	s = edgeWs.ReplaceAllString(s, "")
	s = whiteSpace.ReplaceAllString(s, ",")
	s = repeatComma.ReplaceAllString(s, ",")
	res := strings.Split(s, ",")
	for i, arg := range res {
		res[i] = phReg.ReplaceAllString(arg, us)
	}
	return res, nil
}

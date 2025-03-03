package main

import "regexp"

const (
	PrefixVersionCode = "versionCode"
	//
	ErrEmptyRegexpResult = "empty regexp result"
	//
	ErrfInvalidMatchResult = "invalid match result: %s"
	ErrfInvalidVersionCode = "invalid versionCode: %s, %s"
	ErrfEmptyRegexpResult  = "empty regexp result: %s"
)

var (
	RegexpVersionCode     = regexp.MustCompile(`(versionCode[\s]{1,}[0-9]{1,})`)
	byteSpace             = []byte(" ")
	bytePrefixVersionCode = []byte(PrefixVersionCode)
)

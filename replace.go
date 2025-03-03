package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func Extract(b []byte) (old int, e error) {

	finds := RegexpVersionCode.FindAll(b, -1)
	if len(finds) < 1 {
		e = errors.New(ErrEmptyRegexpResult)
		return
	}
	find := finds[0]
	find = bytes.TrimSpace(find)
	if !bytes.HasPrefix(find, bytePrefixVersionCode) {
		e = fmt.Errorf(ErrfEmptyRegexpResult, find)
		return
	}

	pair := bytes.Split(find, byteSpace)
	if len(pair) < 2 {
		e = fmt.Errorf(ErrfInvalidMatchResult, find)
		return
	}

	sd := string(pair[1])
	d, err := strconv.Atoi(sd)
	if err != nil {
		e = fmt.Errorf(ErrfInvalidVersionCode, sd, err.Error())
		return
	}
	old = d
	return
}

func Replace(b []byte, old int, new int) []byte {
	oldS := fmt.Sprintf("%s %d", PrefixVersionCode, old)
	newS := fmt.Sprintf("%s %d", PrefixVersionCode, new)
	return bytes.Replace(b, []byte(oldS), []byte(newS), 1)
}

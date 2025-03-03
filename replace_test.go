package main

import (
	"bytes"
	"testing"
)

func TestExtract(t *testing.T) {

	b := []byte(`
		testversiontest
		versionCode 15
		testversiontest
	`)
	actual, err := Extract(b)
	if err != nil {
		t.Errorf("got: error: %+v", err)
	}
	expected := 15
	if actual != expected {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}

func TestReplace(t *testing.T) {

	oldB := []byte(`
		testversiontest
		versionCode 15
		testversiontest
	`)
	actual := Replace(oldB, 15, 16)

	expected := []byte(`
		testversiontest
		versionCode 16
		testversiontest
	`)
	if !bytes.Equal(actual, expected) {
		t.Errorf("got: actual: %v, expected: %v", actual, expected)
	}

}

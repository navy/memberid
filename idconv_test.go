package main

import (
	"testing"
)

type convertIdTest struct {
	Expected string
	Id       string
	From     string
	To       string
}

var config, err = LoadJson("sample.json")

var convertIdTests = []convertIdTest{
	{"jojo", "jojo", "", ""},
	{"jojo", "jotarok", "github", ""},
	{"jotarok", "jojo", "", "github"},
	{"jotaro.kujo", "jotarok", "github", "facebook"},
	{"jojo", "jotarok", "github", "google"},
}

func TestConvertId(t *testing.T) {
	for _, d := range convertIdTests {
		r := ConvertId(d.Id, config, d.From, d.To)
		if r != d.Expected {
			t.Errorf("ConvertId Failed: `-from %s -to %s %s` -> %s (expect: %s)", d.From, d.To, d.Id, r, d.Expected)
		}
	}
}

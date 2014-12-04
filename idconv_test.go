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

var r, err = LoadConfig("sample.json")

var convertIdTests = []convertIdTest{
	{"jojo", "jojo", "", ""},
	{"jojo", "jotarok", "github", ""},
	{"jotarok", "jojo", "", "github"},
	{"jotaro.kujo", "jotarok", "github", "facebook"},
	{"jojo", "jotarok", "github", "google"},
}

func TestConvertId(t *testing.T) {
	for _, d := range convertIdTests {
		result := r.ConvertId(d.Id, d.From, d.To)
		if result != d.Expected {
			t.Errorf("ConvertId Failed: `-from %s -to %s %s` -> %s (expect: %s)", d.From, d.To, d.Id, result, d.Expected)
		}
	}
}

package registry

import (
	"reflect"
	"sort"
	"testing"
)

type resolveIdTest struct {
	Expected string
	Id       string
	From     string
	To       string
}

var resolveIdTests = []resolveIdTest{
	{"jotaro", "jotaro", "", ""},
	{"jotaro", "jotarok", "github", ""},
	{"jotarok", "jotaro", "", "github"},
	{"jotaro.kujo", "jotarok", "github", "facebook"},
	{"jotaro", "jotarok", "github", "google"},
}

func TestResolveId(t *testing.T) {
	var r, err = LoadConfig("../sample.json")
	if err != nil {
		t.Errorf("Load config error: [sample.json] %s", err)
	}

	for _, d := range resolveIdTests {
		result := r.ResolveId(d.Id, d.From, d.To)
		if result != d.Expected {
			t.Errorf("ResolveId Failed: `-from %s -to %s %s` -> %s (expect: %s)", d.From, d.To, d.Id, result, d.Expected)
		}
	}
}

func TestIdsAll(t *testing.T) {
	var r, err = LoadConfig("../sample.json")
	if err != nil {
		t.Errorf("Load config error: [sample.json] %s", err)
	}

	expected := []string{"joseph", "jotaro", "dio"}
	actual := r.Ids()

	sort.Strings(expected)
	sort.Strings(actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Ids Failed: expected=%v, actual=%v", expected, actual)
	}
}

func TestIdsForGroup(t *testing.T) {
	var r, err = LoadConfig("../sample.json")
	if err != nil {
		t.Errorf("Load config error: [sample.json] %s", err)
	}

	expected := []string{"dio", "joseph"}
	actual := r.Ids("part1", "part2")

	sort.Strings(expected)
	sort.Strings(actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Ids Failed: expected=%v, actual=%v", expected, actual)
	}
}

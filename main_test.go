package main

import "testing"

var fn = "read-files/log_2lines.log"
var result = []lineInfo {
	{"67.248.219.84", "/a", "9670"},
	{"133.141.46.13", "/a/c/d", "738"},
}

func TestProcessSeq(t *testing.T) {
	res := processSeq(fn)

	if len(res) != len(result) {
		t.Errorf("different result")
	}

	for i, should := range result {
		got := res[i]

		if *got == should {
			continue
		}

		t.Errorf("exp '%v' got '%v'",
			should, got)
	}
}

func TestProcessThreadPool(t *testing.T) {
	res := processThreadPool(fn, 2)

	if len(res) != len(result) {
		t.Errorf("different result")
	}

	for i, should := range result {
		got := res[i]

		if *got == should {
			continue
		}

		t.Errorf("exp '%v' got '%v'",
			should, got)
	}
}


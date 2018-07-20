package main

import "testing"

var fnBench = "read-files/log_200kb.log"
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
	res := processThreadPool(fn, 1)

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


func BenchmarkProcessSeq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSeq(fnBench)
	}
}

func BenchmarkProcessThreadPool1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processThreadPool(fnBench, 1)
	}
}

func BenchmarkProcessThreadPool2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processThreadPool(fnBench, 2)
	}
}

func BenchmarkProcessThreadPool3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processThreadPool(fnBench, 3)
	}
}

func BenchmarkProcessThreadPool5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processThreadPool(fnBench, 5)
	}
}


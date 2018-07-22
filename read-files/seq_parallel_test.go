package read_files

import (
	"strings"
	"testing"
)

var tableNLs = map[int64]string{
	0: "\nabcdef\n",
	1: "a\nbcd",
	3: "   \n",
}

var tableDivide = [][]string{
	{"aaaaaa", "bbbbbb", "cccccccc"},
	{"aa", "bb", "cc"},
}

func TestFindNextNewLine(t *testing.T) {
	for should, input := range tableNLs {
		r := strings.NewReader(input)
		got := findNextNewLine(r, 0)

		if got == should {
			continue
		}
		t.Errorf("for '%s' got '%d' should '%d'",
			input, got, should)
	}
}

func TestDivideToChunks(t *testing.T) {
	for _, test := range tableDivide {
		full := strings.Join(test, "\n")
		r := strings.NewReader(full)
		res := divideToChunks(r, int64(len(full)), len(test))

		if len(res) != len(test) {
			t.Errorf("failed for '%s'", test)
		}

		//var s int64
		//for i, subfile := range res {
		//	should := test[i]
		//}
		//WIP
	}
}

func TestParallelFileio(t *testing.T) {
	var lines []string
	f := func(b []byte) {
		lines = append(lines, string(b))
	}

	ParallelFileio(fnSmall, 64, f, 3)

	if len(lines) != 1000 {
		t.Errorf("lines exp %v got %v",
			1000, len(lines))
	}
}

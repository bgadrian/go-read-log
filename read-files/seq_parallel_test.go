package read_files

import (
	"strings"
	"sync"
	"testing"
)

var tableNLs = map[int64]string{
	0: "\nabcdef\n",
	1: "a\nbcd",
	2: "12\n",
	3: "1234", //EOF
}

var tableDivide = [][]string{
	{"aaaaaa", "bbbbbb", "cccccc"},
	{"a\naaaaa", "bbb\nbbb", "ccccc\nc"},
	{"aaaaaa", "bbbbbb", "cccccc", "ddddd"},
	{"aa", "b"},
}

type lineBucket struct {
	lines []string
	sync.Mutex
}

func TestFindNextNewLine(t *testing.T) {
	for should, input := range tableNLs {
		r := strings.NewReader(input)
		got, err := findNextNewLine(r, 0)

		if err != nil {
			t.Error(err)
		}

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
		res := splitToChunks(r, int64(len(full)), len(test))

		if len(res) != len(test) {
			t.Errorf("failed for '%s' got '%s'", test, res)
			continue
		}

		var s int64

		if res[0].start != 0 {
			t.Errorf("first chunk doesnt start at 0")
		}

		for i, subfile := range res {
			division := test[i]
			should := chunk{s, s + int64(len(division))}
			//remove the EOF
			if i == len(res)-1 {
				should.end--
			}
			s += subfile.Size()

			if i > 0 {
				prevSubFile := res[i-1]
				if subfile.start != prevSubFile.end+1 {
					t.Errorf("subfile overlapse")
				}
			}

			if subfile == should {
				continue
			}
			t.Errorf("chunk should '%v' got '%v'",
				should, subfile)
		}

		if res[len(res)-1].end != int64(len(full)-1) {
			t.Errorf("last chunk doesnt end at last byte")
		}
	}
}

func TestParallelFileio(t *testing.T) {
	bucket := lineBucket{}
	f := func(b []byte) {
		bucket.Lock()
		defer bucket.Unlock()
		bucket.lines = append(bucket.lines, string(b))
	}

	//ParallelFileio("log_2lines.log", 64, f, 1)
	ParallelFileio(fnSmall, 64, f, 3)

	if len(bucket.lines) != 1000 {
		t.Errorf("lines exp %v got %v",
			1000, len(bucket.lines))
	}
}

func TestParallelMmap(t *testing.T) {
	bucket := lineBucket{}
	f := func(b []byte) {
		bucket.Lock()
		defer bucket.Unlock()
		bucket.lines = append(bucket.lines, string(b))
	}

	ParallelMmap(fnSmall, 64, f, 3)

	if len(bucket.lines) != 1000 {
		t.Errorf("lines exp %v got %v",
			1000, len(bucket.lines))
	}
}

//
//func BenchmarkReadFileParallFileB64T1(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		ParallelFileio(fnMedium, 64, parseLine, 1)
//	}
//}
//
//func BenchmarkReadFileParallFileB128T1(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		ParallelFileio(fnMedium, 128, parseLine, 1)
//	}
//}

func BenchmarkReadFileParallFileB258T1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParallelFileio(fnMedium, 258, parseLine, 3)
	}
}

func BenchmarkReadFileParallMmapB258T1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParallelMmap(fnMedium, 258, parseLine, 3)
	}
}

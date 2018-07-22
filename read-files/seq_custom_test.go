package read_files

import "testing"

var parseLine = func(line []byte) {
	//fmt.Println(string(line))
}

func BenchmarkFileio128Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fileio(fnMedium, 128, parseLine)
	}
}

func BenchmarkMmapExp128Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mmapExp(fnMedium, 128, parseLine)
	}
}

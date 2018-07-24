package read_files

import "testing"

var parseLine = func(line []byte) {
	//fmt.Println(string(line))
}

func BenchmarkReadFileFileio256Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fileio(fnMedium, 256, parseLine)
	}
}

func BenchmarkReadFileMmapExp256Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mmapExp(fnMedium, 256, parseLine)
	}
}

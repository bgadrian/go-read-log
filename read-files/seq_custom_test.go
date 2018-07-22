package read_files

import "testing"

var parseLine = func(line []byte) {
	//fmt.Println(string(line))
}

//
//func TestSingle(t *testing.T) {
//	mmapExp(fnSmall, 0, parseLine)
//}

//func BenchmarkMmapSingle32Medium(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		mmapExp(fnMedium, 32, parseLine)
//	}
//}
//
//func BenchmarkMmapSingle64Medium(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		mmapExp(fnMedium, 64, parseLine)
//	}
//}

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

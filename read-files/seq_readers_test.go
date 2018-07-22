package read_files

import "testing"

var fnSmall = "log_200kb.log"
var fnMedium = "log_1MB.log"

//var fnBig = "../../log_big.log"

func BenchmarkBufioReaderLineSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerReadLine(fnSmall)
	}
}

func BenchmarkBufioReaderStringSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerString(fnSmall)
	}
}

func BenchmarkBufioScannerStringSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scannerString(fnSmall)
	}
}

func BenchmarkBufioScannerBytesSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scannerBytes(fnSmall)
	}
}

func BenchmarkBufioReaderLineMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerReadLine(fnMedium)
	}
}

func BenchmarkBufioReaderStringMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerString(fnMedium)
	}
}

func BenchmarkBufioScannerStringMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scannerString(fnMedium)
	}
}

func BenchmarkBufioScannerBytesMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scannerBytes(fnMedium)
	}
}

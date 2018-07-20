package read_files

import "testing"

var fn = "log_200kb.log"
//var fn = "log_1MB.log"

func BenchmarkBufioReaderLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerReadLine(fn)
	}
}

func BenchmarkBufioReaderString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerString(fn)
	}
}

func BenchmarkBufioScannerString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scannerString(fn)
	}
}

func BenchmarkBufioScannerBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scannerBytes(fn)
	}
}
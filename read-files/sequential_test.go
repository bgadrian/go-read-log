package read_files

import "testing"

//var fn = "apache_log_200kb.log"
var fn = "apache_log_1MB.log"

func BenchmarkBufioReaderBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerBytes(fn)
	}
}

func BenchmarkBufioReaderString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readerString(fn)
	}
}

func BenchmarkFileRead16KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read(fn, 16*1024)
	}
}

func BenchmarkFileRead32KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read(fn, 32*1024)
	}
}

func BenchmarkFileRead128KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read(fn, 128*1024)
	}
}

func BenchmarkFileRead512KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read(fn, 512*1024)
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
package read_files

import (
	"io"
	"log"

	"os"

	"golang.org/x/exp/mmap"
)

//simplified new like detection.
const nl = '\n'

type process func(line []byte)

func Fileio(fname string, bufferSize int64, f process) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	readerAtSplitToLines(file, bufferSize, f)
}

func mmapExp(fname string, bufferSize int64, f process) {
	readerAt, err := mmap.Open(fname)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer readerAt.Close()

	readerAtSplitToLines(readerAt, bufferSize, f)
}

func readerAtSplitToLines(r io.ReaderAt, bufferSize int64, f process) {
	var n int
	buffer := make([]byte, bufferSize)
	var offset int64
	var prefix []byte
	var err error
	for err == nil {
		n, err = r.ReadAt(buffer, offset)
		offset += int64(n)

		//for each byte that was read from the file
		for i := 0; i < n; i++ {
			b := buffer[i]
			if b == nl || err == io.EOF {
				//we detected a new line
				f(prefix)
				prefix = prefix[:0] //clear, keep capacity
				continue            //ignore the \n byte
			}
			prefix = append(prefix, b)
		}
	}

	if err != io.EOF {
		log.Fatal(err)
	}
}

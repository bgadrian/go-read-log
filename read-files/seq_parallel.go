package read_files

/**
	The main idea is that your process function
will be called with each line of the file,
but not in sequential order.
*/
import (
	"io"
	"log"
	"os"
	"sync"
)

type chunk struct {
	start, end int64
	//end byte should be a newLine or EOF
}

// return the first byte position that is newLine or EOF
func findNextNewLine(r io.ReaderAt, startAt int64) int64 {
	buffer := make([]byte, 32)
	offset := startAt
	var err error
	var n, i int
	for err == nil {
		n, err = r.ReadAt(buffer, offset)

		for i = 0; i < n; i++ {
			if buffer[i] == nl {
				return offset + int64(i)
			}
		}
		offset += int64(n)
	}

	if err == io.EOF {
		return offset
	}
	return -1
}

func divideToChunks(r io.ReaderAt, totalSize int64, count int) []chunk {
	var chunks []chunk
	aproxChunkSize := totalSize / int64(count)
	for i := 0; i < count; i++ {
		var startByte int64
		if i > 0 {
			//start is the previous end
			startByte = chunks[i-1].end + 1
		}
		aproxEndByte := startByte + aproxChunkSize
		endByte := findNextNewLine(r, aproxEndByte)
		chunks = append(chunks, chunk{startByte, endByte})
		if chunks[i].end == totalSize {
			break
		} //done earlier
	}
	return chunks
}

func ParallelFileio(fname string, bufferSize int64, f process, tasksCount int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//difficult thing is to split the file at exactly new line bytes
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	chunks := divideToChunks(file, stat.Size(), tasksCount)

	//we are using SectionReaders to easy the logic and reuse the code
	//each chunk will think is a different file
	var wg sync.WaitGroup
	wg.Add(len(chunks))

	for _, subfile := range chunks {
		size := subfile.end - subfile.start
		r := io.NewSectionReader(file, subfile.start, size)

		go func(r io.ReaderAt) {
			defer wg.Done()
			readerAtSplitToLines(r, bufferSize, f)
		}(r)
	}

	wg.Wait()

}

//func mmapExp(fname string, bufferSize int64, f process) {
//	readerAt, err := mmap.Open(fname)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	defer readerAt.Close()
//
//	readerAtSplitToLines(readerAt, bufferSize, f)
//}

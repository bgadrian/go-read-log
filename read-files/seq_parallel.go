package read_files

/**
	The main idea is that your process function
will be called with each line of the file,
but not in sequential order.
*/
import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/exp/mmap"
)

//chunk Slice of a long byte slice.
// Start/end are byte positions
// Start/End are both inclusive.
type chunk struct {
	start, end int64
	//file[end] can be a /n, if not probably its the EOF
}

func (c chunk) Size() int64 {
	return c.end - c.start + 1
}

func (c chunk) String() string {
	return fmt.Sprintf("{%d:%d(%d)}", c.start, c.end, c.Size())
}

// return the first byte position that is newLine or
// the last byte position before EOF
func findNextNewLine(r io.ReaderAt, startAt int64) (int64, error) {
	buffer := make([]byte, 32)
	offset := startAt
	var err error
	var n, i int
	for err == nil {
		n, err = r.ReadAt(buffer, offset)

		for i = 0; i < n; i++ {
			if buffer[i] == nl {
				return offset + int64(i), nil
			}
		}
		offset += int64(n)
	}

	if err == io.EOF {
		return offset - 1, nil
	}
	return -1, errors.New("no new line or EOF")
}

func splitToChunks(r io.ReaderAt, totalSize int64, count int) []chunk {
	var result []chunk
	aproxChunkSize := totalSize / int64(count)
	leftBytes := totalSize
	for i := 0; i < count; i++ {
		var startByte int64
		if i > 0 {
			//start is the previous end
			startByte = result[i-1].end + 1
		}
		aproxEndByte := startByte + aproxChunkSize
		if aproxEndByte >= totalSize {
			aproxEndByte = totalSize - 1
		}
		endByte, err := findNextNewLine(r, aproxEndByte)
		if err != nil {
			log.Panic(err)
		}

		result = append(result, chunk{startByte, endByte})

		leftBytes -= result[i].Size()
		if leftBytes == 0 {
			break //done earlier
		}
		if leftBytes < 0 {
			log.Panicf("more bytes in result than total %s  %d", result, totalSize)
		}
	}
	if leftBytes > 0 {
		//we have leftovers, we put them to the end
		result[len(result)-1].end = totalSize - 1
	}
	return result
}

func ParallelFileio(fname string, bufferSize int64, f process, tasksCount int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	readerParall(file, stat.Size(), bufferSize, f, tasksCount)
}

func ParallelMmap(fname string, bufferSize int64, f process, tasksCount int) {
	readerAt, err := mmap.Open(fname)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer readerAt.Close()
	readerParall(readerAt, int64(readerAt.Len()), bufferSize, f, tasksCount)
}

func readerParall(r io.ReaderAt, fileSize int64, bufferSize int64, f process, tasksCount int) {
	//difficult thing is to split the file at exactly new line bytes

	chunks := splitToChunks(r, fileSize, tasksCount)

	//we are using SectionReaders to easy the logic and reuse the code
	//each chunk will think is a different file
	var wg sync.WaitGroup
	wg.Add(len(chunks))

	for _, subfile := range chunks {
		r := io.NewSectionReader(r, subfile.start, subfile.Size())

		go func(r io.ReaderAt) {
			defer wg.Done()
			readerAtSplitToLines(r, bufferSize, f)
		}(r)
	}

	wg.Wait()
}

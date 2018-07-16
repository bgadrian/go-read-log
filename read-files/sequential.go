package read_files

import (
	"os"
	"log"
	"io"
	"bufio"
)

// readerBytes bufio.reader.ReadBytes each line.
func readerBytes(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)

	var line []byte
	fake := func(s []byte) {

	}
	for  ; err != nil; line, err = bufferedReader.ReadBytes('\n'){
		fake(line)
	}
}

// readerString bufio.reader.ReadString each line.
// converts to a string so should be slower  than readBytes
func readerString(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)
	fake := func(s string) {

	}
	var line string

	for  ; err != nil; line, err = bufferedReader.ReadString('\n'){
		fake(line)
	}
}

// read It doesn't read by line, but by buffer, so we have to
// implement that logic
func read(fname string, bufferSize int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	buf := make([]byte, bufferSize) // define your buffer size here.

	for {
		n, err := file.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}
}

// scanner bufio.scanner.Text (string each line)
func scannerString(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {             // internally, it advances token based on sperator
		scanner.Text() // token in unicode-char
	}
}

// scanner bufio.scanner.Bytes (each line)
func scannerBytes(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {             // internally, it advances token based on sperator
		scanner.Bytes()
	}
}



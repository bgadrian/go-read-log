package read_files

import (
	"os"
	"log"
	"io"
	"bufio"
)

func readerReadLine(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedReader := bufio.NewReader(file)

	var line []byte
	var isPrefix bool
	fake := func(s []byte) {

	}
	for {
		line, isPrefix, err = bufferedReader.ReadLine()
		if isPrefix {
			log.Fatal("isPrefix logic to concatenate lines is not implemented")
		}
		if err != nil {
			break
		}
		fake(line)
	}

	if err != io.EOF {
		log.Fatal(err)
	}
}

func readerString(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedReader := bufio.NewReader(file)
	fake := func(s string) {

	}
	var line string

	for {
		line, err = bufferedReader.ReadString('\n')
		if  err != nil {
			break
		}
		fake(line)
	}

	if err != io.EOF {
		log.Fatal(err)
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



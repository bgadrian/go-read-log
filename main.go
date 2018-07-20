package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"sync"
)

type lineInfo struct {
	a,b,c string
}

// processSeq 1 thread, open file and parse each line
func processSeq(fname string)[]*lineInfo {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lineResult []string
	var allData []*lineInfo

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//0, 6 ,8
		lineResult = strings.SplitN(scanner.Text(), " ", 8+2)
		if len(lineResult) < 10 {
			continue
		}
		allData = append(allData, &lineInfo{lineResult[0], lineResult[6], lineResult[8]})

	}

	return allData
}

// processThreadPool split the work between jobs.
func processThreadPool(fname string, workers int) []*lineInfo {
	jobs := make(chan string, workers)
	splitData := make([][]*lineInfo, workers)
	var wg sync.WaitGroup

	worker := func(work chan string, workerID int) {
		defer wg.Done()
		var lineResult []string
		for  line := range jobs {
			lineResult = strings.SplitN(line, " ", 8+2)
			if len(lineResult) < 10 {
				continue
			}
			splitData[workerID] = append(splitData[workerID], &lineInfo{lineResult[0], lineResult[6], lineResult[8]} )
		}
	}

	wg.Add(workers)
	for i := 0 ; i < workers ; i++ {
		go func(id int) {
			worker(jobs, id)
		}(i)
	}

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//cannot use Bytes(), race condition
		jobs <- scanner.Text()
	}
	//wait for all workers finish
	close(jobs)
	wg.Wait()

	//collect their work to one result
	var allData []*lineInfo
	for _, data := range splitData {
		allData = append(allData, data...)
	}

	return allData
}

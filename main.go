package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/roatt-tilma/CyberSec-subDomainEnumeration/brutus"
	"github.com/roatt-tilma/CyberSec-subDomainEnumeration/logger"
	"github.com/roatt-tilma/CyberSec-subDomainEnumeration/progress"
)

func main() {
	flag.Parse()

	if len(urls) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if wordlist == "" {
		flag.Usage()
		os.Exit(1)
	}

	if enumerationtype > 1 || enumerationtype < 0 {
		enumerationtype = 0
	}

	if cores < 1 || cores > 8 {
		cores = 1
	}

	if len(success) == 0 {
		success = append(success, "200")
	}

	now := time.Now()
	defer func() {
		duration := time.Since(now)
		logger.Info("Execution time: " + duration.String())
	}()

	workers := cores

	if enumerationtype == 0 {
		logger.SubDomain()
	} else {
		logger.Directory()
	}

	if workers == 1 {
		logger.Info("Starting " + strconv.Itoa(workers) + " worker thread...")
	} else {
		logger.Info("Starting " + strconv.Itoa(workers) + " worker threads...")
	}

	successCodes := make(map[string]bool)
	for _, code := range success {
		successCodes[code] = true
	}

	var wg sync.WaitGroup
	var bar progress.Progress
	c := make(chan *brutus.Brute)
	logs := make(chan logger.Log)

	go logger.Start(logs)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for b := range c {
				b.Try(successCodes, logs)
			}
		}()
	}

	file, err := os.Open(wordlist)
	if err != nil {
		logger.Fatal("Could not open wordlist: " + wordlist)
	}

	defer file.Close()

	stat, _ := file.Stat()
	size := stat.Size()

	bar.New(0, int(size))

	count := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		for _, url := range urls {
			c <- brutus.New(url, text, enumerationtype)
		}

		count += len(text) + 1
		bar.Play(count, logs)
	}

	close(c)

	wg.Wait()

	// logger.Info("\nEnumeration: " + strconv.Itoa(count) + "\nTotal: " + strconv.FormatInt(size-1, 10))
}

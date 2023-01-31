package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"sync"

	"babushka/helper"
)

// get english words from site https://www.ef.com/wwen/english-resources/english-vocabulary/top-3000-words/

type Word struct {
	Score float64 // precision is not important
	Word  string
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// file size
	fileInfo, err := file.Stat()
	if err != nil {
		println(err)
		return
	}

	sizeOfFile := fileInfo.Size()
	proc := runtime.GOMAXPROCS(0) - 1
	kbPerGoroutine := int(math.Ceil(float64(sizeOfFile) / float64(proc) / 1024))

	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, kbPerGoroutine*1024)
		return lines
	}}

	newWordch := make(chan Word, 100)

	var wg2 sync.WaitGroup

	wg2.Add(1)

	topWords := make([]Word, 0, 10)

	go func() {
		for word := range newWordch {
			if len(topWords) < cap(topWords) {
				topWords = append(topWords, word)
				continue
			}

			for i := 0; i < len(topWords); i++ {
				if word.Score < topWords[i].Score {
					topWords[i] = word
					break
				}
			}
		}

		wg2.Done()
		println("channel closed")
	}()

	var wg sync.WaitGroup

	reader := bufio.NewReader(file)

loop:
	for i := 0; i < proc; i++ {
		buf, ok := linesPool.Get().([]byte)
		if !ok {
			println(buf)
			panic("linesPool.Get() returned not []byte")
		}

		bytesNums, err := reader.Read(buf)

		switch {
		case err == io.EOF && bytesNums == 0:
			break loop
		case err != nil:
			fmt.Println(err)
			break loop
		}

		buf = buf[:bytesNums]
		// if was readed half of the word.
		endOfWord, err := reader.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, endOfWord...)
		}

		wg.Add(1)

		go func() {
			readWordsFromChunk(&buf, &linesPool, newWordch)
			wg.Done()
		}()
	}

	wg.Wait()
	close(newWordch)
	wg2.Wait()
	println("all goroutines done")

	for _, word := range topWords {
		println(word.Word, word.Score)
	}
}

func readWordsFromChunk(buffer *[]byte, chankPool *sync.Pool, newWordch chan<- Word) {
	scanner := bufio.NewScanner(bytes.NewReader(*buffer))
	chankPool.Put(buffer)

	for scanner.Scan() {
		word := scanner.Text()
		score := getScore(word)
		newWordch <- Word{Score: score, Word: word}
	}
}

func getScore(word string) float64 {
	var distance uint8

	for i := 0; i < len(word)-1; i++ {
		distance += helper.GetDistance(word[i], word[i+1])
	}

	return float64(distance) / float64(len(word))
}

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
)

// get english words from site https://www.ef.com/wwen/english-resources/english-vocabulary/top-3000-words/

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

	newWordch := make(chan string, 100)

	var wg2 sync.WaitGroup

	wg2.Add(1)

	go func() {
		i := 0
		for word := range newWordch {
			i++
			println(i, word)
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
}

func readWordsFromChunk(buffer *[]byte, chankPool *sync.Pool, newWordch chan<- string) {
	scanner := bufio.NewScanner(bytes.NewReader(*buffer))
	chankPool.Put(buffer)

	for scanner.Scan() {
		word := scanner.Text()
		newWordch <- word
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"

	"babushka/helper"
)

// get english words from site https://www.ef.com/wwen/english-resources/english-vocabulary/top-3000-words/

type Word struct {
	Score float64 // precision is not important
	Word  string
}

const TOP = 20

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

	// TODO: compare with a x2 goroutines. (while gorutine make syscall for read, another goroutine can read from buffer)?
	sizeOfFile := fileInfo.Size()
	proc := runtime.GOMAXPROCS(0) - 1
	kbPerGoroutine := int(math.Ceil(float64(sizeOfFile) / float64(proc) / 1024))

	// TODO: need more Pool for word?
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, kbPerGoroutine*1024)
		return lines
	}}

	newWordch := make(chan Word, 100)

	var wg2 sync.WaitGroup

	wg2.Add(1)

	topWords := make([]Word, 0, TOP)

	// TODO: benchmark, compare with >1 read goroutines || 1 read 1 write goroutine
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

	// for _, word := range topWords {
	// 	println(word.Word, word.Score)
	// }

	println("\n", "Babushka here's your password! ", chooseWordsForPassword(topWords))
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

type Top struct {
	WordFirst  int
	WordSecond int
	Distance   uint8
}

func chooseWordsForPassword(topWords []Word) (response string) {
	// create slice of all possible combinations
	// and sort it by distance
	top := make([]Top, 0, TOP*TOP)

	// TODO: how to optimize...?
	for i := 0; i < len(topWords); i++ {
		for j := 0; j < len(topWords); j++ {
			if i == j {
				continue // skip same words? or not?
			}

			firstChar := topWords[j].Word[0]
			lastChar := topWords[i].Word[len(topWords[i].Word)-1]
			dist := helper.GetDistance(lastChar, firstChar)

			top = append(top, Top{i, j, dist})
		}
	}

	// it won't get any worse =)
	sort.Slice(top, func(i, j int) bool {
		return top[i].Distance < top[j].Distance
	})

	// Ужас... Будто на питоне пишу
	sumLen := 0
	index := 0
	wasWord := make(map[int]bool)
	// Choose first two word from top, because they have the smallest distance.
	sumLen += len(topWords[top[index].WordFirst].Word) + len(topWords[top[index].WordSecond].Word)
	response += topWords[top[index].WordFirst].Word + topWords[top[index].WordSecond].Word

	for {
		wasWord[top[index].WordFirst] = true
		wasWord[top[index].WordSecond] = true
		index = getIndexNexWord(&top, top[index].WordSecond, wasWord)

		if sumLen > 20 && sumLen < 30 {
			break
		}

		sumLen += len(topWords[top[index].WordSecond].Word)
		response += topWords[top[index].WordSecond].Word
	}

	return response
}

func getIndexNexWord(top *[]Top, wordSecond int, wasWord map[int]bool) int {
	indexMinDist := 0

	for i := 0; i < len(*top); i++ {
		// between WordFirst and wordSecond should be min distance.
		if (*top)[i].WordFirst == wordSecond && !wasWord[(*top)[i].WordSecond] {
			return i // we alreaady sorded slice by distance
		}
	}

	return indexMinDist
}

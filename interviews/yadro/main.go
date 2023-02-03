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
	"sync/atomic"

	"babushka/helper"
)

// get english words from site https://www.ef.com/wwen/english-resources/english-vocabulary/top-3000-words/

type Word struct {
	Dist uint32
	Word string
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

	Words := make([]Word, 0, 5000)

	// TODO: benchmark, compare with >1 read goroutines || 1 read 1 write goroutine
	go func() {
		for word := range newWordch {
			Words = append(Words, word)
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

	println("\n", "Babushka here's your password! ", chooseWordsForPassword(Words))
}

func readWordsFromChunk(buffer *[]byte, chankPool *sync.Pool, newWordch chan<- Word) {
	scanner := bufio.NewScanner(bytes.NewReader(*buffer))
	chankPool.Put(buffer)

	for scanner.Scan() {
		word := scanner.Text()
		score := getScore(word)
		newWordch <- Word{Dist: score, Word: word}
	}
}

func getScore(word string) uint32 {
	var distance uint32

	for i := 0; i < len(word)-1; i++ {
		distance += uint32(helper.GetDistance(word[i], word[i+1]))
	}

	return distance
}

func chooseWordsForPassword(words []Word) (response string) {

	// for calculated word
	var (
		score atomic.Uint32
		mI    sync.Map
		mJ    sync.Map
		mK    sync.Map
		wg    sync.WaitGroup
	)

	score.Store(65535)

	from := 0
	delta := len(words) / runtime.GOMAXPROCS(0)
	var pass atomic.Value

	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)

		go func(from2 int) {

			for i := from2; i < from2+delta; i++ {
				var saveWordI Word

				lastLetterI := words[i].Word[len(words[i].Word)-1]
				if v, ok := mI.Load(string(lastLetterI) + fmt.Sprint(len(words[i].Word))); ok {
					dist := getDist([]Word{v.(Word), words[i]})
					if score.Load() > dist {
						println(dist, words[i].Word, 4)
						pass.Store(words[i].Word + v.(Word).Word)
						score.Store(dist)
					}

					println("mI hit")
					continue
				}

				for j := 0; j < len(words); j++ {
					var saveWordJ Word

					lastLetterJ := words[j].Word[len(words[j].Word)-1]
					if v, ok := mJ.Load(string(lastLetterJ) + fmt.Sprint(len(words[i].Word)+len(words[j].Word))); ok {
						dist := getDist([]Word{v.(Word), words[i], words[j]})
						if score.Load() > dist {
							println(dist, words[i].Word, words[j].Word, 3)
							pass.Store(words[i].Word + words[j].Word + v.(Word).Word)
							score.Store(dist)
						}

						// println("mJ hit")
						continue
					}

					for k := 0; k < len(words); k++ {
						var saveWordK Word

						lastLetterK := words[k].Word[len(words[k].Word)-1]
						if v, ok := mK.Load(string(lastLetterK) + fmt.Sprint(len(words[i].Word)+len(words[j].Word)+len(words[k].Word))); ok {
							dist := getDist([]Word{v.(Word), words[i], words[j], words[k]})
							if score.Load() > dist {
								println(dist, words[i].Word, words[j].Word, words[k].Word, 2)
								pass.Store(words[i].Word + words[j].Word + words[k].Word + v.(Word).Word)
								score.Store(dist)
							}

							// println("mK hit")
							continue
						}

						for l := 0; l < len(words); l++ {
							combination := []Word{words[i], words[j], words[k], words[l]}
							if len(combination[0].Word)+len(combination[1].Word)+len(combination[2].Word)+len(combination[3].Word) >= 20 &&
								len(combination[0].Word)+len(combination[1].Word)+len(combination[2].Word)+len(combination[3].Word) <= 24 {

								dist := getDist(combination)
								if score.Load() > dist {
									saveWordI.Word = combination[1].Word + combination[2].Word + combination[3].Word
									saveWordI.Dist = getDist(combination[1:])

									saveWordJ.Word = combination[2].Word + combination[3].Word
									saveWordJ.Dist = getDist(combination[2:])

									saveWordK.Word = combination[3].Word
									saveWordK.Dist = getDist(combination[3:])

									score.Store(dist)
									pass.Store(combination[0].Word + combination[1].Word + combination[2].Word + combination[3].Word)
									println(score.Load(), combination[0].Word, combination[1].Word, combination[2].Word, combination[3].Word)
								}
							}
						}

						if saveWordK.Word != "" {
							mK.Store(string(lastLetterK)+fmt.Sprint(len(words[i].Word)+len(words[j].Word)+len(words[k].Word)), saveWordK)
						}
					}

					if saveWordJ.Word != "" {
						mJ.Store(string(lastLetterJ)+fmt.Sprint(len(words[i].Word)+len(words[j].Word)), saveWordJ)
					}
				}

				if saveWordI.Word != "" {
					mI.Store(string(lastLetterI)+fmt.Sprint(len(words[i].Word)), saveWordI)
				}
			}

			wg.Done()
		}(from)

		from += delta
	}

	wg.Wait()

	println(score.Load())

	return pass.Load().(string)
}

func getDist(combination []Word) (resp uint32) {
	for i := 0; i < len(combination); i++ {
		resp += combination[i].Dist
	}

	for i := 0; i < len(combination)-1; i++ {
		resp += helper.GetDistance(combination[i].Word[len(combination[i].Word)-1], combination[i+1].Word[0])
	}

	return
}

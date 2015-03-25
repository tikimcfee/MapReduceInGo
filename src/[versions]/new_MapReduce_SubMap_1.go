package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func findAndSplitWork(dirname string, keyMap map[string]chan int) {
	// channel that mappers call out to when they are finished parsing out their text
	doneChannel := make(chan bool)

	dir, _ := os.Open(dirname)
	dirnames, _ := dir.Readdirnames(-1)
	amDone := 0
	for i := 0; i < len(dirnames); i++ {
		fullpath := dirname + "/" + dirnames[i]
		file, _ := os.Stat(fullpath)
		if file.IsDir() {
		} else {
			// create a mapper for the file, and let it handle it
			thisFile, err := os.Open(fullpath)
			if err != nil {
				return
			}
			go mapper(thisFile, keyMap, doneChannel)
			amDone++
		}
	}

	// the main go routine should not move forward and close the reducing channels until every mapper has completed.
	// this is another point of failure; if we somehow have a set of buffered channels to increase channel speed, what
	// does closing a buffered channel do?
	for amDone != 0 {
		<-doneChannel
		amDone--
	}
}

func mapper(input *os.File, keyToReducerMap map[string]chan int, done chan bool) {
	// create a 'routine local' mapping of the keys we're interested in looking for
	localKeyMapChannel := make(chan int)
	stringToPositionMapping := make(map[string]int)
	mappedTotals := make([]int, len(keyToReducerMap))

	i := 0
	for key, _ := range keyToReducerMap {
		stringToPositionMapping[key] = i
		i++
	}

	// create a
	reader := bufio.NewReader(input)
	r, _ := regexp.Compile("[!-@]*")
	submapperFinished := make(chan bool)
	allParsed := make(chan bool)
	subMapperCount := 0
	go func() {
		for {
			line, err := reader.ReadString('\n')
			parsedWords := strings.TrimSpace(line)
			parsedWords = r.ReplaceAllString(parsedWords, "")
			// parsedWordsArray := strings.Split(parsedWords, " ")

			// for _, word := range parsedWordsArray {
			// 	// word = strings.TrimSpace(word)
			// 	if word == "" {
			// 		continue
			// 	}

			// 	if _, exists := keyToReducerMap[word]; exists {
			// 		keys[word]++
			// 	}
			// }

			if err == io.EOF {
				go subMapper(parsedWords, stringToPositionMapping, localKeyMapChannel, submapperFinished)
				subMapperCount++
				break
			} else {
				go subMapper(parsedWords, stringToPositionMapping, localKeyMapChannel, submapperFinished)
				subMapperCount++
			}
		}
		allParsed <- true
	}()

	go func() {
		for mapped := range localKeyMapChannel {
			// fmt.Println("Mapped", mapped)
			mappedTotals[mapped]++
		}
	}()

	<-allParsed
	go func() {
		for subMapperCount != 0 {
			<-submapperFinished
			subMapperCount--
		}
		allParsed <- true
	}()
	<-allParsed

	close(localKeyMapChannel)
	close(submapperFinished)

	for key, value := range stringToPositionMapping {
		keyToReducerMap[key] <- mappedTotals[value]
	}

	input.Close()
	done <- true
}

func subMapper(input string, keyMap map[string]int, newMappedValueChannel chan int, submapperFinished chan bool) {
	// parsedWords := strings.TrimSpace(input)
	// parsedWords = r.ReplaceAllString(parsedWords, "")
	parsedWordsArray := strings.Split(input, " ")

	for _, word := range parsedWordsArray {
		// word = strings.TrimSpace(word)
		if word == "" {
			continue
		}

		if _, exists := keyMap[word]; exists {
			newMappedValueChannel <- keyMap[word]
		}
	}

	submapperFinished <- true
}

func reducer(listeningWord string, mappedStringCount chan int, outputChannel chan map[string]int) {
	total := 0

	// range continuously over the channels we've created. the main routine will close the channel
	// only *after* all mappers have reported in a complete parsing.
	for countedValue := range mappedStringCount {
		total += countedValue
	}

	// the channel has been closed (range complete), so send off results on our output channel
	outputChannel <- map[string]int{listeningWord: total}
}

func main() {
	// includine runtime package and set MAX_PROCs to whatever machine we have at hand to paralleliz all fast and stuff
	// runtime.setMaxProcs(X)

	runtime.GOMAXPROCS(8)
	timeStart := time.Now()

	// get the file to find what keys the user wants to parse; the default name is 'key_file'
	thisFile, err := os.Open("key_file")
	if err != nil {
		return
	}

	// final channel that reducers will send their totals to
	finalOutputChannel := make(chan map[string]int)

	/** include keyword finding / parsing for later channels to be made
	* 	command line more demoable?
	**/

	// create the empty keyMap of channels
	keyMap := make(map[string]chan int)

	// create the channels from the above user input that the mappers will send on
	reader := bufio.NewReader(thisFile)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// a non-empty word requires a new reducer to listen for output on.
		// spawn a goroutine for that reducer listener, add it to the map, and continue
		wordChannel := make(chan int, 10000)
		keyMap[line] = wordChannel
		go reducer(line, wordChannel, finalOutputChannel)

		if err == io.EOF {
			break
		}
	}

	// fire off the mappers in a go routine ; this will change frequently, as what mappers do, and
	// what they are given initially, will change. the directory we search for file in is 'input'
	findAndSplitWork("input", keyMap)

	// inform the for := range's of the reducers that all input has been parsed, and that they may shut down. Also, accept
	// from them their final word counts
	for _, channel := range keyMap {
		close(channel)
		fmt.Println(<-finalOutputChannel)
		// <-finalOutputChannel
	}

	timeEnd := time.Since(timeStart)

	fmt.Printf("Execution took %s\n", timeEnd)
}

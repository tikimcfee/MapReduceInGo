package main

import (
	"fmt"
	"os"
	"strings"
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
			// _find_files(fullpath, output)
		} else {
			// create a mapper for the file, and let it handle it
			thisFile, err := os.Open(filename)
			if err != nil {
				return
			}
			mapper(thisFile, keyMap, doneChannel)
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

func mapper(input File, keyToReducerMap map[string]chan int, done chan bool) {
	keys := make(map[string]int)

	reader := bufio.NewReader(thisFile)
	for {
		line, err := reader.ReadString('\n')
		keys := strings.Split(line, " ")

		for _, word := range keys {
			keys[word]++
		}

		if err == io.EOF {
			break
		}
	}

	for key, value := range keys {
		// fmt.Printf("%s : %d\n", key, value)
		// for each key in keyToReduceMap
		// get the value found in keys above, and send value to channel for corresponding key
		// fmt.Printf("sending %s to %s\n...", value, keyToReducerMap[key])
		keyToReducerMap[key] <- value
		// fmt.Printf("...sent %s to %s\n", value, keyToReducerMap[key])
	}
	done <- true
}

func reducer(listeningWord string, mappedStringCount chan int, outputChannel chan map[string]int) {
	total := 0

	for countedValue := range mappedStringCount {
		total += countedValue
	}

	outputChannel <- map[string]int{listeningWord: total}
}

func main() {
	// includine runtime package and set MAX_PROCs to whatever machine we have at hand to paralleliz all fast and stuff
	// runtime.setMaxProcs(X)
	thisFile, err := os.Open("key_file")

	reader := bufio.NewReader(thisFile)
	for {
		line, err := reader.ReadString('\n')
		keys := strings.Split(line, " ")

		for _, word := range keys {
			keys[word]++
		}

		if err == io.EOF {
			break
		}
	}

	// include keyword finding / parsing for later channels to be made
	// command line more demoable?
	keyMap := make(map[string]chan int)
	keys := strings.Split(in, " ")
	keys2 := strings.Split(in2, " ")

	// final channel that reducers will send their totals to
	finalOutputChannel := make(chan map[string]int)

	// create the channels from the above user input that the mappers will send on
	for _, word := range keys {
		wordChannel := make(chan int)
		keyMap[word] = wordChannel
		go reducer(word, wordChannel, finalOutputChannel)
	}

	// ditto
	for _, word := range keys2 {
		if _, exists := keyMap[word]; !exists {
			wordChannel := make(chan int)
			keyMap[word] = wordChannel
			go reducer(word, wordChannel, finalOutputChannel)
		}
	}

	// fire off the mappers in a go routine ; this will change frequently, as what mappers do, and
	// what they are given initially, will change
	go mapper(keys, keyMap, doneChannel)
	go mapper(keys2, keyMap, doneChannel)

	// the main go routine should not move forward and close the reducing channels until every mapper has completed.
	// this is another point of failure; if we somehow have a set of buffered channels to increase channel speed, what
	// does closing a buffered channel do?
	amDone := 2
	for amDone != 0 {
		<-doneChannel
		amDone--
	}

	// inform the for := range's of the reducers that all input has been parsed, and that they may shut down. Also, except
	// from them their final word counts
	for _, channel := range keyMap {
		close(channel)
		fmt.Println(<-finalOutputChannel)
	}
}

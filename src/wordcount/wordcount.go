package main

import (
	"fileiter"
	"flag"
	"fmt"
	"mapreduce"
	"os"
	"regexp"
	"runtime"
	"time"
)

var input_directory = flag.String("input_directory", "input", "directory of all input files for parsing")

func find_files(dirname string) chan interface{} {
	output := make(chan interface{})
	go func() {
		_find_files(dirname, output)
		close(output)
	}()
	return output
}

func _find_files(dirname string, output chan interface{}) {
	dir, _ := os.Open(dirname)
	dirnames, _ := dir.Readdirnames(-1)
	for i := 0; i < len(dirnames); i++ {
		fullpath := dirname + "/" + dirnames[i]
		file, _ := os.Stat(fullpath)
		if file.IsDir() {
			_find_files(fullpath, output)
		} else {
			output <- fullpath
		}
	}
}

func wordcount(filename interface{}, output chan interface{}) {
	results := map[string]int{}
	wordsRE := regexp.MustCompile(`[A-Za-z0-9_]*`)
	for line := range fileiter.EachLine(filename.(string)) {
		for _, match := range wordsRE.FindAllString(line, -1) {
			results[match]++
		}
	}
	output <- results
}

func reducer(input chan interface{}, output chan interface{}) {
	results := map[string]int{}
	for new_matches := range input {
		for key, value := range new_matches.(map[string]int) {
			previous_count, exists := results[key]
			if !exists {
				results[key] = value
			} else {
				results[key] = previous_count + value
			}
		}
	}
	output <- results
}

func main() {
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		fmt.Println("bad input directory; specify an existing folder with at least 1 text file")
		os.Exit(1)
	}

	runtime.GOMAXPROCS(1)
	timeStart := time.Now()
	// fmt.Println(mapreduce.MapReduce(wordcount, reducer, find_files("input"), 20))
	mapreduce.MapReduce(wordcount, reducer, find_files(dir), 20)
	timeEnd := time.Since(timeStart)
	fmt.Printf("Execution took %s\n", timeEnd)
}

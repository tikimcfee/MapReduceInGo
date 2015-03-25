package main

import (
	"fmt"
	"strings"
)

func mapper(input []string, keyToReducerMap map[string]chan int, done chan bool) {

	keys := make(map[string]int)
	for _, word := range input {
		keys[word]++
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
	in := "Bacon ipsum dolor amet turkey kielbasa ham hock, doner t-bone salami corned beef landjaeger cow beef ribs sirloin. Ball tip t-bone bresaola, shankle tenderloin biltong landjaeger porchetta. Ground round cupim alcatra, jerky porchetta swine pastrami. Pastrami tenderloin ball tip ham hamburger beef ribs meatloaf pork loin brisket venison jowl leberkas tail pork belly biltong. Doner ribeye frankfurter venison pastrami picanha. Kevin turducken meatball, shankle alcatra hamburger kielbasa landjaeger pig pork chop chuck boudin. Bacon ipsum dolor amet turkey kielbasa ham hock, doner t-bone salami corned beef landjaeger cow beef ribs sirloin. Ball tip t-bone bresaola, shankle tenderloin biltong landjaeger porchetta. Ground round cupim alcatra, jerky porchetta swine pastrami. Pastrami tenderloin ball tip ham hamburger beef ribs meatloaf pork loin brisket venison jowl leberkas tail pork belly biltong. Doner ribeye frankfurter venison pastrami picanha. Kevin turducken meatball, shankle alcatra hamburger kielbasa landjaeger pig pork chop chuck boudin. Bacon ipsum dolor amet turkey kielbasa ham hock, doner t-bone salami corned beef landjaeger cow beef ribs sirloin. Ball tip t-bone bresaola, shankle tenderloin biltong landjaeger porchetta. Ground round cupim alcatra, jerky porchetta swine pastrami. Pastrami tenderloin ball tip ham hamburger beef ribs meatloaf pork loin brisket venison jowl leberkas tail pork belly biltong. Doner ribeye frankfurter venison pastrami picanha. Kevin turducken meatball, shankle alcatra hamburger kielbasa landjaeger pig pork chop chuck boudin. Bacon ipsum dolor amet turkey kielbasa ham hock, doner t-bone salami corned beef landjaeger cow beef ribs sirloin. Ball tip t-bone bresaola, shankle tenderloin biltong landjaeger porchetta. Ground round cupim alcatra, jerky porchetta swine pastrami. Pastrami tenderloin ball tip ham hamburger beef ribs meatloaf pork loin brisket venison jowl leberkas tail pork belly biltong. Doner ribeye frankfurter venison pastrami picanha. Kevin turducken meatball, shankle alcatra hamburger kielbasa landjaeger pig pork chop chuck boudin. "
	in2 := "Bacon ipsum dolor amet ground round rump short ribs bresaola kevin beef, pastrami chuck hamburger alcatra prosciutto. Ground round ham hock porchetta, cupim spare ribs doner ball tip bresaola tail chicken meatball. Salami meatball ground round boudin, chicken pork chop pastrami ham hock pork loin meatloaf. Rump sirloin corned beef boudin. Pancetta short loin boudin prosciutto chicken turducken ball tip corned beef t-bone sirloin tenderloin venison. Pork belly hamburger pig biltong, frankfurter capicola doner. Bresaola chicken sausage, bacon brisket swine boudin sirloin capicola flank kevin turducken. Bacon ipsum dolor amet ground round rump short ribs bresaola kevin beef, pastrami chuck hamburger alcatra prosciutto. Ground round ham hock porchetta, cupim spare ribs doner ball tip bresaola tail chicken meatball. Salami meatball ground round boudin, chicken pork chop pastrami ham hock pork loin meatloaf. Rump sirloin corned beef boudin. Pancetta short loin boudin prosciutto chicken turducken ball tip corned beef t-bone sirloin tenderloin venison. Pork belly hamburger pig biltong, frankfurter capicola doner. Bresaola chicken sausage, bacon brisket swine boudin sirloin capicola flank kevin turducken. Bacon ipsum dolor amet ground round rump short ribs bresaola kevin beef, pastrami chuck hamburger alcatra prosciutto. Ground round ham hock porchetta, cupim spare ribs doner ball tip bresaola tail chicken meatball. Salami meatball ground round boudin, chicken pork chop pastrami ham hock pork loin meatloaf. Rump sirloin corned beef boudin. Pancetta short loin boudin prosciutto chicken turducken ball tip corned beef t-bone sirloin tenderloin venison. Pork belly hamburger pig biltong, frankfurter capicola doner. Bresaola chicken sausage, bacon brisket swine boudin sirloin capicola flank kevin turducken. "
	keyMap := make(map[string]chan int)
	keys := strings.Split(in, " ")
	keys2 := strings.Split(in2, " ")

	finalOutputChannel := make(chan map[string]int)
	doneChannel := make(chan bool)

	for _, word := range keys {
		wordChannel := make(chan int)
		keyMap[word] = wordChannel
		go reducer(word, wordChannel, finalOutputChannel)
	}

	for _, word := range keys2 {
		if _, exists := keyMap[word]; !exists {
			wordChannel := make(chan int)
			keyMap[word] = wordChannel
			go reducer(word, wordChannel, finalOutputChannel)
		}
		// go reducer(word, wordChannel, finalOutputChannel)
	}

	go mapper(keys, keyMap, doneChannel)
	go mapper(keys2, keyMap, doneChannel)

	amDone := 2
	for amDone != 0 {
		<-doneChannel
		amDone--
	}

	for _, channel := range keyMap {
		close(channel)
		fmt.Println(<-finalOutputChannel)
	}
}

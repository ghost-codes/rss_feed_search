package search
    //Check the response once we return from the function

import (
	"log"
	"sync"
)

// A map of registered Matchers
var matchers = make(map[string]Matcher)

// Run search logic
func Run(searchTerm string) {
	// Retrieve list of all feed to search through
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	//Create an unbuffered channel to receive results
	results := make(chan *Result)

	//Setup a wait group so e can process all the feeds
	var waitGroup sync.WaitGroup

	//Set the number of go routines we need to wait for while
	//they process the individual feeds
	waitGroup.Add(len(feeds))

	//Launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		//Retrieve a matcher for the search.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//Launch the goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//Launch a goroutine to monitor when all the work is done
	go func() {
		//Wait for everything to be proccessd
		waitGroup.Wait()

		//Close the channel to signal the Display
		//function that we can exit the program.
		close(results)
	}()
	//Start displaying results as they are available and
	// return after the final result is displayed
	Display(results)
}

//Register is called to register a matcher fo use by the program
func Register(feedType string, matcher Matcher){
    if _,exists:=matchers[feedType];exists{
        log.Fatalln(feedType,"Matcher already registered")
    }

    log.Println("Register",feedType,"matcher")
    matchers[feedType]=matcher
}

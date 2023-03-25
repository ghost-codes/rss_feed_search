package search

import "log"

// defaultmatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// Match is launched as a goroutine for each individual feed to run
// searches concurently
func match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {

	//Perfrom the serach against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)

	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

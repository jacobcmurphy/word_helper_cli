package main

import (
	"flag"
	"fmt"
)

var (
	topics          string
	count           string
	follows         string
	preceeds        string
	regex           string
	soundsLike      string
	rhymesWith      string
	slantRhymesWith string
	means           string
	associatedWith  string
	verbose         bool
)

func init() {
	flag.StringVar(&rhymesWith, "rhyme", "", "A word that the desired word rhymes with")
	flag.StringVar(&slantRhymesWith, "slant-rhyme", "", "A word that the desired word has a slant rhyme with")
	flag.StringVar(&soundsLike, "sounds-like", "", "A word that the desired word should sound like")
	flag.StringVar(&topics, "topics", "", "Up to 5 words relating to the desired word")
	flag.StringVar(&follows, "comes-after", "", "The word before the desired word")
	flag.StringVar(&preceeds, "comes-before", "", "The word after the desired word")
	flag.StringVar(&regex, "spelling", "", "Spelling of desired word; use ? as a placeholder for 1 letter; use * as a placeholder for multiple letters")
	flag.StringVar(&means, "means", "", "Word or phrase with a similar meaning to the desired word")
	flag.StringVar(&associatedWith, "related-to", "", "Words commonly used in works including the desired word")
	flag.StringVar(&count, "c", "100", "Number of results to return")
	flag.BoolVar(&verbose, "v", false, "Show verbose API response")
	flag.Parse()
}

func main() {
	q := query{
		topics:          topics,
		count:           count,
		follows:         follows,
		preceeds:        preceeds,
		regex:           regex,
		soundsLike:      soundsLike,
		rhymesWith:      rhymesWith,
		slantRhymesWith: slantRhymesWith,
		means:           means,
		associatedWith:  associatedWith,
	}

	responses, err := getData(q)
	if err != nil {
		fmt.Println("Error getting data", err)
		return
	}
	display(responses)
}

func display(responses []apiResponse) {
	for _, response := range responses {
		fmt.Print(response.Word)
		if verbose {
			fmt.Print(fmt.Sprintf(" (%d)", response.Score))
		}
		fmt.Println()
	}
}

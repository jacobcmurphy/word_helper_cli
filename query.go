package main

import "net/url"

type query struct {
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
}

func (q query) String() string {
	params := make(url.Values)
	paramPtr := &params

	addToQuery("topics", q.topics, paramPtr)
	addToQuery("max", q.count, paramPtr)
	addToQuery("lc", q.follows, paramPtr)
	addToQuery("rc", q.preceeds, paramPtr)
	addToQuery("sp", q.regex, paramPtr)
	addToQuery("sl", q.soundsLike, paramPtr)
	addToQuery("rel_rhy", q.rhymesWith, paramPtr)
	addToQuery("rel_nry", q.slantRhymesWith, paramPtr)
	addToQuery("ml", q.means, paramPtr)
	addToQuery("rel_trg", q.associatedWith, paramPtr)

	return "https://api.datamuse.com/words?" + params.Encode()
}

func addToQuery(key, val string, q *url.Values) {
	if len(val) > 0 {
		q.Add(key, val)
	}
}

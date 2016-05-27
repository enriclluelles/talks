package main

//This file doesn't compile, it's just here for the slides :P

import (
	"fmt"
	"net/http"
	"time"
)

// START RATE OMIT
type RateLimitedError struct {
	error
	WaitFor time.Duration
}

// END RATE OMIT

// START PARSE OMIT
type ParseMessageError interface {
	error
	Line() int
	Column() int
}

// END PARSE OMIT

type Message struct {
}

// START STDERROR OMIT
func parse(*Message) (*Payload, error) {
	//...
	t, err := parseToken(rdr)
	if err != nil {
		return nil, fmt.Errorf("Payload serialization failed on line %d column %d", line, colum)
	}
	//...
}

// END STDERROR OMIT

// START STDERROR2 OMIT
func Get(url) (*Response, error) {
	t, err := client.Get(url)
	if t.StatusCode > 400 {
		//...
		switch t.StatusCode {
		//...
		case 418:
			when := checkNextWindow(ctx)
			return nil, fmt.Errorf("Rate limited can't send until %s", when.Format("3:04PM"))
		}
		//...
	}
	return t
}

// END STDERROR2 OMIT

func m() {
	http.Get
}

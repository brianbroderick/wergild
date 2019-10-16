package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/brianbroderick/agora"
)

/**
 * Stuff that should probably exist in Go (in my opinion)
 */

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func random_int(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func resolveQuery(query string) DgraphResponse {
	j := agora.QueryDgraph(query)

	var r DgraphResponse
	err := json.Unmarshal(j, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r
}

func resolveQueryWithVars(query string, variables map[string]string) DgraphResponse {
	j := agora.QueryDgraphWithVars(query, variables)

	var r DgraphResponse
	err := json.Unmarshal(j, &r)
	if err != nil {
		r.Errors = append(r.Errors, err)
	}

	return r
}

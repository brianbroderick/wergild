package mud

// func resolveQuery(query string) DgraphResponse {
// 	j := agora.QueryDgraph(query)

// 	var r DgraphResponse
// 	err := json.Unmarshal(j, &r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return r
// }

// func resolveQueryWithVars(query string, variables map[string]string) DgraphResponse {
// 	j := agora.QueryDgraphWithVars(query, variables)

// 	var r DgraphResponse
// 	err := json.Unmarshal(j, &r)
// 	if err != nil {
// 		r.Errors = append(r.Errors, err)
// 	}

// 	return r
// }

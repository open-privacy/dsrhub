package goautoneg

import (
	"mime"
	"sort"
	"strconv"
	"strings"
)

// Accept is a structure to represent a clause in an HTTP Accept Header.
type Accept struct {
	Type, SubType string
	Q             float64
	Params        map[string]string
}

// ParseAccept parses the given string as an Accept header as defined in
// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.1.
// Some rules are only loosely applied and might not be as strict as defined in the RFC.
func ParseAccept(header string) []Accept {
	parts := strings.Split(header, ",")
	clauses := []Accept{}

	for _, part := range parts {
		mt, params, err := mime.ParseMediaType(part)
		if err != nil {
			continue
		}

		accept := Accept{
			Q:      1.0, // "[...] The default value is q=1"
			Params: params,
		}

		// A media-type is defined as
		// "*/*" | ( type "/" "*" ) | ( type "/" subtype )
		types := strings.Split(mt, "/")
		switch {
		// This case is not defined in the spec keep it to mimic the original code.
		case len(types) == 1 && types[0] == "*":
			accept.Type = "*"
			accept.SubType = "*"
		case len(types) == 2:
			accept.Type = types[0]
			accept.SubType = types[1]
		default:
			continue
		}

		if qVal, ok := params["q"]; ok {
			// A parsing failure will set Q to 0.
			accept.Q, _ = strconv.ParseFloat(qVal, 64)
			delete(params, "q")
		}

		clauses = append(clauses, accept)
	}

	sort.SliceStable(clauses, func(i, j int) bool {
		return clauses[i].Q > clauses[j].Q
	})

	return clauses
}

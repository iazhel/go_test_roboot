package epd_api_model

import (
	"fmt"
	"time"
)

const (
	TestUrl = "http://192.168.1.200:8005"
	login    = "admin"
	password = "password"

	// Count of founded records. They will be printed in cycle.
	nOutput = 0
	// Prints every api call duration.
	printDurations = false
	// Pints api call url during execution.
	printUrl = false
	// Will print curl commands, if test fail.
	printCurlCommand = false
)

var urlHistory []string

// will return last action, if flag set.
func LastHistory() string {
	// if url printed before, will notthing log
	if printUrl && !printCurlCommand {
		return ""
	}
	if l := len(urlHistory); l > 0 {
		return urlHistory[l-1]
	}
	return "URL history is empty."
}

// Makes commands history.
func saveCommand(cmd, url, token, data string) string {
	if printCurlCommand {
		headers := "curl -H 'Authorization: Token " + token + "' -H 'Content-Type: application/json' "
		switch cmd {
		case "GET":
			return headers + url
		case "POST":
			return headers + "-X POST --data '" + data + "' " + url
		case "PATCH":
			return headers + "-X PATCH --data '" + data + "' " + url
		case "DELETE":
			return headers + "-X DELETE " + url
		}
	}
	return fmt.Sprintf("[BAD ANSWER]:  %-10s%s", cmd, url)
}

// Makes int slice of last n elements from (0..l)
// lastIndexes(100, 3) returns {97, 98, 99}
// Uses for print last slice elements.
func lastIndexes(l, n int) []int {
	r := []int{}
	for i := l - n; i < l; i++ {
		if i < 0 {
			continue
		}
		r = append(r, i)
	}
	return r
}


func printDurationCase(d time.Duration) {
	if printDurations {
		fmt.Printf("        %s", d)
	}
}

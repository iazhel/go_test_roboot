package access_api

import (
	"fmt"
	"time"
)

var (

	// Prints every api call duration.
	printDurations bool = false
	// Pints api call url during execution.
	printUrl bool = true
	// Will print curl commands, if test fail.
	printCurlCommand bool = false
)

var UrlHistory []string

// will return last action, if flag set.
func LastHistory() string {
	// if url printed before, will notthing log
	if !(printUrl || printCurlCommand){
		return ""
	}
	if l := len(UrlHistory); l > 0 {
		return UrlHistory[l-1]
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
		case "HEAD":
			return headers + "-X HEAD " + url
		case "OPTIONS":
			return headers + "-X OPTIONS " + url
		case "POST":
			return headers + "-X POST --data '" + data + "' " + url
		case "PUT":
			return headers + "-X PUT --data '" + data + "' " + url
		case "PATCH":
			return headers + "-X PATCH --data '" + data + "' " + url
		case "DELETE":
			return headers + "-X DELETE " + url
		}
	}
	return fmt.Sprintf("[BAD ANSWER]:  %-10s%s", cmd, url)
}

func printDuration(d time.Duration) {
	if printDurations {
		fmt.Printf("        %s", d)
	}
}

func (self *RestApiClient) PrintDurations(b bool) {
	printDurations = b
}

func (self *RestApiClient) PrintUrl(b bool) {
	printUrl = b
}

func (self *RestApiClient) PrintCurlCommand(b bool) {
	printCurlCommand = b
}


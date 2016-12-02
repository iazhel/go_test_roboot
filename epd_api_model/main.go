package epd_api_model

// v 1.3
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	stamp       = "2006-01-02 15:04:05" // time input format
	formatJSON  = "/?format=json"
	apiCallAuth = "/api-token-auth/"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RestApiClient struct {
	Url        string
	Token      TokenCredential
	httpClient *http.Client
	apiCall    string
}

type TokenCredential struct {
	Str string `json:"token"`
	Id  int    `json:"id"`
}

type UserCredential struct {
	Password string `json:"password"`
	UserName string `json:"username"`
}

// LOGIN to REST API server
func Login(restUrl, login, pass string) (*RestApiClient, error, time.Duration) {
	self := &RestApiClient{
		Url:        restUrl,
		httpClient: &http.Client{},
	}
	// POST login and password
	// make JSON with credential
	user := UserCredential{
		UserName: login,
		Password: pass,
	}
	b, err := json.Marshal(user)

	// make authorization request
	url := self.Url + apiCallAuth
	urlHistory = append(urlHistory, saveCommand("POST", url, self.Token.Str, string(b)))
	if printUrl {
		defer fmt.Println()
		fmt.Print("Login    ", url)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err, 0
	}
	request.Header.Set("Content-Type", "application/json")
	tStart := time.Now()

	self.httpClient.Timeout = time.Duration(time.Second * 30)
	// self.httpClient.Timeout = 0
	resp, err := self.httpClient.Do(request)
	if err != nil {
		return nil, err, time.Since(tStart)

	}
	d := time.Since(tStart)
	printDurationCase(d)
	defer resp.Body.Close()

	// get token from JSON response
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err, d
	}
	theToken := TokenCredential{}
	err = json.Unmarshal(response, &theToken)
	if err != nil {
		return nil, err, d

	}

	if len(theToken.Str) == 0 {
		return nil, errors.New(string(response)), d

	}
	self.Token = theToken
	return self, nil, d

}

func (self *RestApiClient) doRequest(apiCall string, data []byte, action string) (resp *http.Response, err error, d time.Duration) {
	// make request to REST API
	url := self.Url + apiCall + formatJSON
	urlHistory = append(urlHistory, saveCommand(action, url, self.Token.Str, string(data)))
	if printUrl {
		defer fmt.Println()
		fmt.Print((action + "         ")[:8], url)
	}
	request, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Token " + self.Token.Str)

	// Execute request
	tStart := time.Now()
	self.httpClient.Timeout = time.Duration(time.Second * 30)
	resp, err = self.httpClient.Do(request)
	if err != nil {
		return
	}
	d = time.Since(tStart)
	printDurationCase(d)

//	defer resp.Body.Close()
	/* read reaponse body
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	/ check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 && statusCode != 201 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	*/

	return
}

func (self *RestApiClient) GetBodies(apiCall string) ([][]byte, error, time.Duration) {

	// shall look Next url in json answer
	type withNext struct {
		Next string `json:"next"`
	}
	// the first url to get
	NextUrl := apiCall
	allBodies := make([][]byte, 0)
	duration := time.Duration(0) * time.Second
	for len(NextUrl) > 0 {
		unmarshaled := &withNext{}
		body, err, d := self.doGetRequest(NextUrl)
		if err != nil {
			return allBodies, err, duration + d
		}
		err = json.Unmarshal(body, unmarshaled)
		if err != nil {
			return allBodies, err, duration + d
		}
		NextUrl = unmarshaled.Next
		allBodies = append(allBodies, body)
		duration += d
	}
	return allBodies, nil, duration
}
func (self *RestApiClient) Get(apiCall string) (respBody []byte, err error, d time.Duration) {
	// read reaponse body
	resp, err, d := self.doRequest(apiCall, []byte{}, "GET")
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	return self.doGetRequest(apiCall)
}

func (self *RestApiClient) Head(apiCall string) (respBody []byte, err error, d time.Duration) {
	// read reaponse body
	resp, err, d := self.doRequest(apiCall, []byte{}, "HEAD")
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	return self.doGetRequest(apiCall)
}

func (self *RestApiClient) Options(apiCall string) (respBody []byte, err error, d time.Duration) {
	// read reaponse body
	resp, err, d := self.doRequest(apiCall, []byte{}, "OPTIONS")
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	return self.doGetRequest(apiCall)
}

func (self *RestApiClient) GetHeaders(apiCall string) (headers http.Header, err error, d time.Duration) {
	// read reaponse body
	resp, err, d := self.doRequest(apiCall, []byte{}, "GET")
	headers = resp.Header
	defer resp.Body.Close()
	return
}

func (self *RestApiClient) doGetRequest(apiCall string) (respBody []byte, err error, d time.Duration) {
	// make request to REST API
	url := self.Url + apiCall + formatJSON
	// if apiCall contains parameters /?, will use apiCall only
	if strings.Contains(apiCall, "/?") {
		url = apiCall
	}
	urlHistory = append(urlHistory, saveCommand("GET", url, self.Token.Str, ""))
	if printUrl {
		defer fmt.Println()
		fmt.Print("GET    ", url)
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Token "+self.Token.Str)

	// Execute request
	tStart := time.Now()
	self.httpClient.Timeout = time.Duration(time.Second * 30)
	// self.httpClient.Timeout = 0
	resp, err := self.httpClient.Do(request)
	if err != nil {
		return
	}
	d = time.Since(tStart)
	printDurationCase(d)
	defer resp.Body.Close()
	// read reaponse body
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	return
}

func (self *RestApiClient) doPostRequest(apiCall string, data []byte) (respBody []byte, err error, d time.Duration) {
	// make request to REST API
	url := self.Url + apiCall // + formatJSON
	urlHistory = append(urlHistory, saveCommand("POST", url, self.Token.Str, string(data)))
	if printUrl {
		defer fmt.Println()
		fmt.Print("POST   ", url)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Token "+self.Token.Str)

	// Execute request
	tStart := time.Now()
	self.httpClient.Timeout = time.Duration(time.Second * 30)
	// self.httpClient.Timeout = 0
	resp, err := self.httpClient.Do(request)
	if err != nil {
		return
	}
	d = time.Since(tStart)
	printDurationCase(d)
	defer resp.Body.Close()
	// read reaponse body
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 && statusCode != 201 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	return
}

func (self *RestApiClient) doPatchRequest(apiCall string, data []byte) (respBody []byte, err error, d time.Duration) {

	// make request to REST API
	url := self.Url + apiCall
	urlHistory = append(urlHistory, saveCommand("PATCH", url, self.Token.Str, string(data)))
	if printUrl {
		defer fmt.Println()
		fmt.Print("PATCH  ", url)
	}
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Token "+self.Token.Str)

	// Execute request
	tStart := time.Now()
	self.httpClient.Timeout = time.Duration(time.Second * 30)
	// self.httpClient.Timeout = 0
	resp, err := self.httpClient.Do(request)
	if err != nil {
		return
	}
	d = time.Since(tStart)
	printDurationCase(d)
	defer resp.Body.Close()
	// read reaponse body
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 && statusCode != 201 && statusCode != 204 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(respBody, errStruct)
		if err != nil {
			return nil, appendToError(err, statusCode, respBody), d
		}
		return nil, appendToError(err, statusCode, respBody), d
	}
	return
}

func (self *RestApiClient) doDelRequest(apiCall string) (err error, d time.Duration) {
	// make request to REST API
	url := self.Url + apiCall
	urlHistory = append(urlHistory, saveCommand("DELETE", url, self.Token.Str, ""))
	if printUrl {
		defer fmt.Println()
		fmt.Print("DELETE ", url)
	}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err, d
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Token "+self.Token.Str)

	// Execute request
	tStart := time.Now()
	self.httpClient.Timeout = time.Duration(time.Second * 30)
	// self.httpClient.Timeout = 0
	resp, err := self.httpClient.Do(request)
	if err != nil {
		return err, d
	}

	d = time.Since(tStart)
	printDurationCase(d)
	defer resp.Body.Close()
	// read reaponse body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, d
	}
	// check statusCode
	statusCode := resp.StatusCode
	if statusCode != 200 && statusCode != 204 {
		errStruct := &ErrorResponse{}
		err = json.Unmarshal(body, errStruct)
		if err != nil {
			return appendToError(err, statusCode, body), d
		}
		return appendToError(err, statusCode, body), d
	}
	return err, d
}

// appends to error status code and part of the body
func appendToError(err error, statusCode int, body []byte) error {
	trimmedBody := body
	if len(body) > 1000 {
		trimmedBody = append(body[:500], []byte("...")...)
	}
	if err == nil {
		err = errors.New("nil")
	}
	return errors.New(fmt.Sprintf("StatusCode: %d\nError: %s\nBody: %s",
		statusCode,
		err,
		trimmedBody))
}

// converts string "hh:mm:ss" or "d hh:mm:ss" into time.Duration type
func durationParse(s string) (d time.Duration, err error) {
	l := len(s)
	if l < 8 {
		return d, errors.New("Parse Duration: The string variable is too short")
	}
	seconds, errSec := strconv.ParseInt(s[l-2:], 10, 32)
	minutes, _ := strconv.ParseInt(s[l-5:l-3], 10, 32)
	hours, _ := strconv.ParseInt(s[l-8:l-6], 10, 32)
	d = time.Duration(seconds) * time.Second
	d += time.Duration(minutes) * time.Minute
	d += time.Duration(hours) * time.Hour
	// Append days
	if len(s) >= 10 {
		days, _ := strconv.ParseInt(s[:l-9], 10, 32)
		d += time.Duration(24*days) * time.Hour
	}
	return d, errSec
}

func durationMsParse(s string) (d time.Duration, err error) {
	l := len(s)
	if l < 4 {
		return d, errors.New("Parse Duration: The string variable is too short")
	}
	n := strings.Index(s, ".")
	if n == -1 {
		n = len(s)
	}
	milliseconds, _ := strconv.ParseInt(s[n:], 10, 32)
	seconds, _ := strconv.ParseInt(s[:n], 10, 32)
	d = time.Duration(seconds) * time.Second
	d += time.Duration(milliseconds) * time.Millisecond
	return
}
func durationMsFloatParse(s float64) (d time.Duration, err error) {
	seconds := int(s)
	milliseconds := int((s - float64(seconds)) * 1000)
	d = time.Duration(seconds) * time.Second
	d += time.Duration(milliseconds) * time.Millisecond
	return
}

// Uses before POST, PATCH. Checks duration string variable on format.
// If format wrong, will reset variable by DefaultValue. Example:
// timeFormatCheck(&getSt.LastRunTimeStr)
func durationsFormatCheck(d *string) {
	_, err := durationParse(*d)
	if err != nil {
		*d = "00:00:00" // DefaultValue
	}
}

// Uses before POST, PATCH. Checks time string variable on format.
// If format wrong, will reset variable by DefaultValue. Example:
// timeFormatCheck(&getSt.LastRunTimeStr)
func timeFormatCheck(t *string) {
	_, err := time.Parse(stamp, *t)
	if err != nil {
		*t = "0000-00-00 00:00:00" // DefaultValue
	}
}

func timeFormatCheckNotNull(t *string) {
	_, err := time.Parse(stamp, *t)
	if err != nil {
		*t = "2000-01-01 00:00:00" // DefaultValue
	}
}

// converts interface to string,
// trims pair brakes [] at the border.
func toString(a ...interface{}) string {
	out := fmt.Sprintf("%v", a)
	for strings.HasPrefix(out, "[") {
		out = strings.TrimPrefix(out, "[")
		out = strings.TrimSuffix(out, "]")
	}
	return out
}

func (self *RestApiClient) SetPrintUrl(b bool) {
}

package access_api

// v 1.5
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiCallAuth = "api-token-auth/"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RestApiClient struct {
	BaseUrl    string
	Token      TokenCredential
	HttpClient *http.Client
}

type TokenCredential struct {
	Str string `json:"token"`
	Id  int    `json:"id"`
}

type UserCredential struct {
	Password string `json:"password"`
	UserName string `json:"username"`
}

// Sends GET request,
func (self *RestApiClient) Get(apiCall string) ([]byte, error, time.Duration) {
	return self.receiveBody(apiCall, "GET", []int{200})
}

func (self *RestApiClient) Head(apiCall string) ([]byte, error, time.Duration) {
	return self.receiveBody(apiCall, "HEAD", []int{200})
}

func (self *RestApiClient) Options(apiCall string) ([]byte, error, time.Duration) {
	return self.receiveBody(apiCall, "OPTIONS", []int{200})
}

func (self *RestApiClient) Del(apiCall string) (error, time.Duration) {
	_, err, d := self.receiveBody(apiCall, "DELETE", []int{204})
	return err, d
}

func (self *RestApiClient) Post(apiCall string, data []byte) ([]byte, int, error, time.Duration) {
	return self.update(apiCall, "POST", data, []int{200, 201})
}

func (self *RestApiClient) Patch(apiCall string, data []byte) ([]byte, int, error, time.Duration) {
	return self.update(apiCall, "PATCH", data, []int{200, 201})
}

func (self *RestApiClient) Put(apiCall string, data []byte) ([]byte, int, error, time.Duration) {
	return self.update(apiCall, "PUT", data, []int{200, 201})
}

func (self *RestApiClient) GetHeaders(apiCall string) (http.Header, error, time.Duration) {
	// read reaponse body
	resp, err, d := self.doRequest(apiCall, "GET", []byte{}, []int{200})
	if err != nil {
		return nil, err, d
	}
	return resp.Header, nil, d
}

// Will returns all bodies in slice
// if GET response has many pages.
func (self *RestApiClient) GetAll(apiCall string) ([][]byte, error, time.Duration) {
	// shall look Next url in json answer
	type typeWithNext struct {
		Next string `json:"next"`
	}
	ofset := 0
	limit := 1000
	// make first url
	NextUrl := apiCall + fmt.Sprintf("&limit=%d&offset=%d", limit, ofset)
	allBodies := make([][]byte, 0)
	duration := time.Duration(0) * time.Second

	// Do cycle until Next field is existed
	for {
		body, err, d := self.Get(NextUrl)
		if err != nil {
			return allBodies, err, duration + d
		}
		// unmarshal got structure
		structWithNext := &typeWithNext{}
		err = json.Unmarshal(body, structWithNext)
		if err != nil {
			fmt.Println("JSON: ERROR" )
			return allBodies, err, duration + d
		}
		// make next url
		ofset += limit
		NextUrl = apiCall + fmt.Sprintf("&limit=%d&offset=%d", limit, ofset)
		allBodies = append(allBodies, body)
		duration += d
		// exit point
		if len(structWithNext.Next) == 0 {
			break
		}
	}
	return allBodies, nil, duration
}

func (self *RestApiClient) update(apiCall, action string, data []byte, expectedStatuses []int) (body []byte, status int, err error, d time.Duration) {
	// do request
	resp, err, d := self.doRequest(apiCall, action, data, expectedStatuses)
	if err != nil  {
		return []byte{}, 0, err, d
	}
	defer resp.Body.Close()
	status = resp.StatusCode
	// read body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, status, err, d
	}
	return body, status, nil, d
}

// does request and returns body
func (self *RestApiClient) receiveBody(apiCall string, action string, expectedStatuses []int) (body []byte, err error, d time.Duration) {

	// do request
	resp, err, d := self.doRequest(apiCall, action, []byte{}, expectedStatuses)
	if err != nil {
		return []byte{}, err, d
	}
	defer resp.Body.Close()

	// read body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body,  err, d
	}
	return body, nil, d
}

func (self *RestApiClient) doRequest(apiCall, action string, data []byte, expectedStatuses []int) (resp *http.Response, err error, d time.Duration) {

	// make url
	url := self.BaseUrl + apiCall

	// save history log
	UrlHistory = append(UrlHistory, saveCommand(action, url, self.Token.Str, string(data)))

	// print info
	if printUrl {
		fmt.Print("   ", (action + "         ")[:8], url)
	}
	if printCurlCommand {
		fmt.Printf("\n%s\n", LastHistory())
	}

	// make request
	request, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	if err != nil {
		return resp, err, 0
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Token "+self.Token.Str)

	// Execute request
	startTime := time.Now()
	self.HttpClient.Timeout = time.Duration(time.Second * 30)
	resp, err = self.HttpClient.Do(request)
	if err != nil {
		return resp, err, time.Since(startTime)
	}
	d = time.Since(startTime)

	// check statuses
	if !contains(expectedStatuses, resp.StatusCode) {
		body, err := returnBody(resp)
		if err != nil {
			return resp, err, d
		}
		return resp, makeError(errors.New("Not expected status."), resp.StatusCode, body), d
	}
	printDuration(d)
	return
}

// appends to error status code and part of the body
func makeError(err error, statusCode int, body []byte) error {
	trimmedBody := body
	if len(body) > 1000 {
		trimmedBody = append(body[:500], []byte("...")...)
	}
	if err == nil {
		err = errors.New("nil")
	}
	return errors.New(fmt.Sprintf("\nAPI access error: \nStatusCode: %d\nError: %s\nBody: %s",
		statusCode,
		err,
		trimmedBody))
}

func returnBody(self *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(self.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func Login(restUrl, login, pass string) (*RestApiClient, error, time.Duration) {
	self := &RestApiClient{
		BaseUrl:    restUrl,
		HttpClient: &http.Client{},
	}
	// POST login and password
	// make JSON with credential
	user := UserCredential{
		UserName: login,
		Password: pass,
	}
	b, err := json.Marshal(user)

	// make authorization request
	url := self.BaseUrl + apiCallAuth
	UrlHistory = append(UrlHistory, saveCommand("POST", url, self.Token.Str, string(b)))
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

	self.HttpClient.Timeout = time.Duration(time.Second * 30)
	// self.HttpClient.Timeout = 0
	resp, err := self.HttpClient.Do(request)
	if err != nil {
		return nil, err, time.Since(tStart)
	}
	d := time.Since(tStart)
	printDuration(d)
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

func contains(set []int, a int) bool {
	for _, n := range set {
		if a == n {
			return true
		}
	}
	return false
}

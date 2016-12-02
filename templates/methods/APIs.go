package methods

import (
	"access_api"
	"encoding/json"
	"fmt"
	"time"
)

type RestApiClient struct {
	*access_api.RestApiClient
}

func Login(restUrl, login, pass string) (*RestApiClient, error, time.Duration) {
	client, err, d := access_api.Login(restUrl, login, pass)
	return &RestApiClient{client,}, err, d
}

// API GET all /audit_history/view
func (self *RestApiClient) GetAuditHistoryView(a ...interface{}) (*AuditHistoryViewType, error, time.Duration) {
	out := &AuditHistoryViewType{}
	bodies, err, d := self.GetAll(AuditHistoryViewCall + "/?format=json")
	if err != nil {
		return out, err, d
	}
	for _, body := range bodies {
		gotSt := &AuditHistoryViewType{}
		err = json.Unmarshal(body, gotSt)
		if err != nil {
			return out, err, d
		}
		out.Count = gotSt.Count
		out.AuditHistoryViewFields = append(out.AuditHistoryViewFields, gotSt.AuditHistoryViewFields...)
	}
	return out, nil, d
}

// API POST /audit_history/view
func (self *RestApiClient) PostAuditHistoryView(postSt *AuditHistoryView) (*AuditHistoryView, error, time.Duration) {
	gotSt := &AuditHistoryView{}
	b, err := json.Marshal(postSt)
	if err != nil {
		return nil, err, 0
	}
	body, status, err, d := self.Post(AuditHistoryViewCall, b)
	if err != nil {
		return gotSt, err, d
	}
	// get a new structure
	if status == 201 {
		err = json.Unmarshal(body, gotSt)
		if err != nil {
			return gotSt, err, d
		}
	}
	return gotSt, nil, d
}

// API PATCH /audit_history/view
func (self *RestApiClient) PatchAuditHistoryView(id int, postSt *AuditHistoryView) (*AuditHistoryView, error, time.Duration) {
	gotSt := &AuditHistoryView{}

	b, err := json.Marshal(postSt)
	if err != nil {
		return nil, err, 0
	}
	body, status, err, d := self.Patch(AuditHistoryViewCall+fmt.Sprintf("%d/", id), b)
	if err != nil {
		return gotSt, err, d
	}
	if status == 201 {
		err = json.Unmarshal(body, gotSt)
		if err != nil {
			return gotSt, err, d
		}
	}
	return gotSt, nil, d
}

// API PUT /audit_history/view
func (self *RestApiClient) PutAuditHistoryView(id int, isPut *AuditHistoryView) (*AuditHistoryView, error, time.Duration) {
	gotSt := &AuditHistoryView{}
	b, err := json.Marshal(isPut)
	if err != nil {
		return nil, err, 0
	}
	body, status, err, d := self.Put(AuditHistoryViewCall+fmt.Sprintf("%d/", id), b)
	if err != nil {
		return gotSt, err, d
	}
	if status == 201 {
		err = json.Unmarshal(body, gotSt)
		if err != nil {
			return gotSt, err, d
		}
	}
	return gotSt, nil, d
}

// API OPTION /audit_history/view
func (self *RestApiClient) OptAuditHistoryView() ([]byte, error, time.Duration) {
	body, err, d := self.Options(AuditHistoryViewCall)
	if err != nil {
		return body, err, d
	}
	return body, nil, d
}

// API GET all /choices/block_size

func (self *RestApiClient) GetChoicesBlockSize(a ...interface{}) (*ChoicesBlockSizeType, error, time.Duration) {
	out := &ChoicesBlockSizeType{}
	bodies, err, d := self.GetAll(ChoicesBlockSizeCall + "/?format=json")
	if err != nil {
		return out, err, d
	}
	for _, body := range bodies {
		gotSt := &ChoicesBlockSizeType{}
		err = json.Unmarshal(body, gotSt)
		if err != nil {
			return out, err, d
		}
		out.Count = gotSt.Count
		out.ChoicesBlockSizeFields = append(out.ChoicesBlockSizeFields, gotSt.ChoicesBlockSizeFields...)
	}
	return out, nil, d
}

// API OPTION /choices/block_size
func (self *RestApiClient) OptChoicesBlockSize() ([]byte, error, time.Duration) {
	resp, err, d := self.Options(ChoicesBlockSizeCall)
	if err != nil {
		return resp, err, d
	}
	return resp, nil, d
}

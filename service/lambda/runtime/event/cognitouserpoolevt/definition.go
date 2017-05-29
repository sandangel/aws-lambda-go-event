package cognitouserpoolevt

import "encoding/json"

type CallerContext struct {
	AwsSdkVersion string `json:"awsSdkVersion"`
	ClientId      string `json:"clientId"`
}

type Request struct {
	UserAttributes map[string]string `json:"userAttributes"`
	ValidationData map[string]string `json:"validationData,omitempty"`
}

type Response struct {
	AutoConfirmUser bool `json:"autoConfirmUser,omitempty"`
	AutoVerifyPhone bool `json:"autoVerifyPhone,omitempty"`
	AutoVerifyEmail bool `json:"autoVerifyEmail,omitempty"`
}

type Event struct {
	UserName      string         `json:"userName"`
	Version       string         `json:"version"`
	Region        string         `json:"region"`
	UserPoolId    string         `json:"userPoolId"`
	TriggerSource string         `json:"triggerSource"`
	CallerContext *CallerContext `json:"callerContext"`
	Request       *Request       `json:"request"`
	Response      *Response      `json:"response"`
}

func (e *Event) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}

func (e *Event) GoString() string {
	return e.String()
}

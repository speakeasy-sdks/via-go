// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetUsersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A JSON array of user names
	GetUsers200ApplicationJSONStrings []string
}

func (o *GetUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUsersResponse) GetGetUsers200ApplicationJSONStrings() []string {
	if o == nil {
		return nil
	}
	return o.GetUsers200ApplicationJSONStrings
}
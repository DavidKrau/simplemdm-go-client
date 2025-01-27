package simplemdm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

// TODO: To implement : List all, List install...

// GetAssignmentGroup - Returns a specifc assignment group
func (c *Client) AppGet(id string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/apps/%s", c.HostName, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	app := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

// TODO: binary string represent the pkg file so, type need to be change (I think)

// AppCreate - Create new application
func (c *Client) AppCreate(
	appStoreId string,
	bundleId string,
	binary string,
	name string) (*SimplemdmDefaultStruct, error) {

	url := fmt.Sprintf("https://%s/api/v1/apps", c.HostName)

	body := &bytes.Buffer{}

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	q.Add("app_store_id", appStoreId)

	if len(bundleId) > 0 {
		q.Add("bundle_id", bundleId)
	}
	if len(binary) > 0 {
		q.Add("binary", binary)
	}
	if name != "" {
		q.Add("name", name)
	}

	req.URL.RawQuery = q.Encode()

	resBody, err := c.RequestResponse201(req)
	if err != nil {
		return nil, err
	}

	app := SimplemdmDefaultStruct{}
	err = json.Unmarshal(resBody, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

// AppDelete - Delete an application
func (c *Client) AppDelete(appId string) error {
	url := fmt.Sprintf("https://%s/api/v1/apps/%s", c.HostName, appId)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	body, err := c.RequestResponse204(req)

	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}

// TODO: binary string represent the pkg file so, type need to be change (I think)
// AppUpdate - Updates an application
func (c *Client) AppUpdate(appId string, binary string, name string, deployTo string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/apps/%s", c.HostName, appId)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	q.Add("name", name)

	q.Add("binary", binary) // Need to be defined as multipart/form-data.

	q.Add("deploy_to", deployTo)

	// encoding all parameters
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", writer.FormDataContentType())

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	app := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

package simplemdm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) ScriptCreate(name string, variableSupport bool, scriptFile string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/scripts/", c.HostName)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(scriptFile)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(scriptFile))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return nil, errFile1
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// adding parameter name with variable name
	q.Add("name", name)

	switch {
	case variableSupport:
		q.Add("variable_support", "true")
	default:
		q.Add("variable_support", "false")
	}

	// encoding all parameters
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", writer.FormDataContentType())

	body, err := c.RequestResponse201(req)
	if err != nil {
		return nil, err
	}

	script := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &script)
	if err != nil {
		return nil, err
	}

	return &script, nil
}

// DeleteProfile - Deletes an profile
func (c *Client) ScriptDelete(scriptID string) error {
	url := fmt.Sprintf("https://%s/api/v1/scripts/%s", c.HostName, scriptID)
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

// UpdateProfile - Updates an profile
func (c *Client) ScriptUpdate(name string, variableSupport bool, scriptFile string, ID string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/scripts/%s", c.HostName, ID)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(scriptFile)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(scriptFile))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return nil, errFile1
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// adding parameter name with variable name
	q.Add("name", name)

	switch {
	case variableSupport:
		q.Add("variable_support", "true")
	default:
		q.Add("variable_support", "false")
	}
	// encoding all parameters
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", writer.FormDataContentType())

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	script := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &script)
	if err != nil {
		return nil, err
	}

	return &script, nil
}

// GetProfile - Returns a specifc profile
func (c *Client) ScriptGet(scriptID string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/scripts/%s", c.HostName, scriptID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	script := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &script)
	if err != nil {
		return nil, err
	}

	return &script, nil
}

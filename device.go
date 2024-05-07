package simplemdm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) GetDevice(ID string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/devices/%s?include_secret_custom_attributes=true", c.HostName, ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	device := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

// CreateDevice - Create new device
func (c *Client) CreateDevice(name string, groupID string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/devices/", c.HostName)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// adding parameter name with variable name
	q.Add("name", name)
	q.Add("group_id", groupID)

	// encoding all parameters
	req.URL.RawQuery = q.Encode()

	body, err := c.RequestResponse201(req)
	if err != nil {
		return nil, err
	}

	device := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

// UpdateDevice - Updates an device
func (c *Client) UpdateDevice(ID string, name string, deviceMame string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/devices/%s", c.HostName, ID)
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// checking if defaultvalue exist
	q.Add("name", name)
	q.Add("device_name", deviceMame)
	// encoding all parameters
	req.URL.RawQuery = q.Encode()

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	device := SimplemdmDefaultStruct{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

// DeleteDevice - Deletes an device
func (c *Client) DeleteDevice(ID string) error {
	url := fmt.Sprintf("https://%s/api/v1/devices/%s", c.HostName, ID)
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

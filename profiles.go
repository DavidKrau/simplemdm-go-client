package simplemdm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ProfileGet - Returns a specific profile
func (c *Client) ProfileGet(profileID string) (*SimplemdmDefaultStruct, error) {
	url := fmt.Sprintf("https://%s/api/v1/profiles/%s", c.HostName, profileID)
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

// Assign profile to the group
func (c *Client) ProfileAssignToGroup(profileID string, deviceGroupID string) error {
	url := fmt.Sprintf("https://%s/api/v1/profiles/%s/device_groups/%s", c.HostName, profileID, deviceGroupID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	body, err := c.RequestResponse204or409(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}

// Unasiign profile from group
func (c *Client) ProfileUnAssignToGroup(profileID string, deviceGroupID string) error {
	url := fmt.Sprintf("https://%s/api/v1/profiles/%s/device_groups/%s", c.HostName, profileID, deviceGroupID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	body, err := c.RequestResponse204or409(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}

// Assign profile to the device
func (c *Client) ProfileAssignToDevice(profileID string, deviceID string) error {
	url := fmt.Sprintf("https://%s/api/v1/profiles/%s/devices/%s", c.HostName, profileID, deviceID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	body, err := c.RequestResponse204or409(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}

// Unasiign profile from device
func (c *Client) ProfileUnAssignToDevice(profileID string, deviceID string) error {
	url := fmt.Sprintf("https://%s/api/v1/profiles/%s/devices/%s", c.HostName, profileID, deviceID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	body, err := c.RequestResponse204or409(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}

package simplemdmAPIClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAssignmentGroup - Returns a specifc assignment group
func (c *Client) GetApp(id string) (*SimplemdmDefaultStruct, error) {
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

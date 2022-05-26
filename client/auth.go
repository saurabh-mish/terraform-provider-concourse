package client

import (
	"encoding/json"
	//"errors"
	"fmt"
	"net/http"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("username and password undefined")
	}
	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}
	fmt.Println(rb)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	fmt.Println(req)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(body)

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}
	fmt.Println(ar)

	return &ar, nil
}

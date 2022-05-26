package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "https://auth.prod.concourselabs.io/api/v1"
//const Username string = "saurabh+113@concourselabs.com"
//const Password string = "S@ura8hM2906"
//const GrantType string = "password"
//const Scope string = "INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
}

// authentication response
type AuthResponse struct {
	AccessToken  string     `json:"access_token"`
	TokenType    string     `json:"bearer"`
	RefreshToken string     `json:"refresh_token"`
	ExpiresIn    int16      `json:"expires_in"`
	Scope        string     `json:"scope"`
	Extra        Extra_info `json:"extra"`
	Jti          string     `json:"jti"`
}

type Extra_info struct {
	InstitutionID int
	UserID        int
	UserEmail     string
	GroupIDs      []int
	SurfaceIDs    []int
}

// NewClient -
func NewClient(host, username, password, grantType, scope *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
		Auth: AuthStruct{
			Username:  *username,
			Password:  *password,
			GrantType: *grantType,
			Scope:     *scope,
		},
	}

	if host != nil {
		c.HostURL = *host
	}
	fmt.Println(host)


	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.AccessToken

	return &c, nil
}


func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)


	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

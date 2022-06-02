package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"fmt"
)

type ConcourseAuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	GrantType string `json:"grant_type"`
	Scope string `json:"scope"`
}

type ConcourseAuthResp struct {
	AccessToken  string     `json:"access_token"`
	TokenType    string     `json:"token_type"`
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


type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       ConcourseAuthReq
}


const endpoint string = "https://auth.prod.concourselabs.io/api/v1/oauth"


func NewClient(user , pass *string) (*Client, error) {

	jsonData := ConcourseAuthReq{
		Username: *user,
		Password: *pass,
		GrantType: "password",
		Scope: "INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA",
	}

	formData := url.Values{
		"username":   {jsonData.Username},
		"password":   {jsonData.Password},
		"grant_type": {jsonData.GrantType},
		"scope":      {jsonData.Scope},
	}

	cl := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL: endpoint,
		Auth: jsonData,
	}

	//ar, err := c.SignIn()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", cl.HostURL), strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	body, err := cl.doRequest(req)
	if err != nil {
		return nil, err
	}

	// unmarshall
	log.Println("Unmarshalling ...\n")
	authResp := ConcourseAuthResp{}
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return nil, err
	}
	log.Println(authResp)
	log.Println("\nUnmarshal complete\n")

	cl.Token = authResp.AccessToken

	return &cl, nil
}


func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	return body, err
}


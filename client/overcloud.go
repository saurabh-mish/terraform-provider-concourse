package client

import (
	//"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ConcourseAuth struct {
	AccessToken  string     `json:"access_token"`
	TokenType    string     `json:"token_type"`
	RefreshToken string     `json:"refresh_token"`
	ExpiresIn    int16      `json:"expires_in"`
	/*
	// Ignore (do not process) the below fields in response data

	Scope        string     `json:"scope"`
	Extra        ExtraInfo  `json:"extra"`
	Jti          string     `json:"jti"`
	*/
}

/*
// Ignore inner struct

type ExtraInfo struct {
	InstitutionID int
	UserID        int
	UserEmail     string
	GroupIDs      []int
	SurfaceIDs    []int
}
*/


func CheckCredentials() (*string, *string) {
	var concourseUser string
	var concoursePass string
	var present bool

	concourseUser, present = os.LookupEnv("CONCOURSE_USERNAME")
	if concourseUser == "" || !present {
		fmt.Fprint(os.Stdout, "Environment variable 'CONCOURSE_USERNAME' empty or not set ...")
	} else {
		flag.StringVar(&concourseUser, "username", concourseUser, "Username (Email) for Concourse Labs")
	}

	concoursePass, present = os.LookupEnv("CONCOURSE_PASSWORD")
	if concoursePass == "" || !present {
		fmt.Fprint(os.Stdout, "Environment variable 'CONCOURSE_PASSWORD' empty or not set ...")
	} else {
		flag.StringVar(&concoursePass, "password", concoursePass, "Password for Concourse Labs")
	}

	flag.Parse()
	return &concourseUser, &concoursePass
}

func GetAuthData(user *string, pass *string) ConcourseAuth {
	var endpoint string = "https://auth.prod.concourselabs.io/api/v1/oauth/token"
	payload := url.Values{
		"username":   {*user},
		"password":   {*pass},
		"grant_type": {"password"},
		"scope":      {"INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA"},
	}

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(payload.Encode()))
	if err != nil {
		log.Fatalf("Unable to perform 'post' request: %v", err)
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept": {"application/json"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData ConcourseAuth
	json.Unmarshal(body, &jsonData)  // body is of []byte


	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return jsonData
	} else {
		return ConcourseAuth{}
	}
}


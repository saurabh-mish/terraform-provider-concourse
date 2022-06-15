package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const endp = "https://prod.concourselabs.io/api/model/v1"
const resource = "/institutions/113/attribute-tags"

func (c *Client) GetAttributeTag(tagID string) (*AttributeTag, error) {
	endpoint := endp + resource + "/" + tagID
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println("GET endpoint unavailable ...")
		return nil, err
	}

	apiToken := "Bearer " + c.Token
	req.Header.Add("Authorization", apiToken)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	attrTag := AttributeTag{}
	err = json.Unmarshal(body, &attrTag)
	if err != nil {
		log.Println("Error unmarshalling ...")
		return nil, err
	}

	return &attrTag, nil
}

func (c *Client) CreateAttributeTag(attTag AttrTagReq) (*AttributeTag, error) {

	endpoint := endp + resource

	jsonPayload := &AttrTagReq{
		Name:        attTag.Name,
		Description: attTag.Description,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(jsonPayload)
	req, err := http.NewRequest(http.MethodPost, endpoint, payloadBuf)
	if err != nil {
		log.Println("POST endpoint unavailable ...")
	}

	apiToken := "Bearer " + c.Token
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	attrTagResp := AttributeTag{}
	err = json.Unmarshal(body, &attrTagResp)
	if err != nil {
		log.Println("Error unmarshalling ...")
		return nil, err
	}

	return &attrTagResp, nil
}

func (c *Client) UpdateAttributeTag(tagID string, attTag AttrTagReq) (*AttributeTag, error) {

	endpoint := endp + resource + "/" + tagID

	jsonPayload := &AttrTagReq{
		Name:        attTag.Name,
		Description: attTag.Description,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(jsonPayload)
	req, err := http.NewRequest(http.MethodPut, endpoint, payloadBuf)
	if err != nil {
		log.Println("PUT endpoint unavailable ...")
	}

	apiToken := "Bearer " + c.Token
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	attrTagResp := AttributeTag{}
	err = json.Unmarshal(body, &attrTagResp)
	if err != nil {
		log.Println("Error unmarshalling ...")
		return nil, err
	}

	return &attrTagResp, nil
}

func (c *Client) DeleteAttributeTag(tagID string) error {
	endpoint := endp + resource + "/" + tagID
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		log.Println("DELETE endpoint unavailable ...")
		return err
	}

	apiToken := "Bearer " + c.Token
	req.Header.Add("Authorization", apiToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New(resp.Status)
	}

	return nil
}

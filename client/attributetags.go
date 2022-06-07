package client

import (
	"encoding/json"
	"log"
	"net/http"
)

const endp = "https://prod.concourselabs.io/api/model/v1"
const resource = "/institutions/113/attribute-tags"

// GetOrder - Returns a specifc order
func (c *Client) GetAttributeTag(tagID string) (*AttributeTag, error) {
	//attrTag := strconv.Itoa(tagId)
	endpoint := endp + resource + "/" + tagID

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println("Endpoint unavailable ...")
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

/*
func ReadAttributeTag(tagId int) {
	attrTag := strconv.Itoa(tagId)
	endpoint := url + resource + "/" + attrTag

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	var jsonData AttrTagResp
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &jsonData)

	log.Println(jsonData)
}
*/

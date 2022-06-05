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

package client

// JSON response data structure of attribute tag response object
// for create, read, and update operations
type AttributeTag struct {
	ID            int    `json:"id"`
	Version       int    `json:"version"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	CreatedBy     int    `json:"createdBy"`
	UpdatedBy     int    `json:"updatedBy"`
	InstitutionId int    `json:"institutionId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

type AttrTagReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

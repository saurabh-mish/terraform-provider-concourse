package client

/*
// Response data structure for create, update, and delete operations

{
	"id": 212891,
	"version": 0,
	"created": "2022-05-29T20:18:50.190Z",
	"updated": "2022-05-29T20:18:50.190Z",
	"createdBy": 101685,
	"updatedBy": 101685,
	"institutionId": 113,
	"name": "saurabh_test_name",
	"description": "saurabh_test_description"
}
*/

// JSON data structure of attribute tag response object
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

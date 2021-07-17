package model

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type ResponseType string

const (
	RJSON ResponseType = "application/json"
	RXML  ResponseType = "application/xml"
	RTEXT ResponseType = "application/plain"
)

type ResponseCode int

const (
	OK             ResponseCode = 200
	BAD_REQUEST    ResponseCode = 400
	NOT_FOUND      ResponseCode = 404
	INTERNAL_ERROR ResponseCode = 500
)

type URI struct {
	URL          string       `json:"url"`
	AuthToken    string       `json:"auth_token"`
	Method       Method       `json:"method"`
	Body         string       `json:"body"`
	ResponseType ResponseType `json:"response_type"`
	Delay        int          `json:"delay"`
}

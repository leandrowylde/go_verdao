package model

import "time"

type Result struct {
	CreatedAt    time.Time    `json:"created_at"`
	RequestURI   string       `json:"uri" gorm:"column:uri;type:varchar"`
	ResponseCode ResponseCode `json:"code" gorm:"column:code;type:int"`
	ResponseData string       `json:"data" gorm:"column:data;type:varchar"`
}

type ResultRepositoryInterface interface {
	SaveResult(r *Result) error
}

func NewResult(uri URI, code ResponseCode, data string) (*Result, error) {
	result := Result{
		CreatedAt:    time.Now(),
		RequestURI:   uri.URL,
		ResponseCode: code,
		ResponseData: data,
	}

	return &result, nil
}

// func (r *Result) Save() {

// }

package usecase

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cogny/go_verdao/application/model"
)

type TestAPIUseCase struct {
	API              model.URI
	ResultRepository model.ResultRepositoryInterface
}

func (a *TestAPIUseCase) Auth(token string) error {
	req, err := http.NewRequest(string(a.API.Method), a.API.URL, bytes.NewBufferString(a.API.Body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("accept", string(a.API.ResponseType))
	req.Header.Add("Authorization", a.API.AuthToken)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response from %v: %v", a.API.URL, body)
	return err
}

func (a *TestAPIUseCase) MakeRequest(uri model.URI) (responseCode model.ResponseCode, responseData string, err error) {
	req, err := http.NewRequest(string(a.API.Method), a.API.URL, bytes.NewBufferString(a.API.Body))
	if err != nil {
		return
	}
	req.Header.Add("accept", string(a.API.ResponseType))
	req.Header.Add("Authorization", a.API.AuthToken)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	responseCode = model.ResponseCode(resp.StatusCode)
	responseData = string(body)

	result, err := model.NewResult(uri, responseCode, responseData)
	if err != nil {
		return
	}

	err = a.ResultRepository.SaveResult(result)
	return
}

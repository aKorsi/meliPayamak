package meliPayamak

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RestService struct {
	UserName string
	PassWord string
}

type Response struct {
	Value        string `json:"Value"`
	RetStatus    int64  `json:"RetStatus"`
	StrRetStatus string `json:"StrRetStatus"`
}

func (s *RestService) makeRequest(jsonData map[string]string, op string) ([]byte, error) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://rest.payamak-panel.com/api/SendSMS/"+op, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error %s\n" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *RestService) SendSMS(to string, from string, text string, isFlash bool) (Response, error) {
	jsonData := map[string]string{
		"username": s.UserName,
		"password": s.PassWord,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "SendSMS")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *RestService) SendByBaseNumber(text string, to string, bodyId int64) (Response, error) {
	jsonData := map[string]string{
		"username": s.UserName,
		"password": s.PassWord,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "BaseServiceNumber")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *RestService) GetDeliveries2(recID int64) (Response, error) {
	jsonData := map[string]string{
		"username": s.UserName,
		"password": s.PassWord,
		"recID":    strconv.FormatInt(recID, 10),
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "GetDeliveries2")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *RestService) GetMessages(location int, from string, index string, count bool) (Response, error) {
	jsonData := map[string]string{
		"UserName": s.UserName,
		"PassWord": s.PassWord,
		"Location": strconv.Itoa(location),
		"From":     from,
		"Index":    index,
		"Count":    strconv.FormatBool(count),
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "GetMessages")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *RestService) GetCredit() (Response, error) {
	jsonData := map[string]string{
		"UserName": s.UserName,
		"PassWord": s.PassWord,
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "GetCredit")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *RestService) GetBasePrice() (Response, error) {
	jsonData := map[string]string{
		"UserName": s.UserName,
		"PassWord": s.PassWord,
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "GetBasePrice")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *RestService) GetUserNumbers() (Response, error) {
	jsonData := map[string]string{
		"UserName": s.UserName,
		"PassWord": s.PassWord,
	}
	res := Response{}
	data, err := s.makeRequest(jsonData, "GetUserNumbers")
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, err
}

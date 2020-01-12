package meliPayamak

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type SoapService struct {
	UserName string
	PassWord string
}

func (s *SoapService) makeRequestToSend(Data string) ([]byte, error) {
	dataAsByte := []byte(Data)

	response, err := http.Post("https://api.payamak-panel.com/post/send.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *SoapService) makeRequestToReceive(Data string) ([]byte, error) {
	dataAsByte := []byte(Data)
	response, err := http.Post("https://api.payamak-panel.com/post/receive.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *SoapService) makeRequestToContacts(Data string) ([]byte, error) {
	dataAsByte := []byte(Data)
	response, err := http.Post("https://api.payamak-panel.com/post/contacts.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s *SoapService) makeRequestToActions(Data string) ([]byte, error) {

	dataAsByte := []byte(Data)
	response, err := http.Post("https://api.payamak-panel.com/post/contacts.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s *SoapService) makeRequestToSchedule(Data string) ([]byte, error) {

	dataAsByte := []byte(Data)
	response, err := http.Post("https://api.payamak-panel.com/post/schedule.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s *SoapService) makeRequestToTickets(Data string) ([]byte, error) {

	dataAsByte := []byte(Data)
	response, err := http.Post("https://api.payamak-panel.com/post/tickets.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s *SoapService) makeRequestToUsers(Data string) ([]byte, error) {

	dataAsByte := []byte(Data)
	response, err := http.Post("https://api.payamak-panel.com/post/users.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		return nil, errors.New("The HTTP request failed with error \n\r" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *SoapService) SendSimpleSMS2(to string, from string, msg string, flash bool) (mxj.Map, error) {
	_func := "SendSimpleSMS2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendSimpleSMS(to []string, from string, msg string, flash bool) (mxj.Map, error) {
	_to := "<string>" + strings.Join(to, "</string><string>") + "</string>"
	_func := "SendSimpleSMS"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + _to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendSms(to []string, from string, msg string, flash bool, udh string, recid int64) (mxj.Map, error) {

	_to := "<string>" + strings.Join(to, "</string><string>") + "</string>"
	_recid := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recid)), "</long><long>"), "[]") + "</long>"
	_func := "SendSms"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + _to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash><udh>" + udh + "</udh><recId>" + _recid + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendWithDomain(to string, from string, msg string, flash bool, domain string) (mxj.Map, error) {

	_func := "SendWithDomain"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash><domainName>" + domain + "</domainName></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendByBaseNumber(text []string, to string, bodyId int) (mxj.Map, error) {

	_text := "<string>" + strings.Join(text, "</string><string>") + "</string>"
	_func := "SendByBaseNumber"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><text>" + _text + "</text><to>" + to + "</to><bodyId>" + strconv.Itoa(bodyId) + "</bodyId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendByBaseNumber2(text string, to string, bodyId int) (mxj.Map, error) {

	_func := "SendByBaseNumber2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><text>" + text + "</text><to>" + to + "</to><bodyId>" + strconv.Itoa(bodyId) + "</bodyId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) getMessages(location int, from string, index int, count int) (mxj.Map, error) {

	_func := "getMessages"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetSmsPrice(irancellCount int, mtnCount int, from string, text string) (mxj.Map, error) {
	_func := "GetSmsPrice"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><irancellCount>" + strconv.Itoa(irancellCount) + "</irancellCount><mtnCount>" + strconv.Itoa(mtnCount) + "</mtnCount><from>" + from + "</from><text>" + text + "</text></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMultiDelivery2(recId string) (mxj.Map, error) {

	_func := "GetMultiDelivery2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMultiDelivery(recId string) (mxj.Map, error) {

	_func := "GetMultiDelivery"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetInboxCount(isRead bool) (mxj.Map, error) {

	_func := "GetInboxCount"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><isRead>" + strconv.FormatBool(isRead) + "</isRead></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetDelivery2(recId string) (mxj.Map, error) {

	_func := "GetDelivery2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetDelivery(recId string) (mxj.Map, error) {

	_func := "GetDelivery"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetDeliveries3(recIds []string) (mxj.Map, error) {

	_recids := "<string>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recIds)), "</string><string>"), "[]") + "</string>"
	_func := "GetDeliveries3"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><recId>" + _recids + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetDeliveries2(recId string) (mxj.Map, error) {

	_func := "GetDeliveries2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetDeliveries(recIds []int64) (mxj.Map, error) {

	_recids := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recIds)), "</long><long>"), "[]") + "</long>"
	_func := "GetDeliveries"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><recId>" + _recids + "</recId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetCredit() (mxj.Map, error) {

	_func := "GetCredit"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSend(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

// //Receive API Operations
func (s *SoapService) RemoveMessages2(location int, msgIds string) (mxj.Map, error) {

	_func := "RemoveMessages2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><msgIds>" + msgIds + "</msgIds></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

//use Received or Sent or Removed or Deleted for location in the next method
func (s *SoapService) RemoveMessages(location string, msgIds string) (mxj.Map, error) {

	_func := "RemoveMessages"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + location + "</location><msgIds>" + msgIds + "</msgIds></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUsersMessagesByDate(location int, from string, index int, count int, dateFrom string, dateTo string) (mxj.Map, error) {

	_func := "GetUsersMessagesByDate"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetOutBoxCount() (mxj.Map, error) {

	_func := "GetOutBoxCount"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessagesWithChangeIsRead(location int, from string, index int, count int, isRead bool, changeIsRead bool) (mxj.Map, error) {

	_func := "GetMessagesWithChangeIsRead"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count><isRead>" + strconv.FormatBool(isRead) + "</isRead><changeIsRead>" + strconv.FormatBool(changeIsRead) + "</changeIsRead></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessagesReceptions(msgId int, fromRows int) (mxj.Map, error) {

	_func := "GetMessagesReceptions"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><msgId>" + strconv.Itoa(msgId) + "</msgId><fromRows>" + strconv.Itoa(fromRows) + "</fromRows></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessagesFilterByDate(location int, from string, index int, count int, dateFrom string, dateTo string, isRead bool) (mxj.Map, error) {

	_func := "GetMessagesFilterByDate"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo><isRead>" + strconv.FormatBool(isRead) + "</isRead></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessagesByDate(location int, from string, index int, count int, dateFrom string, dateTo string) (mxj.Map, error) {

	_func := "GetMessagesByDate"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessagesAfterIDJson(location int, from string, count int, msgId int) (mxj.Map, error) {

	_func := "GetMessagesAfterIDJson"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><count>" + strconv.Itoa(count) + "</count><msgId>" + strconv.Itoa(msgId) + "</msgId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessagesAfterID(location int, from string, count int, msgId int) (mxj.Map, error) {

	_func := "GetMessagesAfterID"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><count>" + strconv.Itoa(count) + "</count><msgId>" + strconv.Itoa(msgId) + "</msgId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessages(location int, from string, index int, count int) (mxj.Map, error) {

	_func := "GetMessages"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetMessageStr(location int, from string, index int, count int) (mxj.Map, error) {

	_func := "GetMessageStr"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) ChangeMessageIsRead(msgIds string) (mxj.Map, error) {

	_func := "ChangeMessageIsRead"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><msgIds>" + msgIds + "</msgIds></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToReceive(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

// //Users API Operations
func (s *SoapService) AddPayment(name string, family string, bankName string, code string, amount int, cardNumber string) (mxj.Map, error) {

	_func := "AddPayment"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><name>" + name + "</name><family>" + family + "</family><bankName>" + bankName + "</bankName><code>" + code + "</code><amount>" + strconv.Itoa(amount) + "</amount><cardNumber>" + cardNumber + "</cardNumber></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddUser(productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) (mxj.Map, error) {

	_func := "AddUser"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><productId>" + strconv.Itoa(productId) + "</productId><descriptions>" + descriptions + "</descriptions><mobileNumber>" + mobileNumber + "</mobileNumber><emailAddress>" + emailAddress + "</emailAddress><nationalCode>" + nationalCode + "</nationalCode><name>" + name + "</name><family>" + family + "</family><corporation>" + corporation + "</corporation><phone>" + phone + "</phone><fax>" + fax + "</fax><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><certificateNumber>" + certificateNumber + "</certificateNumber></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddUserComplete(productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int, howFindUs string, commercialCode string, saleId string, recommanderId string) (mxj.Map, error) {

	_func := "AddUserComplete"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><productId>" + strconv.Itoa(productId) + "</productId><descriptions>" + descriptions + "</descriptions><mobileNumber>" + mobileNumber + "</mobileNumber><emailAddress>" + emailAddress + "</emailAddress><nationalCode>" + nationalCode + "</nationalCode><name>" + name + "</name><family>" + family + "</family><corporation>" + corporation + "</corporation><phone>" + phone + "</phone><fax>" + fax + "</fax><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><certificateNumber>" + certificateNumber + "</certificateNumber><country>" + strconv.Itoa(country) + "</country><province>" + strconv.Itoa(province) + "</province><city>" + strconv.Itoa(city) + "</city><howFindUs>" + howFindUs + "</howFindUs><commercialCode>" + commercialCode + "</commercialCode><saleId>" + saleId + "</saleId><recommanderId>" + recommanderId + "</recommanderId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddUserWithLocation(productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int) (mxj.Map, error) {
	_func := "AddUserWithLocation"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><productId>" + strconv.Itoa(productId) + "</productId><descriptions>" + descriptions + "</descriptions><mobileNumber>" + mobileNumber + "</mobileNumber><emailAddress>" + emailAddress + "</emailAddress><nationalCode>" + nationalCode + "</nationalCode><name>" + name + "</name><family>" + family + "</family><corporation>" + corporation + "</corporation><phone>" + phone + "</phone><fax>" + fax + "</fax><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><certificateNumber>" + certificateNumber + "</certificateNumber><country>" + strconv.Itoa(country) + "</country><province>" + strconv.Itoa(province) + "</province><city>" + strconv.Itoa(city) + "</city></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddUserWithMobileNumber(productId int, mobileNumber string, firstName string, lastName string, email string) (mxj.Map, error) {

	_func := "AddUserWithMobileNumber"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><productId>" + strconv.Itoa(productId) + "</productId><mobileNumber>" + mobileNumber + "</mobileNumber><firstName>" + firstName + "</firstName><lastName>" + lastName + "</lastName><email>" + email + "</email></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddUserWithUserNameAndPass(targetUserName string, targetUserPassword string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) (mxj.Map, error) {

	_func := "AddUserWithUserNameAndPass"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUserName>" + targetUserName + "</targetUserName><targetUserPassword>" + targetUserPassword + "</targetUserPassword><productId>" + strconv.Itoa(productId) + "</productId><descriptions>" + descriptions + "</descriptions><mobileNumber>" + mobileNumber + "</mobileNumber><emailAddress>" + emailAddress + "</emailAddress><nationalCode>" + nationalCode + "</nationalCode><name>" + name + "</name><family>" + family + "</family><corporation>" + corporation + "</corporation><phone>" + phone + "</phone><fax>" + fax + "</fax><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><certificateNumber>" + certificateNumber + "</certificateNumber></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AuthenticateUser() (mxj.Map, error) {

	_func := "AuthenticateUser"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) ChangeUserCredit(amount int, description string, targetUsername string, GetTax bool) (mxj.Map, error) {

	_func := "ChangeUserCredit"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><amount>" + strconv.Itoa(amount) + "</amount><description>" + description + "</description><targetUsername>" + targetUsername + "</targetUsername><GetTax>" + strconv.FormatBool(GetTax) + "</GetTax></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) ChangeUserCredit2(amount string, description string, targetUsername string, GetTax bool) (mxj.Map, error) {
	_func := "ChangeUserCredit2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><amount>" + amount + "</amount><description>" + description + "</description><targetUsername>" + targetUsername + "</targetUsername><GetTax>" + strconv.FormatBool(GetTax) + "</GetTax></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) ChangeUserCreditBySeccondPass(seccondPassword string, amount int, description string, targetUsername string, GetTax bool) (mxj.Map, error) {
	_func := "ChangeUserCreditBySeccondPass"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" +
		s.UserName + "</UserName><seccondPassword>" + seccondPassword + "</seccondPassword><amount>" + strconv.Itoa(amount) + "</amount><description>" + description + "</description><targetUsername>" + targetUsername + "</targetUsername><GetTax>" + strconv.FormatBool(GetTax) + "</GetTax></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

func (s *SoapService) DeductUserCredit(amount int, description string) (mxj.Map, error) {

	_func := "DeductUserCredit"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" +
		s.UserName + "</UserName><PassWord>" +
		s.PassWord + "</PassWord><amount>" + strconv.Itoa(amount) + "</amount><description>" + description + "</description></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

func (s *SoapService) ForgotPassword(mobileNumber string, emailAddress string, targetUsername string) (mxj.Map, error) {

	_func := "ForgotPassword"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><mobileNumber>" + mobileNumber + "</mobileNumber><emailAddress>" + emailAddress + "</emailAddress><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetCities(provinceId int) (mxj.Map, error) {

	_func := "GetCities"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><provinceId>" + strconv.Itoa(provinceId) + "</provinceId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetEnExpireDate() (mxj.Map, error) {

	_func := "GetEnExpireDate"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetExpireDate() (mxj.Map, error) {

	_func := "GetExpireDate"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetProvinces() (mxj.Map, error) {

	_func := "GetProvinces"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserBasePrice(targetUsername string) (mxj.Map, error) {

	_func := "GetUserBasePrice"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserByUserID(pass string, userId int) (mxj.Map, error) {

	_func := "GetUserByUserID"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><pass>" + pass + "</pass><userId>" + strconv.Itoa(userId) + "</userId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserCredit(targetUsername string) (mxj.Map, error) {

	_func := "GetUserCredit"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserCreditBySecondPass(secondPassword string, targetUsername string) (mxj.Map, error) {

	_func := "GetUserCreditBySecondPass"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><secondPassword>" + secondPassword + "</secondPassword><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserDetails(targetUsername string) (mxj.Map, error) {
	_func := "GetUserDetails"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserIsExist(targetUsername string) (mxj.Map, error) {

	_func := "GetUserIsExist"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserNumbers() (mxj.Map, error) {

	_func := "GetUserNumbers"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserTransactions(targetUsername string, creditType string, dateFrom string, dateTo string, keyword string) (mxj.Map, error) {

	_func := "GetUserTransactions"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUsername>" + targetUsername + "</targetUsername><creditType>" + creditType + "</creditType><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo><keyword>" + keyword + "</keyword></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserWallet() (mxj.Map, error) {

	_func := "GetUserWallet"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUserWalletTransaction(dateFrom string, dateTo string, count int, startIndex int, payType string, PayLoc string) (mxj.Map, error) {

	_func := "GetUserWalletTransaction"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo><count>" + strconv.Itoa(count) + "</count><startIndex>" + strconv.Itoa(startIndex) + "</startIndex><payType>" + payType + "</payType><payLoc>" + PayLoc + "</payLoc></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetUsers() (mxj.Map, error) {

	_func := "GetUsers"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) HasFilter(text string) (mxj.Map, error) {

	_func := "HasFilter"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><text>" + text + "</text></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveUser(targetUsername string) (mxj.Map, error) {

	_func := "RemoveUser"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUsername>" + targetUsername + "</targetUsername></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RevivalUser(targetUsername string) (mxj.Map, error) {

	_func := "RevivalUser"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><targetUserName>" + targetUsername + "</targetUserName></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToUsers(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

// //Contact API Operations
func (s *SoapService) AddContact(groupIds string, firstname string, lastname string, nickname string, corporation string, mobilenumber string, phone string, fax string, birthdate string, email string, gender int, province int, city int, address string, postalCode string, additionaldate string, additionaltext string, descriptions string) (mxj.Map, error) {

	_func := "AddContact"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><groupIds>" + groupIds + "</groupIds><firstname>" + firstname + "</firstname><lastname>" + lastname + "</lastname><nickname>" + nickname + "</nickname><corporation>" + corporation + "</corporation><mobilenumber>" + mobilenumber + "</mobilenumber><phone>" + phone + "</phone><fax>" + fax + "</fax><birthdate>" + birthdate + "</birthdate><email>" + email + "</email><gender>" + strconv.Itoa(gender) + "</gender><province>" + strconv.Itoa(province) + "</province><city>" + strconv.Itoa(city) + "</city><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><additionaldate>" + additionaldate + "</additionaldate><additionaltext>" + additionaltext + "</additionaltext><descriptions>" + descriptions + "</descriptions></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddContactEvents(contactId int, eventName string, eventType int, eventDate string) (mxj.Map, error) {

	_func := "AddContactEvents"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><contactId>" + strconv.Itoa(contactId) + "</contactId><eventName>" + eventName + "</eventName><eventDate>" + eventDate + "</eventDate><eventType>" + strconv.Itoa(eventType) + "</eventType></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddGroup(groupName string, Descriptions string, showToChilds bool) (mxj.Map, error) {

	_func := "AddGroup"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><groupName>" + groupName + "</groupName><Descriptions>" + Descriptions + "</Descriptions><showToChilds>" + strconv.FormatBool(showToChilds) + "</showToChilds></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) ChangeContact(contactId int, mobilenumber string, firstname string, lastname string, nickname string, corporation string, phone string, fax string, email string, gender int, province int, city int, address string, postalCode string, additionaltext string, descriptions string, contactStatus int) (mxj.Map, error) {

	_func := "ChangeContact"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><contactId>" + strconv.Itoa(contactId) + "</contactId><mobilenumber>" + mobilenumber + "</mobilenumber><firstname>" + firstname + "</firstname><lastName>" + lastname + "</lastname><nickname>" + nickname + "</nickname><corporation>" + corporation + "</corporation><phone>" + phone + "</phone><fax>" + fax + "</fax><email>" + email + "</email><gender>" + strconv.Itoa(gender) + "</gender><province>" + strconv.Itoa(province) + "</province><city>" + strconv.Itoa(city) + "</city><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><additionaltext>" + additionaltext + "</additionaltext><descriptions>" + descriptions + "</descriptions><contactStatus>" + strconv.Itoa(contactStatus) + "</contactStatus></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) ChangeGroup(groupId int, groupName string, Descriptions string, showToChilds bool, groupStatus int) (mxj.Map, error) {

	_func := "ChangeGroup"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><groupId>" + strconv.Itoa(groupId) + "</groupId><groupName>" + groupName + "</groupName><Descriptions>" + Descriptions + "</Descriptions><showToChilds>" + strconv.FormatBool(showToChilds) + "</showToChilds><groupStatus>" + strconv.Itoa(groupStatus) + "</groupStatus></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) CheckMobileExistInContact(mobileNumber string) (mxj.Map, error) {

	_func := "CheckMobileExistInContact"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><mobileNumber>" + mobileNumber + "</mobileNumber></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetContactEvents(contactId int) (mxj.Map, error) {

	_func := "GetContactEvents"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><contactId>" + strconv.Itoa(contactId) + "</contactId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetContacts(groupId int, keyword string, from int, count int) (mxj.Map, error) {

	_func := "GetContacts"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><groupId>" + strconv.Itoa(groupId) + "</groupId><keyword>" + keyword + "</keyword><from>" + strconv.Itoa(from) + "</from><count>" + strconv.Itoa(count) + "</count></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetContactsByID(contactId int, status int) (mxj.Map, error) {

	_func := "GetContactsByID"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><contactId>" + strconv.Itoa(contactId) + "</contactId><status>" + strconv.Itoa(status) + "</status></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetGroups() (mxj.Map, error) {

	_func := "GetGroups"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) MergeGroups(originGroupId int, destinationGroupId int, deleteOriginGroup bool) (mxj.Map, error) {

	_func := "MergeGroups"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><originGroupId>" + strconv.Itoa(originGroupId) + "</originGroupId><destinationGroupId>" + strconv.Itoa(destinationGroupId) + "</destinationGroupId><deleteOriginGroup>" + strconv.FormatBool(deleteOriginGroup) + "</deleteOriginGroup></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveContact(mobilenumber string) (mxj.Map, error) {

	_func := "RemoveContact"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><mobilenumber>" + mobilenumber + "</mobilenumber></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveContactByContactID(contactId int) (mxj.Map, error) {

	_func := "RemoveContactByContactID"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><contactId>" + strconv.Itoa(contactId) + "</contactId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveGroup(groupId int) (mxj.Map, error) {

	_func := "RemoveGroup"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><groupId>" + strconv.Itoa(groupId) + "</groupId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToContacts(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

// //ACtions API Operations
func (s *SoapService) AddBranch(branchName string, owner int) (mxj.Map, error) {

	_func := "AddBranch"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><branchName>" + branchName + "</branchName><owner>" + strconv.Itoa(owner) + "</owner></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddBulk(from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) (mxj.Map, error) {

	_func := "AddBulk"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><from>" + from + "</from><branch>" + strconv.Itoa(branch) + "</branch><bulkType>" + strconv.Itoa(bulkType) + "</bulkType><title>" + title + "</title><message>" + message + "</message><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo><DateToSend>" + DateToSend + "</DateToSend><requestCount>" + strconv.Itoa(requestCount) + "</requestCount><rowFrom>" + strconv.Itoa(rowFrom) + "</rowFrom></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddBulk2(from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) (mxj.Map, error) {

	_func := "AddBulk2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><from>" + from + "</from><branch>" + strconv.Itoa(branch) + "</branch><bulkType>" + strconv.Itoa(bulkType) + "</bulkType><title>" + title + "</title><message>" + message + "</message><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo><DateToSend>" + DateToSend + "</DateToSend><requestCount>" + strconv.Itoa(requestCount) + "</requestCount><rowFrom>" + strconv.Itoa(rowFrom) + "</rowFrom></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddNewBulk(from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) (mxj.Map, error) {

	_func := "AddNewBulk"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><from>" + from + "</from><branch>" + strconv.Itoa(branch) + "</branch><bulkType>" + strconv.Itoa(bulkType) + "</bulkType><title>" + title + "</title><message>" + message + "</message><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo><DateToSend>" + DateToSend + "</DateToSend><requestCount>" + strconv.Itoa(requestCount) + "</requestCount><rowFrom>" + strconv.Itoa(rowFrom) + "</rowFrom></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddNumber(branchId int, mobileNumbers []string) (mxj.Map, error) {

	_mobileNumbers := "<string>" + strings.Join(mobileNumbers, "</string><string>") + "</string>"
	_func := "AddNumber"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><branchId>" + strconv.Itoa(branchId) + "</branchId><mobileNumbers>" + _mobileNumbers + "</mobileNumbers></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetBranchs(owner int) (mxj.Map, error) {

	_func := "GetBranchs"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><owner>" + strconv.Itoa(owner) + "</owner></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetBulk() (mxj.Map, error) {

	_func := "GetBulk"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetBulkCount(branch int, rangeFrom string, rangeTo string) (mxj.Map, error) {

	_func := "GetBulkCount"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><branch>" + strconv.Itoa(branch) + "</branch><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetBulkReceptions(bulkId int, fromRows int) (mxj.Map, error) {

	_func := "GetBulkReceptions"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><bulkId>" + strconv.Itoa(bulkId) + "</bulkId><fromRows>" + strconv.Itoa(fromRows) + "</fromRows></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetBulkStatus(bulkId int, sent int, failed int, status int) (mxj.Map, error) {

	_func := "GetBulkStatus"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><bulkId>" + strconv.Itoa(bulkId) + "</bulkId><sent>" + strconv.Itoa(sent) + "</sent><failed>" + strconv.Itoa(failed) + "</failed><status>" + strconv.Itoa(status) + "</status></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

//duplicate method
// func GetMessagesReceptions(  msgId int64, fromRows int) {

//     _func := "GetMessagesReceptions"
// 	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + UserName + "</UserName><PassWord>" + PassWord + "</PassWord><msgId>" + strconv.Itoa(msgId) + "</msgId><fromRows>" + strconv.Itoa(fromRows) + "</fromRows></"+ _func +"></soap:Body></soap:Envelope>"

//     makeRequestToActions(wsReq)
// }
func (s *SoapService) GetMobileCount(branch int, rangeFrom string, rangeTo string) (mxj.Map, error) {

	_func := "GetMobileCount"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><branch>" + strconv.Itoa(branch) + "</branch><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetSendBulk() (mxj.Map, error) {

	_func := "GetSendBulk"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetTodaySent() (mxj.Map, error) {

	_func := "GetTodaySent"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetTotalSent() (mxj.Map, error) {

	_func := "GetTotalSent"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveBranch(branchId int) (mxj.Map, error) {

	_func := "RemoveBranch"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><branchId>" + strconv.Itoa(branchId) + "</branchId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

func (s *SoapService) RemoveBulk(bulkId int) (mxj.Map, error) {

	_func := "RemoveBulk"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><bulkId>" + strconv.Itoa(bulkId) + "</bulkId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendMultipleSMS(to []string, from string, text []string, isflash bool, udh string, recId []int64, status string) (mxj.Map, error) {

	_to := "<string>" + strings.Join(to, "</string><string>") + "</string>"
	_text := "<string>" + strings.Join(text, "</string><string>") + "</string>"
	_recId := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recId)), "</long><long>"), "[]") + "</long>"
	_func := "SendMultipleSMS"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + _to + "</to><from>" + from + "</from><text>" + _text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><udh>" + udh + "</udh><recId>" + _recId + "</recId><status>" + status + "</status></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) SendMultipleSMS2(to []string, from []string, text []string, isflash bool, udh string, recId []int64, status string) (mxj.Map, error) {

	_to := "<string>" + strings.Join(to, "</string><string>") + "</string>"
	_text := "<string>" + strings.Join(text, "</string><string>") + "</string>"
	_from := "<string>" + strings.Join(from, "</string><string>") + "</string>"
	_recId := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recId)), "</long><long>"), "[]") + "</long>"
	_func := "SendMultipleSMS2"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + _to + "</to><from>" + _from + "</from><text>" + _text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><udh>" + udh + "</udh><recId>" + _recId + "</recId><status>" + status + "</status></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) UpdateBulkDelivery(bulkId int) (mxj.Map, error) {

	_func := "UpdateBulkDelivery"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><bulkId>" + strconv.Itoa(bulkId) + "</bulkId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToActions(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

// //Schedule API Operations
func (s *SoapService) AddMultipleSchedule(to []string, from string, text []string, isflash bool, scheduleDateTime []string, period string) (mxj.Map, error) {

	_to := "<string>" + strings.Join(to, "</string><string>") + "</string>"
	_text := "<string>" + strings.Join(text, "</string><string>") + "</string>"
	_schDates := "<string>" + strings.Join(scheduleDateTime, "</string><string>") + "</string>"
	_func := "AddMultipleSchedule"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + _to + "</to><from>" + from + "</from><text>" + _text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleDateTime>" + _schDates + "</scheduleDateTime><period>" + period + "</period></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddNewUsance(to string, from string, text string, isflash bool, scheduleStartDateTime string, countrepeat int, scheduleEndDateTime string, periodType string) (mxj.Map, error) {

	_func := "AddNewUsance"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + to + "</to><from>" + from + "</from><text>" + text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleStartDateTime>" + scheduleStartDateTime + "</scheduleStartDateTime><countrepeat>" + strconv.Itoa(countrepeat) + "</countrepeat><periodType>" + periodType + "</periodType></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddSchedule(to string, from string, text string, isflash bool, scheduleDateTime string, period string) (mxj.Map, error) {

	_func := "AddSchedule"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + to + "</to><from>" + from + "</from><text>" + text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleDateTime>" + scheduleDateTime + "</scheduleDateTime><period>" + period + "</period></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) AddUsance(to string, from string, text string, isflash bool, scheduleStartDateTime string, repeatAfterDays int, scheduleEndDateTime string) (mxj.Map, error) {

	_func := "AddUsance"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><to>" + to + "</to><from>" + from + "</from><text>" + text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleStartDateTime>" + scheduleStartDateTime + "</scheduleStartDateTime><repeatAfterDays>" + strconv.Itoa(repeatAfterDays) + "</repeatAfterDays><scheduleEndDateTime>" + scheduleEndDateTime + "</scheduleEndDateTime></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) GetScheduleStatus(scheduleId int) (mxj.Map, error) {

	_func := "GetScheduleStatus"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><scheduleId>" + strconv.Itoa(scheduleId) + "</scheduleId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveSchedule(scheduleId int) (mxj.Map, error) {

	_func := "RemoveSchedule"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><scheduleId>" + strconv.Itoa(scheduleId) + "</scheduleId></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}
func (s *SoapService) RemoveScheduleList(scheduleIdList []int) (mxj.Map, error) {

	_list := "<int>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scheduleIdList)), "</int><int>"), "[]") + "</int>"
	_func := "RemoveScheduleList"
	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + s.UserName + "</UserName><PassWord>" + s.PassWord + "</PassWord><scheduleIdList>" + _list + "</scheduleIdList></" + _func + "></soap:Body></soap:Envelope>"
	data, err := s.makeRequestToSchedule(wsReq)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(data)
}

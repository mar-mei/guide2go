package src

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://json.schedulesdirect.org/20141201/" // BaseURL: Request URL

var Token string // Token: SD login token

func sdLogin(username, password string) (err error) {

	var data = structToJson(SD_Login{Username: username, Password: password})

	body, err := postDataFromSD(data, "token")
	if err != nil {
		return
	}

	var response SD_Token
	json.Unmarshal(body.([]byte), &response)
	Token = response.Token

	return
}

func sdStatus() (response SD_Status, err error) {

	body, err := postDataFromSD("{}", "status")

	if err != nil {
		return
	}

	json.Unmarshal(body.([]byte), &response)

	logInfo("SD", fmt.Sprintf("Account expires: %s", response.Account.Expires))
	logInfo("SD", fmt.Sprintf("Lineups: %d", len(response.Lineups)))
	logInfo("SD", fmt.Sprintf("Max lineups: %d", response.Account.MaxLineups))
	logInfo("SD", fmt.Sprintf("System status: %s", response.SystemStatus[0].Status))
	logInfo("SD", fmt.Sprintf("System message: %s", response.SystemStatus[0].Message))

	return
}

func sdCountries() (response SD_Countries, err error) {

	body, err := postDataFromSD("", "countries")

	if err != nil {
		return
	}

	json.Unmarshal(body.([]byte), &response)

	return
}

func sdHeadends(lineup string) (response SD_Headends, err error) {

	body, err := postDataFromSD(lineup, "headends")

	if err != nil {
		return
	}
	json.NewDecoder(body.(io.Reader)).Decode(&response)

	return
}

func sdAddLineup(lineup string) (err error) {

	_, err = postDataFromSD(lineup, "lineups")
	return
}

func sdRemoveLineup(lineup string) (err error) {

	_, err = postDataFromSD(lineup, "delete")

	return
}

func sdChannelList(lineup string) (response SD_ChannelList, err error) {

	body, err := postDataFromSD(lineup, "channelList")

	if err != nil {
		return
	}

	json.Unmarshal(body.([]byte), &response)

	return
}

// Schedules
func sdGetSchedules(data string) (response SD_Schedules, err error) {

	body, err := postDataFromSD(data, "schedules")

	if err != nil {
		return
	}
	json.NewDecoder(body.(io.Reader)).Decode(&response)

	return
}

func sdGetPrograms(data string) (response SD_Programs, err error) {

	body, err := postDataFromSD(data, "programs")

	if err != nil {
		return
	}

	newBody, err := gUnzipData(body.(io.Reader))

	if err != nil {
		return
	}
	json.NewDecoder(newBody).Decode(&response)

	return
}

func sdGetMetadata(data string) (response SD_Metadata, err error) {

	body, err := postDataFromSD(data, "metadata")

	if err != nil {
		return
	}
	json.NewDecoder(body.(io.Reader)).Decode(&response)

	return
}

func postDataFromSD(data, reqType string) (body interface{}, err error) {

	var url, connectType string

	switch reqType {

	case "token":
		url = BaseURL + "token"
		connectType = "POST"

	case "status":
		url = BaseURL + "status"
		data = "{}"
		connectType = "GET"

	case "delete":
		url = BaseURL + "lineups/" + data
		connectType = "DELETE"

	case "countries":
		url = BaseURL + "available/countries"
		data = "{}"
		connectType = "GET"

	case "headends":
		url = BaseURL + "headends" + data
		data = "{}"
		connectType = "GET"

	case "lineups":
		url = BaseURL + "lineups/" + data
		data = "{}"
		connectType = "PUT"

	case "channelList":
		url = BaseURL + "lineups/" + data
		data = "{}"
		connectType = "GET"

	case "schedules":
		url = BaseURL + "schedules"
		connectType = "POST"

	case "programs":
		url = BaseURL + "programs"
		connectType = "POST"

	case "metadata":
		url = BaseURL + "metadata/programs"
		connectType = "POST"

	}

	logInfo("URL", url)

	var jsonStr = []byte(data)
	req, err := http.NewRequest(connectType, url, bytes.NewBuffer(jsonStr))

	if len(Token) > 0 {
		req.Header.Set("Token", Token)
	}

	if reqType == "programs" {
		req.Header.Set("Accept-Encoding", "deflate,gzip")
	}

	req.Header.Set("User-Agent", AppName)
	req.Header.Set("X-Custom-Header", AppName)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {

		checkErr(err, true)
		return

	}

	switch reqType {

	case "programs":
		body = resp.Body
		return
	case "metadata":
		body = resp.Body
		return
	case "schedules":
		body = resp.Body
		return
	}

	body, _ = ioutil.ReadAll(resp.Body)

	switch reqType {

	case "headends":
		return

	}

	var response SD_Status
	err = json.Unmarshal(body.([]byte), &response)

	if err != nil {
		return
	}

	err = checkTheServerStatus(response)

	return
}

func checkTheServerStatus(response SD_Status) (err error) {

	switch response.Code {

	case 0:
		if len(response.Message) > 0 {
			logInfo("SD", response.Message)
		}
		break

	default:
		logInfo("SD", response.Message)
		err = errors.New(response.Message)
		break

	}

	return
}

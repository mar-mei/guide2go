package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

// Init : Init Schedules Direct
func (sd *SD) Init() (err error) {

  sd.BaseURL = "https://json.schedulesdirect.org/20141201/"

  sd.Login = func() (err error) {

    sd.Req.URL = sd.BaseURL + "token"
    sd.Req.Type = "POST"
    sd.Req.Call = "login"
    sd.Req.Compression = false
    sd.Token = ""

    var login = Config.Account

    sd.Req.Data, err = json.MarshalIndent(login, "", "  ")
    if err != nil {
      ShowErr(err)
      return
    }

    err = sd.Connect()
    if err != nil {

      if sd.Resp.Login.Code != 0 {
        // SD Account problem
        return
      }

      return
    }

    showInfo("SD", fmt.Sprintf("Login...%s", sd.Resp.Login.Message))

    sd.Token = sd.Resp.Login.Token

    return
  }

  sd.Status = func() (err error) {

    fmt.Println()

    sd.Req.URL = sd.BaseURL + "status"
    sd.Req.Type = "GET"
    sd.Req.Data = nil
    sd.Req.Call = "status"
    sd.Req.Compression = false

    err = sd.Connect()
    if err != nil {
      return
    }

    showInfo("SD", fmt.Sprintf("Account Expires: %v", sd.Resp.Status.Account.Expires))
    showInfo("SD", fmt.Sprintf("Lineups: %d / %d", len(sd.Resp.Status.Lineups), sd.Resp.Status.Account.MaxLineups))

    for _, status := range sd.Resp.Status.SystemStatus {
      showInfo("SD", fmt.Sprintf("System Status: %s [%s]", status.Status, status.Message))
    }

    showInfo("G2G", fmt.Sprintf("Channels: %d", len(Config.Station)))

    return
  }

  sd.Countries = func() (err error) {

    sd.Req.URL = sd.BaseURL + "available/countries"
    sd.Req.Type = "GET"
    sd.Req.Data = nil
    sd.Req.Call = "countries"
    sd.Req.Compression = false

    err = sd.Connect()
    if err != nil {
      return
    }

    return
  }

  sd.Headends = func() (err error) {

    sd.Req.URL = fmt.Sprintf("%sheadends%s", sd.BaseURL, sd.Req.Parameter)
    sd.Req.Type = "GET"
    sd.Req.Data = nil
    sd.Req.Call = "headends"
    sd.Req.Compression = false

    err = sd.Connect()
    if err != nil {
      return
    }

    return
  }

  sd.Lineups = func() (err error) {

    sd.Req.URL = fmt.Sprintf("%slineups%s", sd.BaseURL, sd.Req.Parameter)
    sd.Req.Data = nil
    sd.Req.Call = "lineups"
    sd.Req.Compression = false

    err = sd.Connect()
    if err != nil {
      return
    }

    if len(sd.Resp.Lineup.Message) != 0 {
      showInfo("SD", sd.Resp.Lineup.Message)
    }

    return
  }

  sd.Schedule = func() (err error) {

    sd.Req.URL = fmt.Sprintf("%sschedules", sd.BaseURL)
    sd.Req.Type = "POST"
    sd.Req.Call = "schedule"
    sd.Req.Compression = false

    err = sd.Connect()
    if err != nil {
      return
    }

    return
  }

  sd.Program = func() (err error) {

    sd.Req.Type = "POST"
    sd.Req.Call = "program"
    sd.Req.Compression = true

    err = sd.Connect()
    if err != nil {
      return
    }

    return
  }

  return
}

// Connect : Connect to Schedules Direct

func (sd *SD) Connect() (err error) {

  var sdStatus SDStatus

  showInfo("URL", sd.Req.URL)

  req, err := http.NewRequest(sd.Req.Type, sd.Req.URL, bytes.NewBuffer(sd.Req.Data))
  if err != nil {
    return
  }

  if sd.Req.Compression == true {
    req.Header.Set("Accept-Encoding", "deflate,gzip")
  }

  req.Header.Set("Token", sd.Token)
  req.Header.Set("User-Agent", AppName)
  req.Header.Set("X-Custom-Header", AppName)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    ShowErr(err)
    return
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    ShowErr(err)
    return
  }

  sd.Resp.Body = body

  switch sd.Req.Call {

  case "login":
    err = json.Unmarshal(body, &sd.Resp.Login)
    if err != nil {
      ShowErr(err)
    }

    sdStatus.Code = sd.Resp.Login.Code
    sdStatus.Message = sd.Resp.Login.Message

  case "status":
    err = json.Unmarshal(body, &sd.Resp.Status)
    if err != nil {
      ShowErr(err)
    }

    sdStatus.Code = sd.Resp.Status.Code
    sdStatus.Message = sd.Resp.Status.Message

  case "countries":
    err = json.Unmarshal(body, &sd.Resp.Countries)
    if err != nil {
      ShowErr(err)
    }

  case "headends":
    err = json.Unmarshal(body, &sd.Resp.Headend)
    if err != nil {
      ShowErr(err)
    }

  case "lineups":
    err = json.Unmarshal(body, &sd.Resp.Lineup)
    if err != nil {
      ShowErr(err)
    }
    sd.Resp.Body = body

    sdStatus.Code = sd.Resp.Lineup.Code
    sdStatus.Message = sd.Resp.Lineup.Message

  case "schedule", "program":
    sd.Resp.Body = body

  }

  switch sdStatus.Code {

  case 0:
    //showInfo("SD", sd.Res.Message)

  default:
    err = fmt.Errorf("%s [SD API Error Code: %d]", sdStatus.Message, sdStatus.Code)
    ShowErr(err)

  }

  return
}

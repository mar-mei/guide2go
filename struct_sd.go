package main

import "time"

type SDStatus struct {
  Message string `json:"message"`
  Code    int    `json:"code"`
}

// SD : Schedules Direct API
type SD struct {
  BaseURL string
  Token   string

  // SD Request
  Req struct {
    URL         string
    Data        []byte
    Type        string
    Compression bool
    Parameter   string
    Call        string
  }

  // SD Response
  Resp struct {
    Body []byte

    // Login
    Login struct {
      Message  string    `json:"message"`
      Code     int       `json:"code"`
      ServerID string    `json:"serverID"`
      Datetime time.Time `json:"datetime"`

      Token string `json:"token"`
    }

    // Status
    Status struct {
      Account struct {
        Expires    time.Time     `json:"expires"`
        MaxLineups int64         `json:"maxLineups"`
        Messages   []interface{} `json:"messages"`
      } `json:"account"`
      Code    int    `json:"code"`
      Message string `json:"message"`

      Datetime       string `json:"datetime"`
      LastDataUpdate string `json:"lastDataUpdate"`
      Lineups        []struct {
        Lineup   string `json:"lineup"`
        Modified string `json:"modified"`
        Name     string `json:"name"`
        URI      string `json:"uri"`
      } `json:"lineups"`
      Notifications []interface{} `json:"notifications"`
      ServerID      string        `json:"serverID"`
      SystemStatus  []struct {
        Date    string `json:"date"`
        Message string `json:"message"`
        Status  string `json:"status"`
      } `json:"systemStatus"`
    }

    // Countries
    Countries struct {
      Caribbean []struct {
        FullName          string `json:"fullName"`
        OnePostalCode     bool   `json:"onePostalCode"`
        PostalCode        string `json:"postalCode"`
        PostalCodeExample string `json:"postalCodeExample"`
        ShortName         string `json:"shortName"`
      } `json:"Caribbean"`
      Europe []struct {
        FullName          string `json:"fullName"`
        OnePostalCode     bool   `json:"onePostalCode"`
        PostalCode        string `json:"postalCode"`
        PostalCodeExample string `json:"postalCodeExample"`
        ShortName         string `json:"shortName"`
      } `json:"Europe"`
      LatinAmerica []struct {
        FullName          string `json:"fullName"`
        OnePostalCode     bool   `json:"onePostalCode"`
        PostalCode        string `json:"postalCode"`
        PostalCodeExample string `json:"postalCodeExample"`
        ShortName         string `json:"shortName"`
      } `json:"Latin America"`
      NorthAmerica []struct {
        FullName          string `json:"fullName"`
        PostalCode        string `json:"postalCode"`
        PostalCodeExample string `json:"postalCodeExample"`
        ShortName         string `json:"shortName"`
      } `json:"North America"`
      Zzz []struct {
        FullName          string `json:"fullName"`
        OnePostalCode     bool   `json:"onePostalCode"`
        PostalCode        string `json:"postalCode"`
        PostalCodeExample string `json:"postalCodeExample"`
        ShortName         string `json:"shortName"`
      } `json:"ZZZ"`
    }

    // Headend
    Headend []struct {
      Headend string `json:"headend"`
      Lineups []struct {
        Lineup string `json:"lineup"`
        Name   string `json:"name"`
        URI    string `json:"uri"`
      } `json:"lineups"`
      Location  string `json:"location"`
      Transport string `json:"transport"`
    }

    // Lineup
    Lineup struct {
      // PUT
      ChangesRemaining int       `json:"changesRemaining"`
      Code             int       `json:"code"`
      Datetime         time.Time `json:"datetime"`
      Message          string    `json:"message"`
      Response         string    `json:"response"`
      ServerID         string    `json:"serverID"`

      // GET
      Map []struct {
        StationID string `json:"stationID"`
        Channel   string `json:"channel"`
      } `json:"map"`
      Stations []Station `json:"stations"`
    }
  }

  // SD API Calls
  Login     func() (err error)
  Status    func() (err error)
  Countries func() (err error)
  Headends  func() (err error)
  Lineups   func() (err error)
  Delete    func() (err error)
  Channels  func() (err error)
  Schedule  func() (err error)
  Program   func() (err error)
}

type Station struct {
  StationID           string   `json:"stationID"`
  Name                string   `json:"name"`
  Callsign            string   `json:"callsign"`
  Affiliate           string   `json:"affiliate,omitempty"`
  BroadcastLanguage   []string `json:"broadcastLanguage"`
  DescriptionLanguage []string `json:"descriptionLanguage"`
  StationLogo         []struct {
    URL    string `json:"URL"`
    Height int    `json:"height"`
    Width  int    `json:"width"`
    Md5    string `json:"md5"`
    Source string `json:"source"`
  } `json:"stationLogo,omitempty"`
  Logo struct {
    URL    string `json:"URL"`
    Height int    `json:"height"`
    Width  int    `json:"width"`
    Md5    string `json:"md5"`
  } `json:"logo,omitempty"`
  Broadcaster struct {
    City       string `json:"city"`
    State      string `json:"state"`
    Postalcode string `json:"postalcode"`
    Country    string `json:"country"`
  } `json:"broadcaster,omitempty"`
}

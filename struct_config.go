package main

type config struct {
  File       string   `yaml:"-"`
  ChannelIDs []string `yaml:"-"`

  Account struct {
    Username string `yaml:"Username" json:"username"`
    Password string `yaml:"Password" json:"password"`
  } `yaml:"Account"`

  Files struct {
    Cache string `yaml:"Cache"`
    XMLTV string `yaml:"XMLTV"`
  } `yaml:"Files"`

  Options struct {
    PosterAspect            string `yaml:"Poster Aspect"`
    Schedule                int    `yaml:"Schedule Days"`
    SubtitleIntoDescription bool   `yaml:"Subtitle into Description"`
  } `yaml:"Options"`

  Station []channel `yaml:"Station"`
}

type channel struct {
  Name        string        `yaml:"Name" json:"-" xml:"-"`
  DisplayName []DisplayName `yaml:"-" json:"-" xml:"display-name"`
  ID          string        `yaml:"ID" json:"stationID" xml:"id,attr"`
  Lineup      string        `yaml:"Lineup" json:"-" xml:"-"`
  Date        []string      `yaml:"-" json:"date"`
  Icon        Icon          `yaml:"-" json:"-" xml:"icon"`
}

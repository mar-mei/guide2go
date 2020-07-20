package main

import (
  "sync"
  "time"
)

type cache struct {
  Channel  map[string]G2GCache   `json:"Channel"`
  Program  map[string]G2GCache   `json:"Program"`
  Metadata map[string]G2GCache   `json:"Metadata"`
  Schedule map[string][]G2GCache `json:"Schedule"`

  sync.RWMutex `json:"-"`
}

// G2GCache : Cache data
type G2GCache struct {

  // Global
  Md5       string `json:"md5,omitempty"`
  ProgramID string `json:"programID,omitempty"`

  // Channel
  StationID         string   `json:"stationID,omitempty"`
  Name              string   `json:"name,omitempty"`
  Callsign          string   `json:"callsign,omitempty"`
  Affiliate         string   `json:"affiliate,omitempty"`
  BroadcastLanguage []string `json:"broadcastLanguage"`
  StationLogo       []struct {
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

  // Schedule
  AirDateTime     time.Time `json:"airDateTime,omitempty"`
  AudioProperties []string  `json:"audioProperties,omitempty"`
  Duration        int       `json:"duration,omitempty"`
  LiveTapeDelay   string    `json:"liveTapeDelay,omitempty"`
  New             bool      `json:"new,omitempty"`
  Ratings         []struct {
    Body string `json:"body"`
    Code string `json:"code"`
  } `json:"ratings,omitempty"`
  VideoProperties []string `json:"videoProperties,omitempty"`

  // Program
  Cast []struct {
    BillingOrder  string `json:"billingOrder"`
    CharacterName string `json:"characterName"`
    Name          string `json:"name"`
    NameID        string `json:"nameId"`
    PersonID      string `json:"personId"`
    Role          string `json:"role"`
  } `json:"cast"`
  Crew []struct {
    BillingOrder string `json:"billingOrder"`
    Name         string `json:"name"`
    NameID       string `json:"nameId"`
    PersonID     string `json:"personId"`
    Role         string `json:"role"`
  } `json:"crew"`
  ContentRating []struct {
    Body    string `json:"body"`
    Code    string `json:"code"`
    Country string `json:"country"`
  } `json:"contentRating"`
  Descriptions struct {
    Description1000 []struct {
      Description         string `json:"description"`
      DescriptionLanguage string `json:"descriptionLanguage"`
    } `json:"description1000"`
    Description100 []struct {
      DescriptionLanguage string `json:"descriptionLanguage"`
      Description         string `json:"description"`
    } `json:"description100"`
  } `json:"descriptions"`

  EpisodeTitle150   string   `json:"episodeTitle150,omitempty"`
  Genres            []string `json:"genres,omitempty"`
  HasEpisodeArtwork bool     `json:"hasEpisodeArtwork,omitempty"`
  HasImageArtwork   bool     `json:"hasImageArtwork,omitempty"`
  HasSeriesArtwork  bool     `json:"hasSeriesArtwork,omitempty"`

  Metadata []struct {
    Gracenote struct {
      Episode int `json:"episode"`
      Season  int `json:"season"`
    } `json:"Gracenote"`
  } `json:"metadata",omitempty`

  OriginalAirDate string `json:"originalAirDate,omitempty"`
  ResourceID      string `json:"resourceID,omitempty"`
  ShowType        string `json:"showType,omitempty"`
  Titles          []struct {
    Title120 string `json:"title120"`
  } `json:"titles"`

  // Metadata
  Data []Data `json:"data,omitempty"`
}

// SDSchedule : Schedules Direct schedule data
type SDSchedule struct {

  // Schedule
  Programs []struct {
    AirDateTime     time.Time `json:"airDateTime"`
    AudioProperties []string  `json:"audioProperties"`
    Duration        int       `json:"duration"`
    LiveTapeDelay   string    `json:"liveTapeDelay"`
    New             bool      `json:"new"`
    Md5             string    `json:"md5"`
    ProgramID       string    `json:"programID"`
    Ratings         []struct {
      Body string `json:"body"`
      Code string `json:"code"`
    } `json:"ratings"`
    VideoProperties []string `json:"videoProperties"`
  } `json:"programs"`
  StationID string `json:"stationID"`
}

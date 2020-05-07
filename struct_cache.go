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
  Descriptions struct {
    Description1000 []struct {
      Description         string `json:"description"`
      DescriptionLanguage string `json:"descriptionLanguage"`
    } `json:"description1000"`
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
  Data []struct {
    Aspect string `json:"aspect"`
    Height string `json:"height"`
    URI    string `json:"uri"`
    Width  string `json:"width"`
  } `json:"data,omitempty"`
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

// SDProgram : Schedules Direct program data
type SDProgram struct {

  // Program
  Cast []struct {
    BillingOrder  string `json:"billingOrder"`
    CharacterName string `json:"characterName"`
    Name          string `json:"name"`
    NameID        string `json:"nameId"`
    PersonID      string `json:"personId"`
    Role          string `json:"role"`
  } `json:"cast"`
  ContentAdvisory []string `json:"contentAdvisory"`
  ContentRating   []struct {
    Body    string `json:"body"`
    Code    string `json:"code"`
    Country string `json:"country"`
  } `json:"contentRating"`
  Crew []struct {
    BillingOrder string `json:"billingOrder"`
    Name         string `json:"name"`
    NameID       string `json:"nameId"`
    PersonID     string `json:"personId"`
    Role         string `json:"role"`
  } `json:"crew"`
  Descriptions struct {
    Description1000 []struct {
      Description         string `json:"description"`
      DescriptionLanguage string `json:"descriptionLanguage"`
    } `json:"description1000"`
  } `json:"descriptions"`
  EntityType        string   `json:"entityType"`
  EpisodeTitle150   string   `json:"episodeTitle150"`
  Genres            []string `json:"genres"`
  HasEpisodeArtwork bool     `json:"hasEpisodeArtwork"`
  HasImageArtwork   bool     `json:"hasImageArtwork"`
  HasSeriesArtwork  bool     `json:"hasSeriesArtwork"`
  Md5               string   `json:"md5"`

  Metadata []struct {
    Gracenote struct {
      Episode int `json:"episode"`
      Season  int `json:"season"`
    } `json:"Gracenote"`
  } `json:"metadata"`

  OriginalAirDate string `json:"originalAirDate"`
  ProgramID       string `json:"programID"`
  ResourceID      string `json:"resourceID"`
  ShowType        string `json:"showType"`
  Titles          []struct {
    Title120 string `json:"title120"`
  } `json:"titles"`
}

//SDMetadata : Schedules Direct meta data
type SDMetadata struct {
  Data []struct {
    Aspect string `json:"aspect"`
    Height string `json:"height"`
    URI    string `json:"uri"`
    Width  string `json:"width"`

    /*
          Caption struct {
            Content string `json:"content"`
            Lang    string `json:"lang"`
          } `json:"caption"`
       Category string `json:"category"`
       Primary  string `json:"primary"`
       Size     string `json:"size"`
       Text     string `json:"text"`
       Tier     string `json:"tier"`
    */
  } `json:"data"`
  ProgramID string `json:"programID"`
}

type SDStation struct {
  Map []struct {
    Channel   string `json:"channel"`
    StationID string `json:"stationID"`
  } `json:"map"`
  Metadata struct {
    Lineup    string `json:"lineup"`
    Modified  string `json:"modified"`
    Transport string `json:"transport"`
  } `json:"metadata"`
  Stations []struct {
    Affiliate         string   `json:"affiliate"`
    BroadcastLanguage []string `json:"broadcastLanguage"`
    Broadcaster       struct {
      City       string `json:"city"`
      Country    string `json:"country"`
      Postalcode string `json:"postalcode"`
      State      string `json:"state"`
    } `json:"broadcaster"`
    Callsign            string   `json:"callsign"`
    DescriptionLanguage []string `json:"descriptionLanguage"`
    Logo                struct {
      URL    string `json:"URL"`
      Height int    `json:"height"`
      Width  int    `json:"width"`
      Md5    string `json:"md5"`
    } `json:"logo,omitempty"`
    Name        string `json:"name"`
    StationID   string `json:"stationID"`
    StationLogo []struct {
      URL    string `json:"URL"`
      Height int    `json:"height"`
      Md5    string `json:"md5"`
      Source string `json:"source"`
      Width  int    `json:"width"`
    } `json:"stationLogo"`
  } `json:"stations"`
}

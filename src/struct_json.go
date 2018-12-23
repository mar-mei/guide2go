package src

import (
  "encoding/json"
  "time"
)

// Schedules Direct
type SD_Login struct {
  Username    string `json:"username"`
  Password    string `json:"password"`
}

type SD_Token struct {
  Code     int       `json:"code"`
  Message  string    `json:"message"`
  ServerID string    `json:"serverID"`
  Datetime time.Time `json:"datetime"`
  Token    string    `json:"token"`
}

type SD_Status struct {
  Message  string    `json:"message"`
  
  Account struct {
    Expires    time.Time     `json:"expires"`
    Messages   []interface{} `json:"messages"`
    MaxLineups int           `json:"maxLineups"`
  } `json:"account"`
  Lineups []struct {
    Lineup   string    `json:"lineup"`
    Modified time.Time `json:"modified"`
    URI      string    `json:"uri"`
    Name     string    `json:"name"`
  } `json:"lineups"`
  LastDataUpdate time.Time     `json:"lastDataUpdate"`
  Notifications  []interface{} `json:"notifications"`
  SystemStatus   []struct {
    Date    time.Time `json:"date"`
    Status  string    `json:"status"`
    Message string    `json:"message"`
  } `json:"systemStatus"`
  ServerID string    `json:"serverID"`
  Datetime time.Time `json:"datetime"`
  Code     int       `json:"code"`

}

type SD_Countries struct {
  NorthAmerica []struct {
    FullName          string `json:"fullName"`
    ShortName         string `json:"shortName"`
    PostalCodeExample string `json:"postalCodeExample"`
    PostalCode        string `json:"postalCode"`
  } `json:"North America"`
  Europe []struct {
    FullName          string `json:"fullName"`
    ShortName         string `json:"shortName"`
    PostalCodeExample string `json:"postalCodeExample"`
    PostalCode        string `json:"postalCode"`
    OnePostalCode     bool   `json:"onePostalCode,omitempty"`
  } `json:"Europe"`
  LatinAmerica []struct {
    FullName          string `json:"fullName"`
    ShortName         string `json:"shortName"`
    PostalCodeExample string `json:"postalCodeExample"`
    PostalCode        string `json:"postalCode"`
    OnePostalCode     bool   `json:"onePostalCode,omitempty"`
  } `json:"Latin America"`
  Caribbean []struct {
    FullName          string `json:"fullName"`
    ShortName         string `json:"shortName"`
    PostalCodeExample string `json:"postalCodeExample"`
    PostalCode        string `json:"postalCode"`
    OnePostalCode     bool   `json:"onePostalCode,omitempty"`
  } `json:"Caribbean"`
  ZZZ []struct {
    FullName          string `json:"fullName"`
    ShortName         string `json:"shortName"`
    PostalCodeExample string `json:"postalCodeExample"`
    PostalCode        string `json:"postalCode"`
    OnePostalCode     bool   `json:"onePostalCode"`
  } `json:"ZZZ"`
}

type SD_Headends []struct {
  Headend   string `json:"headend"`
  Transport string `json:"transport"`
  Location  string `json:"location"`
  Lineups   []struct {
    Name   string `json:"name"`
    Lineup string `json:"lineup"`
    URI    string `json:"uri"`
  } `json:"lineups"`
}

type SD_ChannelList struct {
  Map []struct {
    StationID string `json:"stationID"`
    Channel   string `json:"channel"`
  } `json:"map"`
  Stations []struct {
    StationID           string   `json:"stationID"`
    Name                string   `json:"name"`
    Callsign            string   `json:"callsign"`
    BroadcastLanguage   []string `json:"broadcastLanguage"`
    DescriptionLanguage []string `json:"descriptionLanguage"`
    Broadcaster         struct {
      City       string `json:"city"`
      State      string `json:"state"`
      Postalcode string `json:"postalcode"`
      Country    string `json:"country"`
    } `json:"broadcaster,omitempty"`
    StationLogo []struct {
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
    Affiliate        string `json:"affiliate,omitempty"`
    IsCommercialFree bool   `json:"isCommercialFree,omitempty"`
  } `json:"stations"`
  Metadata struct {
    Lineup    string    `json:"lineup"`
    Modified  time.Time `json:"modified"`
    Transport string    `json:"transport"`
  } `json:"metadata"`
}

type SD_Schedules []struct {
  StationID string `json:"stationID"`
  Programs  []struct {
    ProgramID       string    `json:"programID"`
    AirDateTime     time.Time `json:"airDateTime"`
    Duration        int       `json:"duration"`
    Md5             string    `json:"md5"`
    AudioProperties []string  `json:"audioProperties"`
    LiveTapeDelay   string    `json:"liveTapeDelay,omitempty"`
    New             bool      `json:"new,omitempty"`
    VideoProperties []string  `json:"videoProperties,omitempty"`
    Ratings         []struct {
      Body string `json:"body"`
      Code string `json:"code"`
    } `json:"ratings,omitempty"`
  } `json:"programs"`
  Metadata struct {
    Modified  time.Time `json:"modified"`
    Md5       string    `json:"md5"`
    StartDate string    `json:"startDate"`
  } `json:"metadata"`
}

type SD_Programs []struct {
  ResourceID string `json:"resourceID"`
  ProgramID  string `json:"programID"`
  Titles     []struct {
    Title120 string `json:"title120"`
  } `json:"titles"`
  Descriptions struct {
    Description100 []struct {
      DescriptionLanguage string `json:"descriptionLanguage"`
      Description         string `json:"description"`
    } `json:"description100"`
    Description1000 []struct {
      DescriptionLanguage string `json:"descriptionLanguage"`
      Description         string `json:"description"`
    } `json:"description1000"`
  } `json:"descriptions"`
  OriginalAirDate string   `json:"originalAirDate"`
  Genres          []string `json:"genres"`
  EpisodeTitle150 string   `json:"episodeTitle150,omitempty"`
  Metadata        []struct {
    Gracenote struct {
      Season  int `json:"season"`
      Episode int `json:"episode"`
    } `json:"Gracenote"`
  } `json:"metadata,omitempty"`
  Cast []struct {
    BillingOrder  string `json:"billingOrder"`
    Role          string `json:"role"`
    NameID        string `json:"nameId"`
    PersonID      string `json:"personId"`
    Name          string `json:"name"`
    CharacterName string `json:"characterName"`
  } `json:"cast,omitempty"`
  Crew []struct {
    BillingOrder string `json:"billingOrder"`
    Role         string `json:"role"`
    NameID       string `json:"nameId"`
    PersonID     string `json:"personId"`
    Name         string `json:"name"`
  } `json:"crew,omitempty"`
  EntityType        string `json:"entityType"`
  ShowType          string `json:"showType"`
  HasImageArtwork   bool   `json:"hasImageArtwork"`
  HasSeriesArtwork  bool   `json:"hasSeriesArtwork"`
  HasEpisodeArtwork bool   `json:"hasEpisodeArtwork,omitempty"`
  Md5               string `json:"md5"`
}

type SD_Metadata []struct {
  ProgramID string `json:"programID"`
  Data      []struct {
    Width    string `json:"width"`
    Height   string `json:"height"`
    URI      string `json:"uri"`
    Size     string `json:"size"`
    Aspect   string `json:"aspect"`
    Category string `json:"category"`
    Text     string `json:"text"`
    Primary  string `json:"primary"`
    Tier     string `json:"tier"`
  } `json:"data"`
}
//

// guide2go

type G2G_Programs []struct {
  ProgramID       string      `json:"programID"`
  AirDateTime     time.Time   `json:"airDateTime"`
  Duration        int         `json:"duration"`
  Md5             string      `json:"md5"`
  AudioProperties interface{} `json:"audioProperties"`
  VideoProperties []string    `json:"videoProperties"`
  Ratings         []struct {
    Body string `json:"body"`
    Code string `json:"code"`
  } `json:"ratings,omitempty"`
  New bool `json:"new,omitempty"`
}

type G2G_Channels struct {
  Callsign   string `json:"callsign"`
  LogoHeight int    `json:"logoHeight"`
  LogoURL    string `json:"logoURL"`
  LogoWidth  int    `json:"logoWidth"`
  Name       string `json:"name"`
  StationID  string `json:"stationID"`
}

type G2G_Schedules []struct {
  AirDateTime     time.Time `json:"airDateTime"`
  AudioProperties []string  `json:"audioProperties"`
  Duration        int       `json:"duration"`
  Md5             string    `json:"md5"`
  ProgramID       string    `json:"programID"`
  VideoProperties []string  `json:"videoProperties"`
  New             bool `json:"new,omitempty"`
}

type G2G_Program struct {
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
  Descriptions struct {
    Description100 []struct {
      Description         string `json:"description"`
      DescriptionLanguage string `json:"descriptionLanguage"`
    } `json:"description100"`
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
  Metadata          []struct {
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


type G2G_Metadata []struct {
  Aspect   string `json:"aspect"`
  Category string `json:"category"`
  Height   string `json:"height"`
  Primary  string `json:"primary"`
  Size     string `json:"size"`
  Text     string `json:"text"`
  Tier     string `json:"tier"`
  URI      string `json:"uri"`
  Width    string `json:"width"`
}

//


type SDGO_account struct {
  SdUsername string `json:"sd.username"`
  SdPassword string `json:"sd.password"`
}



func structToJson(data SD_Login)(string) {
  jsonString, _ := json.MarshalIndent(data, "", "  ")
  return string(jsonString)
}




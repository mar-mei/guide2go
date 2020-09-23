package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
  "sync"
)

// Cache : Cache file
var Cache cache

// Init : Inti cache
func (c *cache) Init() {

  if c.Schedule == nil {
    c.Schedule = make(map[string][]G2GCache)
  }

  c.Channel = make(map[string]G2GCache)

  if c.Program == nil {
    c.Program = make(map[string]G2GCache)
  }

  if c.Metadata == nil {
    c.Metadata = make(map[string]G2GCache)
  }

}

func (c *cache) Remove() {

  if len(Config.Files.Cache) != 0 {

    showInfo("G2G", fmt.Sprintf("%s [%s]", getMsg(0301), Config.Files.Cache))
    os.RemoveAll(Config.Files.Cache)

    c.Init()

  }

}

func (c *cache) AddStations(data *[]byte, lineup string) {

  c.Lock()
  defer c.Unlock()

  var g2gCache G2GCache
  var sdData SDStation

  err := json.Unmarshal(*data, &sdData)
  if err != nil {
    ShowErr(err)
    return
  }

  var channelIDs = Config.GetChannelList(lineup)

  for _, sd := range sdData.Stations {

    if ContainsString(channelIDs, sd.StationID) != -1 {

      g2gCache.StationID = sd.StationID
      g2gCache.Name = sd.Name
      g2gCache.Callsign = sd.Callsign
      g2gCache.Affiliate = sd.Affiliate
      g2gCache.BroadcastLanguage = sd.BroadcastLanguage
      g2gCache.Logo = sd.Logo

      c.Channel[sd.StationID] = g2gCache

    }

  }

}

func (c *cache) AddSchedule(data *[]byte) {

  c.Lock()
  defer c.Unlock()

  var g2gCache G2GCache
  var sdData []SDSchedule

  err := json.Unmarshal(*data, &sdData)
  if err != nil {
    ShowErr(err)
    return
  }

  for _, sd := range sdData {

    if _, ok := c.Schedule[sd.StationID]; !ok {
      c.Schedule[sd.StationID] = []G2GCache{}
    }

    for _, p := range sd.Programs {

      g2gCache.AirDateTime = p.AirDateTime
      g2gCache.AudioProperties = p.AudioProperties
      g2gCache.Duration = p.Duration
      g2gCache.LiveTapeDelay = p.LiveTapeDelay
      g2gCache.New = p.New
      g2gCache.Md5 = p.Md5
      g2gCache.ProgramID = p.ProgramID
      g2gCache.Ratings = p.Ratings
      g2gCache.VideoProperties = p.VideoProperties

      c.Schedule[sd.StationID] = append(c.Schedule[sd.StationID], g2gCache)

    }

  }

}

func (c *cache) AddProgram(gzip *[]byte, wg *sync.WaitGroup) {

  c.Lock()

  defer func() {
    c.Unlock()
    wg.Done()
  }()

  b, err := gUnzip(*gzip)
  if err != nil {
    ShowErr(err)
    return
  }

  var g2gCache G2GCache
  var sdData []SDProgram

  err = json.Unmarshal(b, &sdData)
  if err != nil {
    ShowErr(err)
    return
  }

  for _, sd := range sdData {

    g2gCache.Descriptions = sd.Descriptions
    g2gCache.EpisodeTitle150 = sd.EpisodeTitle150
    g2gCache.Genres = sd.Genres

    g2gCache.HasEpisodeArtwork = sd.HasEpisodeArtwork
    g2gCache.HasImageArtwork = sd.HasImageArtwork
    g2gCache.HasSeriesArtwork = sd.HasSeriesArtwork
    g2gCache.Metadata = sd.Metadata
    g2gCache.OriginalAirDate = sd.OriginalAirDate
    g2gCache.ResourceID = sd.ResourceID
    g2gCache.ShowType = sd.ShowType
    g2gCache.Titles = sd.Titles
    g2gCache.ContentRating = sd.ContentRating
    g2gCache.Cast = sd.Cast
    g2gCache.Crew = sd.Crew

    c.Program[sd.ProgramID] = g2gCache

  }

  return
}

func (c *cache) AddMetadata(gzip *[]byte, wg *sync.WaitGroup) {

  c.Lock()
  defer func() {
    c.Unlock()
    wg.Done()
  }()

  b, err := gUnzip(*gzip)
  if err != nil {
    ShowErr(err)
    return
  }

  var tmp = make([]interface{}, 0)

  var g2gCache G2GCache

  err = json.Unmarshal(b, &tmp)
  if err != nil {
    ShowErr(err)
    return
  }

  for _, t := range tmp {

    var sdData SDMetadata

    jsonByte, _ := json.Marshal(t)
    err = json.Unmarshal(jsonByte, &sdData)
    if err != nil {

      var sdError SDError
      err = json.Unmarshal(jsonByte, &sdError)
      if err == nil {

        if Config.Options.SDDownloadErrors == true {
          err = fmt.Errorf("%s [SD API Error Code: %d] Program ID: %s", sdError.Data.Message, sdError.Data.Code, sdError.ProgramID)
          ShowErr(err)
        }

      }

    } else {

      g2gCache.Data = sdData.Data
      c.Metadata[sdData.ProgramID] = g2gCache

    }

  }

  return
}

func (c *cache) GetAllProgramIDs() (programIDs []string) {

  for _, channel := range c.Schedule {

    for _, schedule := range channel {

      if ContainsString(programIDs, schedule.ProgramID) == -1 {
        programIDs = append(programIDs, schedule.ProgramID)
      }

    }

  }

  return
}

func (c *cache) GetRequiredProgramIDs() (programIDs []string) {

  var allProgramIDs = c.GetAllProgramIDs()

  for _, id := range allProgramIDs {

    if _, ok := c.Program[id]; !ok {

      if ContainsString(programIDs, id) == -1 {
        programIDs = append(programIDs, id)
      }

    }

  }

  return
}

func (c *cache) GetRequiredMetaIDs() (metaIDs []string) {

  for id, p := range c.Program {

    if p.HasImageArtwork == true {

      if len(id) > 10 {

        if _, ok := c.Metadata[id[:10]]; !ok {
          metaIDs = append(metaIDs, id[:10])
        }

      }

    }

  }

  return
}

func (c *cache) Open() (err error) {

  data, err := ioutil.ReadFile(Config.Files.Cache)

  if err != nil {
    c.Init()
    c.Save()
    return nil
  }

  // Open config file and convert Yaml to Struct (config)
  err = json.Unmarshal(data, &c)
  if err != nil {
    return
  }

  return
}

func (c *cache) Save() (err error) {

  c.Lock()
  defer c.Unlock()

  data, err := json.MarshalIndent(&c, "", "  ")
  if err != nil {
    return err
  }

  err = ioutil.WriteFile(Config.Files.Cache, data, 0644)
  if err != nil {
    return
  }

  return
}

func (c *cache) CleanUp() {

  var count int
  showInfo("G2G", fmt.Sprintf("Clean up Cache [%s]", Config.Files.Cache))

  var programIDs = c.GetAllProgramIDs()

  for id := range c.Program {

    if ContainsString(programIDs, id) == -1 {

      count++
      delete(c.Program, id)
      delete(c.Metadata, id[0:10])

    }

  }

  c.Channel = make(map[string]G2GCache)
  c.Schedule = make(map[string][]G2GCache)

  showInfo("G2G", fmt.Sprintf("Deleted Program Informations: %d", count))

  err := c.Save()
  if err != nil {
    ShowErr(err)
    return
  }
}

// Get data from cache
func (c *cache) GetTitle(id, lang string) (t []Title) {

  if p, ok := c.Program[id]; ok {

    var title Title

    for _, s := range p.Titles {
      title.Value = s.Title120
      title.Lang = lang
      t = append(t, title)
    }

  }

  if len(t) == 0 {
    var title Title
    title.Value = "No EPG Info"
    title.Lang = "en"
    t = append(t, title)
  }

  return
}

func (c *cache) GetSubTitle(id, lang string) (s SubTitle) {

  if p, ok := c.Program[id]; ok {

    if len(p.EpisodeTitle150) != 0 {

      s.Value = p.EpisodeTitle150
      s.Lang = lang

    } else {

      for _, d := range p.Descriptions.Description100 {

        s.Value = d.Description
        s.Lang = d.DescriptionLanguage

      }

    }

  }

  return
}

func (c *cache) GetDescs(id, subTitle string) (de []Desc) {

  if p, ok := c.Program[id]; ok {

    d := p.Descriptions

    var desc Desc

    for _, tmp := range d.Description1000 {

      switch Config.Options.SubtitleIntoDescription {

      case true:
        if len(subTitle) != 0 {
          desc.Value = fmt.Sprintf("[%s]\n%s", subTitle, tmp.Description)
          break
        }

        fallthrough
      case false:
        desc.Value = tmp.Description

      }

      desc.Lang = tmp.DescriptionLanguage

      de = append(de, desc)
    }

  }

  return
}

func (c *cache) GetCredits(id string) (cr Credits) {

  if Config.Options.Credits == true {

    if p, ok := c.Program[id]; ok {

      // Crew
      for _, crew := range p.Crew {

        switch crew.Role {

        case "Director":
          cr.Director = append(cr.Director, Director{Value: crew.Name})

        case "Producer":
          cr.Producer = append(cr.Producer, Producer{Value: crew.Name})

        case "Presenter":
          cr.Presenter = append(cr.Presenter, Presenter{Value: crew.Name})

        case "Writer":
          cr.Writer = append(cr.Writer, Writer{Value: crew.Name})

        }

      }

      // Cast
      for _, cast := range p.Cast {

        switch cast.Role {

        case "Actor":
          cr.Actor = append(cr.Actor, Actor{Value: cast.Name, Role: cast.CharacterName})

        }

      }

    }

  }

  return
}

func (c *cache) GetCategory(id string) (ca []Category) {

  if p, ok := c.Program[id]; ok {

    for _, g := range p.Genres {

      var category Category
      category.Value = g
      category.Lang = "en"

      ca = append(ca, category)

    }

  }

  return
}

func (c *cache) GetEpisodeNum(id string) (ep []EpisodeNum) {

  var seaseon, episode int

  if p, ok := c.Program[id]; ok {

    for _, m := range p.Metadata {

      seaseon = m.Gracenote.Season
      episode = m.Gracenote.Episode

      var episodeNum EpisodeNum

      if seaseon != 0 && episode != 0 {

        episodeNum.Value = fmt.Sprintf("%d.%d.", seaseon-1, episode-1)
        episodeNum.System = "xmltv_ns"

        ep = append(ep, episodeNum)
      }

    }

    if seaseon != 0 && episode != 0 {

      var episodeNum EpisodeNum
      episodeNum.Value = fmt.Sprintf("S%d E%d", seaseon, episode)
      episodeNum.System = "onscreen"
      ep = append(ep, episodeNum)

    }

    if len(ep) == 0 {

      var episodeNum EpisodeNum

      switch id[0:2] {

      case "EP":
        episodeNum.Value = id[0:10] + "." + id[10:]

      case "SH", "MV":
        episodeNum.Value = id[0:10] + ".0000"

      default:
        episodeNum.Value = id
      }

      episodeNum.System = "dd_progid"

      ep = append(ep, episodeNum)

    }

    if len(p.OriginalAirDate) > 0 {

      var episodeNum EpisodeNum
      episodeNum.Value = p.OriginalAirDate
      episodeNum.System = "original-air-date"
      ep = append(ep, episodeNum)

    }

  }

  return
}

func (c *cache) GetPreviouslyShown(id string) (prev *PreviouslyShown) {

  prev = &PreviouslyShown{}

  if p, ok := c.Program[id]; ok {
    prev.Start = p.OriginalAirDate
  }

  return
}

func (c *cache) GetIcon(id string) (i []Icon) {

  var aspects = []string{"2x3", "4x3", "3x4", "16x9"}
  var uri string
  var width, height int
  var err error

  switch Config.Options.PosterAspect {

  case "all":
    break

  default:
    aspects = []string{Config.Options.PosterAspect}

  }

  if m, ok := c.Metadata[id]; ok {

    for _, aspect := range aspects {

      var maxWidth, maxHeight int

      for _, icon := range m.Data {

        if icon.URI[0:7] != "http://" && icon.URI[0:8] != "https://" {

          icon.URI = fmt.Sprintf("https://json.schedulesdirect.org/20141201/image/%s", icon.URI)

        }

        if icon.Aspect == aspect {

          width, err = strconv.Atoi(icon.Width)
          if err != nil {
            return
          }

          height, err = strconv.Atoi(icon.Height)
          if err != nil {
            return
          }

          if width > maxWidth {
            maxWidth = width
            maxHeight = height
            uri = icon.URI
          }

        }

      }

      if maxWidth > 0 {
        i = append(i, Icon{Src: uri, Height: maxHeight, Width: maxWidth})
      }

    }

  }

  return
}

func (c *cache) GetRating(id, countryCode string) (ra []Rating) {

  if Config.Options.Rating.Guidelines == false {
    return
  }

  var add = func(code, body, country string) {

    switch Config.Options.Rating.CountryCodeAsSystem {

    case true:
      ra = append(ra, Rating{Value: code, System: country})

    case false:
      ra = append(ra, Rating{Value: code, System: body})

    }

  }

  /*
     var prepend = func(code, body, country string) {

       switch Config.Options.Rating.CountryCodeAsSystem {

       case true:
         ra = append([]Rating{{Value: code, System: country}}, ra...)

       case false:
         ra = append([]Rating{{Value: code, System: body}}, ra...)

       }

     }
  */

  if p, ok := c.Program[id]; ok {

    switch len(Config.Options.Rating.Countries) {

    case 0:
      for _, r := range p.ContentRating {

        if len(ra) == Config.Options.Rating.MaxEntries && Config.Options.Rating.MaxEntries != 0 {
          return
        }

        if countryCode == r.Country {
          add(r.Code, r.Body, r.Country)
        }

      }

      for _, r := range p.ContentRating {

        if len(ra) == Config.Options.Rating.MaxEntries && Config.Options.Rating.MaxEntries != 0 {
          return
        }

        if countryCode != r.Country {
          add(r.Code, r.Body, r.Country)
        }

      }

    default:
      for _, cCode := range Config.Options.Rating.Countries {

        for _, r := range p.ContentRating {

          if len(ra) == Config.Options.Rating.MaxEntries && Config.Options.Rating.MaxEntries != 0 {
            return
          }

          if cCode == r.Country {

            add(r.Code, r.Body, r.Country)

          }

        }

      }

    }

  }

  return
}

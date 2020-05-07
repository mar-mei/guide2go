package main

import (
  "encoding/json"
  "fmt"
  "path/filepath"
  "runtime"
  "strings"
  "sync"
  "time"
)

// Update : Update data from Schedules Direct and create the XMLTV file
func (sd *SD) Update(filename string) {

  Config.File = strings.TrimSuffix(filename, filepath.Ext(filename))

  err := Config.Open()
  if err != nil {
    return
  }

  err = sd.Init()
  if err != nil {
    return
  }

  if len(sd.Token) == 0 {

    err = sd.Login()
    if err != nil {
      return
    }

  }

  sd.GetData()

  runtime.GC()

  err = CreateXMLTV2(filename)
  if err != nil {
    ShowErr(err)
    return
  }

  Cache.CleanUp()

  runtime.GC()

}

// GetData : Get data from Schedules Direct
func (sd *SD) GetData() {

  var err error
  var wg sync.WaitGroup
  var count = 0

  err = Cache.Open()
  if err != nil {
    ShowErr(err)
    return
  }
  Cache.Init()

  // Channel list
  sd.Status()
  Cache.Channel = make(map[string]G2GCache)

  var lineup []string

  for _, l := range sd.Resp.Status.Lineups {
    lineup = append(lineup, l.Lineup)
  }

  for _, id := range lineup {

    //var channelIDs = Config.GetChannelList(id)

    sd.Req.Parameter = fmt.Sprintf("/%s", id)
    sd.Req.Type = "GET"

    err = sd.Lineups()

    Cache.AddStations(&sd.Resp.Body, id)

    /*
       for _, station := range sd.Resp.Lineup.Stations {

         if ContainsString(channelIDs, station.StationID) != -1 {

           fmt.Println(station.StationID, station.Name)
           //if station.Name == "Das Erste HD" {
           //fmt.Println(mapToJson(station))
           //}
           Cache.AddChannel(station)

         }

       }
    */

  }

  // Schedule
  showInfo("G2G", fmt.Sprintf("Download Schedule: %d Day(s)", Config.Options.Schedule))

  var limit = 5000

  var days = make([]string, 0)
  var channels = make([]interface{}, 0)

  for i := 0; i < Config.Options.Schedule; i++ {
    var nextDay = time.Now().Add(time.Hour * time.Duration(24*i))
    days = append(days, nextDay.Format("2006-01-02"))
  }

  for i, channel := range Config.Station {

    count++

    channel.Date = days
    channels = append(channels, channel)

    if count == limit || i == len(Config.Station)-1 {

      sd.Req.Data, err = json.Marshal(channels)
      if err != nil {
        ShowErr(err)
        return
      }

      sd.Schedule()

      wg.Add(1)
      go func() {

        Cache.AddSchedule(&sd.Resp.Body)

        wg.Done()

      }()

      count = 0

    }

  }

  wg.Wait()

  // Program and Metadata
  count = 0
  sd.Req.Data = []byte{}

  var types = []string{"programs", "metadata"}
  var programIds = Cache.GetRequiredProgrammIDs()
  var allIDs = Cache.GetAllProgrammIDs()
  var programs = make([]interface{}, 0)

  showInfo("G2G", fmt.Sprintf("Download Program Informations: New: %d / Cached: %d", len(programIds), len(allIDs)-len(programIds)))

  for _, t := range types {

    switch t {
    case "metadata":
      sd.Req.URL = fmt.Sprintf("%smetadata/programs", sd.BaseURL)
      sd.Req.Call = "metadata"
      limit = 500

    case "programs":

      sd.Req.URL = fmt.Sprintf("%sprograms", sd.BaseURL)
      sd.Req.Call = "programs"
      limit = 5000

    }

    for i, p := range programIds {

      count++

      var programID string

      switch t {
      case "metadata":
        programID = p[0:10]

      case "programs":
        programID = p

      }

      programs = append(programs, programID)

      if count == limit || i == len(programIds)-1 {

        sd.Req.Data, err = json.Marshal(programs)
        if err != nil {
          ShowErr(err)
          return
        }

        sd.Program()

        var gzip = sd.Resp.Body

        switch t {
        case "metadata":
          wg.Add(1)
          go func() {
            Cache.AddMetadata(&gzip)
            wg.Done()
          }()

        case "programs":
          wg.Add(1)
          go func() {
            Cache.AddProgram(&gzip)
            wg.Done()
          }()
        }

        count = 0
        programs = make([]interface{}, 0)

      }

    }

  }

  wg.Wait()

  err = Cache.Save()
  if err != nil {
    ShowErr(err)
    return
  }

  return
}

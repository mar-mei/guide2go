package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"

  "gopkg.in/yaml.v3"
)

// Configure : Configure config file
func Configure(filename string) (err error) {

  var menu Menu
  var entry Entry
  var sd SD

  Config.File = strings.TrimSuffix(filename, filepath.Ext(filename))

  err = Config.Open()
  if err != nil {
    return
  }

  sd.Init()

  if len(Config.Account.Username) != 0 || len(Config.Account.Password) != 0 {
    sd.Login()
    sd.Status()
  }

  for {

    menu.Entry = make(map[int]Entry)

    menu.Headline = fmt.Sprintf("%s [%s.yaml]", getMsg(0000), Config.File)
    menu.Select = getMsg(0001)

    // Exit
    entry.Key = 0
    entry.Value = getMsg(0010)
    menu.Entry[0] = entry

    // Account
    entry.Key = 1
    entry.Value = getMsg(0011)
    menu.Entry[1] = entry
    if len(Config.Account.Username) == 0 || len(Config.Account.Password) == 0 {
      entry.account()
      sd.Login()
      sd.Status()
    }

    // Add Lineup
    entry.Key = 2
    entry.Value = getMsg(0012)
    menu.Entry[2] = entry

    // Remove Lineup
    entry.Key = 3
    entry.Value = getMsg(0013)
    menu.Entry[3] = entry

    // Manage Channels
    entry.Key = 4
    entry.Value = getMsg(0014)
    menu.Entry[4] = entry

    // Create XMLTV file
    entry.Key = 5
    entry.Value = fmt.Sprintf("%s [%s]", getMsg(0016), Config.Files.XMLTV)
    menu.Entry[5] = entry

    var selection = menu.Show()

    entry = menu.Entry[selection]

    switch selection {

    case 0:
      os.Exit(0)

    case 1:
      entry.account()
      sd.Login()
      sd.Status()
      break

    case 2:
      entry.addLineup(&sd)
      sd.Status()
      break

    case 3:
      entry.removeLineup(&sd)
      sd.Status()
      break

    case 4:
      entry.manageChannels(&sd)
      sd.Status()
      break

    case 5:
      sd.Update(filename)
      break

    }

  }

  return
}

func (c *config) Open() (err error) {

  data, err := ioutil.ReadFile(fmt.Sprintf("%s.yaml", c.File))

  if err != nil {
    // File is missing, create new config file (YAML)
    c.InitConfig()
    err = c.Save()
    if err != nil {
      return
    }

    return nil
  }

  // Open config file and convert Yaml to Struct (config)
  err = yaml.Unmarshal(data, &c)
  if err != nil {
    return
  }

  return
}

func (c *config) Save() (err error) {

  data, err := yaml.Marshal(&c)
  if err != nil {
    return err
  }

  err = ioutil.WriteFile(fmt.Sprintf("%s.yaml", c.File), data, 0644)
  if err != nil {
    return
  }

  return
}

func (c *config) InitConfig() {

  // Files
  c.Files.Cache = fmt.Sprintf("%s_cache.json", c.File)
  c.Files.XMLTV = fmt.Sprintf("%s.xml", c.File)

  // Options
  c.Options.PosterAspect = "all"
  c.Options.Schedule = 7
  c.Options.SubtitleIntoDescription = false

  return
}

func (c *config) GetChannelList(lineup string) (list []string) {

  for _, channel := range c.Station {

    switch len(lineup) {

    case 0:
      list = append(list, channel.ID)

    default:
      if lineup == channel.Lineup {
        list = append(list, channel.ID)
      }

    }

  }

  return
}

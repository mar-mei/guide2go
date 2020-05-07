package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "time"
)

// AppName : App name
const AppName = "guide2go"

// Version : Version
const Version = "1.1.0"

// Config : Config file (struct)
var Config config

func main() {

  var configure = flag.String("configure", "", "= Create or modify a Schedules Direct configuration file. [filename.yaml]")
  var config = flag.String("config", "", "= Get data from Schedules Direct with configuration file. [filename.yaml]")
  var xmltv = flag.String("xmltv", "", "= Create XMLTV file from configuration. [filename.yaml]")

  var h = flag.Bool("h", false, ": Show help")
  flag.Parse()

  showInfo("G2G", fmt.Sprintf("Version: %s", Version))

  if *h {
    fmt.Println(AppName, Version)
    flag.Usage()
    os.Exit(0)
  }

  if len(*configure) != 0 {
    err := Configure(*configure)
    if err != nil {
      ShowErr(err)
    }
    os.Exit(0)
  }

  if len(*config) != 0 {
    var sd SD
    sd.Update(*config)
  }

  if len(*xmltv) != 0 {
    err := CreateXMLTV2(*xmltv)
    if err != nil {
      ShowErr(err)
    }
    time.Sleep(20 * time.Second)
    os.Exit(0)
  }

}

// ShowErr : Show error on screen
func ShowErr(err error) {
  var msg = fmt.Sprintf("[ERROR] %s", err)
  log.Println(msg)
}

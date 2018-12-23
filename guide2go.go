package main

import (
  "./src"
  "flag"
  "os"
)

func main() {

  var configure  = flag.String("configure",   "",     "= Create or modify a Schedules Direct configuration file. [filename]")
  var config     = flag.String("config",      "",     "= Get data from Schedules Direct with configuration file. [filename]")

  var h         = flag.Bool("h", false, ": Show help")
  flag.Parse()

  // Show help
  if *h {
    flag.Usage()
    os.Exit(0)
  }

  // Init
  src.Init()

  // Create or modify a configuration file
  if len(*configure) > 0 {
    src.CreateModifyConfigurationFile(*configure)
    os.Exit(0)
  }

  // Get data from Schedules Direct
  if len(*config) > 0 {
    
    src.GetData(*config)
    src.CreateXMLTV(*config)
    src.CleanUpTheCache(*config)
    
    os.Exit(0)
  }

  return
}
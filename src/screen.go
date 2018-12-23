package src

import (
  "fmt"
  "log"
)

var Max = 11

func screenInput(screenMsg int)(input string) {
  
  fmt.Print(getMsg(screenMsg) + ": " + getSpace(len(getMsg(screenMsg))))
  fmt.Scanln(&input)
  return

}

func logInfo(key, value string) {

  key = fmt.Sprintf("[%s] ", key)
  var logMsg = key + getSpace(len(key)) + value
  log.Println(logMsg)

}

func showHeadline(screenMsg int) {
  
  fmt.Println()
  fmt.Println(getMsg(screenMsg))
  underline(len(getMsg(screenMsg)))

}

func createScreenMenu(menu map[int]string) {

  for i := 1; i < len(menu); i++ {
    var m = fmt.Sprintf("%d. %s", i, menu[i])
    fmt.Println(m)
  }

  var m = fmt.Sprintf("%d. %s", 0, menu[0])
  fmt.Println(m)

}

func getMsg(screenMsg int)(msg string) {

  switch(screenMsg) {

    case 001: msg = "Schedules Direct username"
    case 002: msg = "Schedules Direct password"
    case 003: msg = "Select entry"
    case 004: msg = "Which lineup should be deleted?"
    case 005: msg = "Enter country"
    case 006: msg = "Enter postalcode"
    case 007: msg = "Select Provider"
    case 020: msg = "(Y) Add channel, (N) Remove channel, (ALL) All channels, (NONE) No channels, (SKIP) Skip all channels"

    case 100: msg = "Create configuration file"
    case 101: msg = "Edit configuration file"
    case 102: msg = "Change credentials from Schedules Direct"
    case 103: msg = "Add lineup from Schedules Direct"
    case 104: msg = "Remove lineup from Schedules Direct"
    case 105: msg = "Account status"
    case 106: msg = "Manage channels"
    case 110: msg = "System info"
    case 111: msg = "Clean up cache file"

    case 150: msg = "Get data from Schedules Direct"
    case 190: msg = "Create XMLTV file"

  }

  return
}

func getSpace(stringLen int)(space string) {

  for i := 0; i < Max - stringLen; i++ {
    space = space + " "
  }
  return

}

func underline(stringLen int) {

  var line string
  for i := 0; i < stringLen; i++ {
    line = line + "-"
  }

  fmt.Println(line)

}
package main

import (
  "errors"
  "fmt"
  "log"
  "sort"
  "strconv"
)

func getMsg(code int) (msg string) {

  switch code {

  // Menu entries
  case 0000:
    msg = "Configuration"
  case 0001:
    msg = "Select Entry"
  case 0010:
    msg = "Exit"
  case 0011:
    msg = "Schedules Direct Account"
  case 0012:
    msg = "Add Lineup"
  case 0013:
    msg = "Remove Lineup"
  case 0014:
    msg = "Manage Channels"
  case 0015:
    msg = "Exit"
  case 0016:
    msg = "Create XMLTV File"

  case 0100:
    msg = "Username"
  case 0101:
    msg = "Password"

  case 0200:
    msg = "Cancel"
  case 0201:
    msg = "Select Country"
  case 0202:
    msg = "Postal Code"
  case 0203:
    msg = "Select Provider"
  case 0204:
    msg = "Select Lineup"

  case 0300:
    msg = "Update Config File"
  case 0301:
    msg = "Remove Cache File"

  }

  return
}

// Show : Show menu on screen
func (m *Menu) Show() (selection int) {

  if len(m.Entry) == 0 {
    return
  }

  fmt.Println()
  fmt.Println(m.Headline)

  for i := 0; i < len(m.Headline); i++ {
    fmt.Print("-")
  }

  fmt.Println()

  for {

    var input string
    var keys []int

    for _, entry := range m.Entry {

      keys = append(keys, entry.Key)

    }

    sort.Ints(keys)

    if keys[0] == 0 {
      keys = keys[1:]
      keys = append(keys, 0)
    }

    for _, key := range keys {

      var entry = m.Entry[key]

      switch len(fmt.Sprintf("%d", entry.Key)) {

      case 1:
        fmt.Print(fmt.Sprintf(" %d. ", entry.Key))
      case 2:
        fmt.Print(fmt.Sprintf("%d. ", entry.Key))

      }

      fmt.Println(entry.Value)

    }

    fmt.Print(fmt.Sprintf("%s: ", m.Select))
    fmt.Scanln(&input)

    selection, err := strconv.Atoi(input)
    if err == nil {

      for _, entry := range m.Entry {

        if selection == entry.Key {
          return selection
        }

      }

    }

    err = errors.New("Invalid Input")
    ShowErr(err)
    fmt.Println()

  }

  return
}

// ShowInfo : Show info on screen
func showInfo(key, msg string) {

  switch len(key) {

  case 1:
    msg = fmt.Sprintf("[%s    ] %s", key, msg)
  case 2:
    msg = fmt.Sprintf("[%s   ] %s", key, msg)
  case 3:
    msg = fmt.Sprintf("[%s  ] %s", key, msg)
  case 4:
    msg = fmt.Sprintf("[%s ] %s", key, msg)
  case 5:
    msg = fmt.Sprintf("[%s] %s", key, msg)

  }

  log.Println(msg)
}

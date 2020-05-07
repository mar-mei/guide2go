package main

import (
  "fmt"
)

func (e *Entry) headline() {

  fmt.Println()
  fmt.Println(e.Value)

  for i := 0; i < len(e.Value); i++ {
    fmt.Print("-")
  }
  fmt.Println()

  return
}

func (e *Entry) account() {

  var username, password string

  e.headline()

  fmt.Print(fmt.Sprintf("%s: ", getMsg(0100)))
  fmt.Scanln(&username)

  fmt.Print(fmt.Sprintf("%s: ", getMsg(0101)))
  fmt.Scanln(&password)

  Config.Account.Username = username
  Config.Account.Password = SHA1(password)

  Config.Save()

  return
}

func (e *Entry) addLineup(sd *SD) (err error) {

  var index, selection int
  var postalcode string
  var menu Menu
  var entry Entry

  menu.Entry = make(map[int]Entry)
  menu.Select = getMsg(0201)
  menu.Headline = e.Value

  err = sd.Countries()
  if err != nil {
    return
  }

  // Cancel
  entry.Key = index
  entry.Value = getMsg(0200)
  menu.Entry[index] = entry

  for _, lineup := range sd.Resp.Countries.NorthAmerica {

    index++
    entry.Key = index
    entry.Value = fmt.Sprintf("%s [%s]", lineup.FullName, lineup.PostalCodeExample)
    entry.Country = lineup.FullName
    entry.Postalcode = lineup.PostalCode
    entry.ShortName = lineup.ShortName
    menu.Entry[index] = entry

  }

  for _, lineup := range sd.Resp.Countries.Europe {

    index++
    entry.Key = index
    entry.Value = fmt.Sprintf("%s [%s]", lineup.FullName, lineup.PostalCodeExample)
    entry.Country = lineup.FullName
    entry.Postalcode = lineup.PostalCode
    entry.ShortName = lineup.ShortName
    menu.Entry[index] = entry

  }

  for _, lineup := range sd.Resp.Countries.LatinAmerica {

    index++
    entry.Key = index
    entry.Value = fmt.Sprintf("%s [%s]", lineup.FullName, lineup.PostalCodeExample)
    entry.Country = lineup.FullName
    entry.Postalcode = lineup.PostalCode
    entry.ShortName = lineup.ShortName
    menu.Entry[index] = entry

  }

  for _, lineup := range sd.Resp.Countries.Caribbean {

    index++
    entry.Key = index
    entry.Value = fmt.Sprintf("%s [%s]", lineup.FullName, lineup.PostalCodeExample)
    entry.Country = lineup.FullName
    entry.Postalcode = lineup.PostalCode
    entry.ShortName = lineup.ShortName
    menu.Entry[index] = entry

  }

  selection = menu.Show()

  switch selection {

  case 0:
    return
  default:
    entry = menu.Entry[selection]

  }

  fmt.Println(entry.Value)

  for {

    fmt.Print(fmt.Sprintf("%s: ", getMsg(0202)))
    fmt.Scanln(&postalcode)

    sd.Req.Parameter = fmt.Sprintf("?country=%s&postalcode=%s", entry.ShortName, postalcode)

    err = sd.Headends()

    if err == nil {
      break
    }

  }

  // Select Linup
  index = 0

  menu.Entry = make(map[int]Entry)
  menu.Select = getMsg(0203)

  // Cancel
  entry.Key = index
  entry.Value = getMsg(0200)
  menu.Entry[index] = entry

  for _, slice := range sd.Resp.Headend {

    for _, lineup := range slice.Lineups {

      index++
      entry.Key = index
      entry.Value = fmt.Sprintf("%s [%s]", lineup.Name, lineup.Lineup)
      entry.Lineup = lineup.Lineup

      menu.Entry[index] = entry

    }

  }

  selection = menu.Show()

  switch selection {

  case 0:
    return
  default:
    entry = menu.Entry[selection]

  }

  sd.Req.Parameter = fmt.Sprintf("/%s", entry.Lineup)
  sd.Req.Type = "PUT"

  err = sd.Lineups()

  return
}

func (e *Entry) removeLineup(sd *SD) (err error) {

  var index, selection int
  var menu Menu
  var entry Entry

  menu.Entry = make(map[int]Entry)
  menu.Select = getMsg(0204)
  menu.Headline = e.Value

  // Cancel
  entry.Key = index
  entry.Value = getMsg(0200)
  menu.Entry[index] = entry

  for _, lineup := range sd.Resp.Status.Lineups {

    index++
    entry.Key = index
    entry.Value = fmt.Sprintf("%s [%s]", lineup.Name, lineup.Lineup)
    entry.Lineup = lineup.Lineup

    menu.Entry[index] = entry

  }

  selection = menu.Show()

  switch selection {

  case 0:
    return

  default:
    entry = menu.Entry[selection]

  }

  sd.Req.Parameter = fmt.Sprintf("/%s", entry.Lineup)
  sd.Req.Type = "DELETE"

  err = sd.Lineups()

  return
}

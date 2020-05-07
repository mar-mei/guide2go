package main

// Menu : Menu
type Menu struct {
  Entry    map[int]Entry
  Headline string
  Select   string
}

// Entry : Menu entry
type Entry struct {
  Key   int
  Value string

  // Add Lineup
  Country    string
  Postalcode string
  ShortName  string
  Lineup     string
}

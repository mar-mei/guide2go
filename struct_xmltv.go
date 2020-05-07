package main

import "encoding/xml"

// XMLTV : XMLTV
type XMLTV struct {
  XMLName   xml.Name     `xml:"tv"`
  Info      string       `xml:"source-info-url,attr"`
  Generator string       `xml:"generator-info-name,attr"`
  Source    string       `xml:"source-info-name,attr"`
  Channel   []ChannelXML `xml:"channel"`
  Program   []Program    `xml:"programme"`
}

type programme struct {
  Channel string `xml:"channel,attr"`
  Start   string `xml:"start,attr"`
  Stop    string `xml:"stop,attr"`

  Title    []Title  `xml:"title"`
  SubTitle SubTitle `xml:"sub-title"`

  Desc        []Desc       `xml:"desc"`
  Categorys   []Category   `xml:"category"`
  EpisodeNums []EpisodeNum `xml:"episode-num"`

  //Icon
  Icon  []Icon `xml:"icon"`
  Video Video  `xml:"video"`
  Audio Audio  `xml:"audio"`

  PreviouslyShown *PreviouslyShown `xml:"previously-shown"`
  New             *New             `xml:"new,omitempty"`
  Live            *Live            `xml:"live"`
}

// ChannelXML : Channel
type ChannelXML struct {
  ID          string        `xml:"id,attr"`
  DisplayName []DisplayName `xml:"display-name"`
  Icon        Icon          `xml:"icon"`
}

// DisplayName : Channel name
type DisplayName struct {
  Value string `xml:",chardata"`
}

// Icon : Channel icon
type Icon struct {
  Src    string `xml:"src,attr"`
  Height int    `xml:"height,attr"`
  Width  int    `xml:"width,attr"`
}

// Program : Channel program
type Program struct {
  Channel string `xml:"channel,attr"`
  Start   string `xml:"start,attr"`
  Stop    string `xml:"stop,attr"`

  Title    []Title  `xml:"title"`
  SubTitle SubTitle `xml:"sub-title"`

  Desc        []Desc       `xml:"desc"`
  Categorys   []Category   `xml:"category"`
  EpisodeNums []EpisodeNum `xml:"episode-num"`

  //Icon
  Icon  []Icon `xml:"icon"`
  Video Video  `xml:"video"`
  Audio Audio  `xml:"audio"`

  PreviouslyShown *PreviouslyShown `xml:"previously-shown"`
  New             *New             `xml:"new,omitempty"`
  Live            *Live            `xml:"live"`
}

type Title struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

type SubTitle struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

type Desc struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

type Category struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

type EpisodeNum struct {
  Value  string `xml:",chardata"`
  System string `xml:"system,attr"`
}

type ProgramIcon struct {
  Src    string `xml:"src,attr"`
  Height int64  `xml:"height,attr"`
  Width  int64  `xml:"width,attr"`
}

type Video struct {
  Present string `xml:"present,omitempty"`
  Colour  string `xml:"colour,omitempty"`
  Aspect  string `xml:"aspect,omitempty"`
  Quality string `xml:"quality,omitempty"`
}

type Audio struct {
  Stereo   string `xml:"stereo,omitempty"`
  Surround string `xml:"surround,omitempty"`
  Mono     string `xml:"mono,omitempty"`
}

type PreviouslyShown struct {
  Start string `xml:"start,attr"`
}

type New struct {
  Value string `xml:",chardata"`
}

type Live struct {
  Value string `xml:",chardata"`
}

package main

import "encoding/xml"

// Programme : Programme
type Programme struct {
  XMLName xml.Name `xml:"programme"`
  Channel string   `xml:"channel,attr"`
  Start   string   `xml:"start,attr"`
  Stop    string   `xml:"stop,attr"`

  Title    []Title  `xml:"title"`
  SubTitle SubTitle `xml:"sub-title"`

  Desc []Desc `xml:"desc"`

  // Credits
  Credits Credits `xml:"credits,omitempty"`

  Categorys   []Category   `xml:"category,omitempty"`
  Language    string       `xml:"language,omitempty"`
  EpisodeNums []EpisodeNum `xml:"episode-num,omitempty"`

  //Icon
  Icon  []Icon `xml:"icon"`
  Video Video  `xml:"video"`
  Audio Audio  `xml:"audio"`

  Rating []Rating `xml:"rating,omitempty"`

  PreviouslyShown *PreviouslyShown `xml:"previously-shown,omitempty"`
  New             *New             `xml:"new"`
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

// Title : Title
type Title struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

// SubTitle : Sub Title
type SubTitle struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

// Desc : Description
type Desc struct {
  Value string `xml:",chardata"`
  Lang  string `xml:"lang,attr"`
}

// Credits : Credits
type Credits struct {
  Director  []Director  `xml:"director,omitempty"`
  Actor     []Actor     `xml:"actor,omitempty"`
  Producer  []Producer  `xml:"producer,omitempty"`
  Presenter []Presenter `xml:"presenter,omitempty"`
  Writer    []Writer    `xml:"writer,omitempty"`
}

type Director struct {
  Value string `xml:",chardata"`
}

type Producer struct {
  Value string `xml:",chardata"`
}

type Presenter struct {
  Value string `xml:",chardata"`
}

type Actor struct {
  Value string `xml:",chardata"`
  Role  string `xml:"role,attr,omitempty"`
}

type Writer struct {
  Value string `xml:",chardata"`
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

type Rating struct {
  System string `xml:"system,attr"`
  Value  string `xml:"value"`
  Icon   []Icon `xml:"icon",omitempty`
}

type Video struct {
  Present string `xml:"present,omitempty"`
  Colour  string `xml:"colour,omitempty"`
  Aspect  string `xml:"aspect,omitempty"`
  Quality string `xml:"quality,omitempty"`
}

type Audio struct {
  Stereo string `xml:"stereo,omitempty"`
}

type PreviouslyShown struct {
  Start string `xml:"start,attr,omitempty"`
}

type New struct {
  Value string `xml:",chardata"`
}

type Live struct {
  Value string `xml:",chardata"`
}

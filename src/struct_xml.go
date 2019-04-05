package src

import (
	"encoding/xml"
)

type Xmltv struct {
	XMLName   xml.Name  `xml:"tv"`
	Info      string    `xml:"source-info-url,attr"`
	Generator string    `xml:"generator-info-name,attr"`
	Source    string    `xml:"source-info-name,attr"`
	Channels  []Channel `xml:"channel"`
	Programs  []Program `xml:"programme"`
}

// Channels
type Channel struct {
	Id           string        `xml:"id,attr"`
	DisplayNames []DisplayName `xml:"display-name"`
	Icon         Icon          `xml:"icon"`
}

type DisplayName struct {
	Value string `xml:",chardata"`
}

type Icon struct {
	Src    string `xml:"src,attr"`
	Height int    `xml:"height,attr"`
	Width  int    `xml:"width,attr"`
}

// Programs
type Program struct {
	Channel string `xml:"channel,attr"`
	Start   string `xml:"start,attr"`
	Stop    string `xml:"stop,attr"`

	Titles   []Title  `xml:"title"`
	SubTitle SubTitle `xml:"sub-title"`

	Descs       []Desc       `xml:"desc"`
	Categorys   []Category   `xml:"category"`
	EpisodeNums []EpisodeNum `xml:"episode-num"`

	//Icon              Icon            `xml:"icon"`
	ProgramIcons []ProgramIcon `xml:"icon"`
	Video        Video         `xml:"video"`
	Audio        Audio         `xml:"audio"`

	PreviouslyShown *PreviouslyShown `xml:"previously-shown"`
	New             *New             `xml:"new"`
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

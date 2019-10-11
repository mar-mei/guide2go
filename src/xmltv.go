package src

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CreateXMLTV(configFile string) {

	showHeadline(190)
	config, err := loadJsonFileToMap(configFile)
	if err != nil {
		checkErr(err, true)
	}

	Config = config

	var outputFile = Config["file.output"].(string)
	var cacheFile = Config["file.cache"].(string)

	loadCachFile()
	if err != nil {
		checkErr(err, true)
	}

	logInfo(AppName, fmt.Sprintf("Config file: %s", configFile))
	logInfo(AppName, fmt.Sprintf("Output file: %s", outputFile))
	logInfo(AppName, fmt.Sprintf("Cache file: %s", cacheFile))

	var xmltv = &Xmltv{Generator: AppName, Source: "Schedules Direct", Info: "http://schedulesdirect.org"}

	// Create Channels

	for _, c := range Cache["channels"].(map[string]interface{}) {

		var jsonString = mapToJson(c)
		var thisChannel G2G_Channels
		json.Unmarshal([]byte(jsonString), &thisChannel)

		var channels Channel
		channels.Id = createChannelID(thisChannel.StationID)
		channels.Icon = Icon{Src: thisChannel.LogoURL, Width: thisChannel.LogoWidth, Height: thisChannel.LogoHeight}

		channels.DisplayNames = append(channels.DisplayNames, DisplayName{Value: thisChannel.Callsign})
		channels.DisplayNames = append(channels.DisplayNames, DisplayName{Value: thisChannel.Name})

		xmltv.Channels = append(xmltv.Channels, channels)

		err, tmp_xmltv := getProgrammData(thisChannel.StationID)
		if err == nil {

			for _, program := range tmp_xmltv.Programs {
				xmltv.Programs = append(xmltv.Programs, program)
			}

		} else {
			checkErr(err, false)
		}

	}

	var content, _ = xml.MarshalIndent(xmltv, "  ", "    ")
	var xmlOutput = []byte(xml.Header + string(content))
	writeByteToFile(outputFile, xmlOutput)

	return
}

func getProgrammData(stationID string) (err error, tmp_xmltv Xmltv) {

	var schedules = Cache["schedules"].(map[string]interface{})

	if s, ok := schedules[stationID]; ok {

		var thisProgram G2G_Schedules
		var jsonString = mapToJson(s)
		json.Unmarshal([]byte(jsonString), &thisProgram)

		for _, p := range thisProgram {

			var program Program
			// Channel ID
			program.Channel = createChannelID(stationID)

			// Time: start and stop
			timeLayout := "2006-01-02 15:04:05 +0000 UTC"
			t, err := time.Parse(timeLayout, p.AirDateTime.Format(timeLayout))
			if err != nil {
				return err, tmp_xmltv
			}

			var dateArray = strings.Fields(t.String())
			var offset = " " + dateArray[2]
			var startTime = t.Format("20060102150405") + offset
			var stopTime = t.Add(time.Second*time.Duration(p.Duration)).Format("20060102150405") + offset
			program.Start = startTime
			program.Stop = stopTime

			// Title
			err, tmp := getProgramTitel(p.ProgramID)
			if err == nil {
				program.Titles = tmp.Titles
			}

			// Sub-Title
			err, tmp = getProgramSubTitle(p.ProgramID)
			if err == nil {
				program.SubTitle = tmp.SubTitle
			}

			// Desc
			err, tmp = getProgramDescription(p.ProgramID)
			if err == nil {
				program.Descs = tmp.Descs
			}

			// Category
			err, tmp = getProgramCategory(p.ProgramID)
			if err == nil {
				program.Categorys = tmp.Categorys
			}

			// Episode-num system="xmltv_ns, original-air-date, onscreen"
			err, tmp = getEpisodeNumSystem(p.ProgramID)
			if err == nil {
				program.EpisodeNums = tmp.EpisodeNums
			}

			// Icon (Cover/Poster)
			err, tmp = getAllProgramIcon(p.ProgramID)
			if err == nil {
				program.ProgramIcons = tmp.ProgramIcons
			}

			// Video
			var videoProperties = p.VideoProperties
			for _, v := range videoProperties {

				var quality = false

				switch v {

				case "hdtv":
					quality = true
				case "sdtv":
					quality = true
				case "uhdtv":
					quality = true
				case "3d":
					quality = true

				}

				if quality == true {
					program.Video.Quality = strings.ToUpper(v)
				}

			}

			// Audio
			var audioProperties = p.AudioProperties
			for _, a := range audioProperties {

				var audio string

				switch a {

				case "stereo":
					audio = "stereo"
				case "DD 5.1":
					audio = "surround"
				case "Atmos":
					audio = "surround"
				case "dubbed":
					audio = "mono"
				case "mono":
					audio = "mono"
				default:
					audio = "stereo"

				}

				switch audio {

				case "stereo":
					program.Audio.Stereo = "dolby digital"
				case "surround":
					program.Audio.Surround = a
				case "mono":
					program.Audio.Mono = "mono"

				}

			}

			// Previously shown
			if p.New == true {

				program.New = &New{Value: ""}

			} else {

				err, tmp = getPreviouslyShown(p.ProgramID)
				if err == nil {
					program.PreviouslyShown = tmp.PreviouslyShown
				}

			}

			// Live
			if p.LiveTapeDelay == "Live" {
				program.Live = &Live{Value: ""}
			}

			tmp_xmltv.Programs = append(tmp_xmltv.Programs, program)
		}

	}

	return
}

func getProgramTitel(programmID string) (err error, program Program) {

	err, p := getProgrammStruct(programmID)

	if err == nil {

		for _, i := range p.Titles {
			program.Titles = append(program.Titles, Title{Value: i.Title120, Lang: getLanguageFromDescription(programmID)})
		}

	}

	return
}

func getProgramSubTitle(programmID string) (err error, program Program) {

	err, p := getProgrammStruct(programmID)
	var subTitle string

	if len(p.EpisodeTitle150) > 0 {

		subTitle = p.EpisodeTitle150

	} else {

		if len(p.Descriptions.Description100) > 0 {
			subTitle = p.Descriptions.Description100[0].Description
		}

	}

	program.SubTitle = SubTitle{Value: subTitle, Lang: getLanguageFromDescription(programmID)}

	return
}

func getProgramDescription(programmID string) (err error, program Program) {

	err, p := getProgrammStruct(programmID)
	if err == nil {

		var desc = p.Descriptions

		for _, i := range desc.Description1000 {

			if len(p.EpisodeTitle150) > 0 {
				program.Descs = append(program.Descs, Desc{Value: i.Description + "\n[" + p.EpisodeTitle150 + "]", Lang: i.DescriptionLanguage})
			} else {
				program.Descs = append(program.Descs, Desc{Value: i.Description, Lang: i.DescriptionLanguage})
			}

		}

	}

	return
}

func getProgramCategory(programmID string) (err error, program Program) {

	err, p := getProgrammStruct(programmID)
	if err == nil {

		for _, i := range p.Genres {
			program.Categorys = append(program.Categorys, Category{Value: i, Lang: "en"})
		}

	}

	return
}

func getEpisodeNumSystem(programmID string) (err error, program Program) {

	err, p := getProgrammStruct(programmID)
	var seaseon, episode int
	var value string

	if err == nil {

		for _, m := range p.Metadata {
			seaseon = m.Gracenote.Season
			episode = m.Gracenote.Episode

			if seaseon > 0 && episode > 0 {
				program.EpisodeNums = append(program.EpisodeNums, EpisodeNum{Value: fmt.Sprintf("%d.%d.", seaseon-1, episode-1), System: "xmltv_ns"})
			}

		}

		if seaseon > 0 && episode > 0 {
			program.EpisodeNums = append(program.EpisodeNums, EpisodeNum{Value: fmt.Sprintf("S%d E%d", seaseon, episode), System: "onscreen"})
		}

		if len(program.EpisodeNums) == 0 {

			switch programmID[0:2] {

			case "EP":
				value = programmID[0:10] + "." + programmID[10:]

			case "SH", "MV":
				value = programmID[0:10] + ".0000"

			default:
				value = programmID
			}

			program.EpisodeNums = append(program.EpisodeNums, EpisodeNum{Value: value, System: "dd_progid"})

		}

		if len(p.OriginalAirDate) > 0 {
			program.EpisodeNums = append(program.EpisodeNums, EpisodeNum{Value: fmt.Sprintf("%s", p.OriginalAirDate), System: "original-air-date"})
		}

	}

	return
}

func getAllProgramIcon(programmID string) (err error, program Program) {

	err, m := getMetadateStruct(programmID[0:10])

	var width, height int64
	var maxWidth int64 = 0

	var aspects = []string{"2x3", "4x3", "3x4", "16x9"}

	if v, ok := Config["poster.aspect"].(string); ok {

		switch v {
		case "all":
			break
		default:
			aspects = []string{v}
		}

	}

	if err == nil {

		for _, aspect := range aspects {

			var uri string

			for _, i := range m {

				if i.Aspect == aspect && i.URI[0:4] == "http" {

					width, err = strconv.ParseInt(i.Width, 10, 64)
					height, err = strconv.ParseInt(i.Height, 10, 64)

					if err == nil {

						if width > maxWidth {
							maxWidth = width
							uri = i.URI

						}

					}

				}

			}

			if maxWidth > 0 && len(uri) > 0 {
				program.ProgramIcons = append(program.ProgramIcons, ProgramIcon{Src: uri, Height: height, Width: width})
			}

		}

	}

	return
}

func getLanguageFromDescription(programmID string) (lang string) {

	err, p := getProgrammStruct(programmID)
	if err == nil {

		if len(p.Descriptions.Description100) > 0 {
			lang = p.Descriptions.Description100[0].DescriptionLanguage
		} else if len(p.Descriptions.Description1000) > 0 {
			lang = p.Descriptions.Description1000[0].DescriptionLanguage
		}

	}

	return
}

func getPreviouslyShown(programmID string) (err error, program Program) {

	err, p := getProgrammStruct(programmID)
	if err == nil {

		if len(p.OriginalAirDate) > 0 {
			program.PreviouslyShown = &PreviouslyShown{Start: p.OriginalAirDate}
		}
	}

	return
}

func getMetadateStruct(programmID string) (err error, thisMetadata G2G_Metadata) {

	var metadata = Cache["metadata"].(map[string]interface{})
	if m, ok := metadata[programmID]; ok {

		var jsonString = mapToJson(m)
		json.Unmarshal([]byte(jsonString), &thisMetadata)

	} else {

		err = errors.New(fmt.Sprintf("No metadata information found [%s]", programmID))

	}

	return
}

func getProgrammStruct(programmID string) (err error, thisProgram G2G_Program) {

	var programs = Cache["programs"].(map[string]interface{})
	if p, ok := programs[programmID]; ok {
		var jsonString = mapToJson(p)
		json.Unmarshal([]byte(jsonString), &thisProgram)
	} else {
		err = errors.New(fmt.Sprintf("No program information found [%s]", programmID))
	}

	return
}

func createChannelID(stationID string) (channelID string) {
	channelID = fmt.Sprintf("%s.%s.schedulesdirect.org", AppName, stationID)
	return
}

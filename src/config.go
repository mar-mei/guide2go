package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

const AppName = "guide2go"
const Version = "0.1.2"

var Config = make(map[string]interface{}) // Configuartion Map
var Cache = make(map[string]interface{})  // Cache Map

func Init() {

	showHeadline(110)
	logInfo(AppName, "Version: "+Version)

	return
}

func CreateModifyConfigurationFile(configFile string) {

	var myChannels = make(map[string]interface{})

	//set or change credentials
	var setCredentials = func(configFile string) (err error) {

		config, err := loadJsonFileToMap(configFile)
		config["sd.username"] = screenInput(001)
		config["sd.password"] = stringToSHA1(screenInput(002))

		err = saveMapToJsonFile(configFile, config)

		return
	}

	// Add lineup
	var addLineup = func(configFile string) (err error) {

		showHeadline(103)

		response, err := sdCountries()

		var menu = make(map[int]string)
		var id = make(map[int]string)

		menu[0] = "Cancel"

		var northAmerica = response.NorthAmerica
		var count = len(menu)

		for i := 0; i < len(northAmerica); i++ {
			menu[i+count] = fmt.Sprintf("%s (%s)", northAmerica[i].FullName, northAmerica[i].PostalCodeExample)
			id[i+count] = fmt.Sprintf("%s", northAmerica[i].ShortName)
		}

		var europe = response.Europe
		count = len(menu)

		for i := 0; i < len(europe); i++ {
			menu[i+count] = fmt.Sprintf("%s (%s)", europe[i].FullName, europe[i].PostalCodeExample)
			id[i+count] = fmt.Sprintf("%s", europe[i].ShortName)
		}

		var latinAmerica = response.LatinAmerica
		count = len(menu)

		for i := 0; i < len(latinAmerica); i++ {
			menu[i+count] = fmt.Sprintf("%s (%s)", latinAmerica[i].FullName, latinAmerica[i].PostalCodeExample)
			id[i+count] = fmt.Sprintf("%s", latinAmerica[i].ShortName)
		}

		var caribbean = response.Caribbean
		count = len(menu)

		for i := 0; i < len(caribbean); i++ {
			menu[i+count] = fmt.Sprintf("%s (%s)", caribbean[i].FullName, caribbean[i].PostalCodeExample)
			id[i+count] = fmt.Sprintf("%s", caribbean[i].ShortName)
		}

		createScreenMenu(menu)

	start:
		var s = screenInput(005)
		input, err := strconv.Atoi(s)

		if err != nil {
			err = errors.New("Invalid Input")
			checkErr(err, false)
			goto start
		}

		switch input {

		case 0:
			CreateModifyConfigurationFile(configFile)
			return err

		}

		if _, ok := menu[input]; ok {

			fmt.Println()
			fmt.Println(menu[input])
			var postalcode = screenInput(006)

			var parameter = "?country=" + id[input] + "&postalcode=" + postalcode

			// Select headends
			headends, err := sdHeadends(parameter)
			if err != nil {
				return err
			}
		plz:

			var menu = make(map[int]string)
			var id = make(map[int]string)
			menu[0] = "Cancel"
			var count = 1

			for _, v := range headends {

				for _, l := range v.Lineups {

					menu[count] = fmt.Sprintf("%s (%s)", l.Name, l.Lineup)
					id[count] = fmt.Sprintf("%s", l.Lineup)
					count++

				}

			}

			createScreenMenu(menu)
			var s = screenInput(007)
			input, err := strconv.Atoi(s)

			if err != nil {
				err = errors.New("Invalid Input")
				checkErr(err, false)
				goto plz

			}

			switch input {
			case 0:
				CreateModifyConfigurationFile(configFile)
				return err
			}

			if _, ok := menu[input]; ok {

				err = sdAddLineup(id[input])
				if err != nil {
					return err
				}

				CreateModifyConfigurationFile(configFile)
				return err

			} else {
				goto plz
			}

		} else {

			err = errors.New("Invalid Input")
			checkErr(err, false)
			goto start

		}

	}

	// Remove liniup
	var removeLineup = func(response SD_Status) (err error) {

		showHeadline(104)
		var menu = make(map[int]string)
		var lineupID = make(map[int]string)

		menu[0] = "Cancel"
		for i, r := range response.Lineups {
			menu[i+1] = fmt.Sprintf("%s (%s)", r.Name, r.Lineup)
			lineupID[i+1] = fmt.Sprintf("%s", r.Lineup)
		}

		createScreenMenu(menu)
	start:
		var s = screenInput(004)
		input, err := strconv.Atoi(s)

		if err != nil {
			err = errors.New("Invalid Input")
			checkErr(err, false)
			goto start
		}

		switch input {

		case 0:
			return

		default:
			err = sdRemoveLineup(lineupID[input])
			if err != nil {
				return
			}

		}

		return
	}

	// Add channel
	var addChannel = func(channel map[string]interface{}) {
		var stationID = channel["stationID"].(string)
		myChannels[stationID] = channel
		//fmt.Println(myChannels)
	}

	// Remove channel
	var removeChannel = func(channel map[string]interface{}) {
		var stationID = channel["stationID"].(string)
		delete(myChannels, stationID)
	}

	// New configuration file
	var newConfiguration = func(configFile string) (err error) {

		showHeadline(100)
		var config = make(map[string]interface{})
		var path = getPlatformPath(configFile)
		var file = "cache_" + getFilenameFromPath(configFile)

		fmt.Println(path + file)

		config["file.output"] = removeFilenameExtension(configFile) + ".xml"
		config["file.cache"] = path + file
		config["poster.aspect"] = "all"
		config["schedule.days"] = 7

		err = saveMapToJsonFile(configFile, config)

		if err != nil {
			return
		}

		err = setCredentials(configFile)
		return
	}

	// Manage channels
	var manageChannels = func(response SD_Status) (err error) {

		showHeadline(106)
		for _, r := range response.Lineups {

			var none = false
			var all = false
			var skip = false
			var name = r.Name

			channelList, err := sdChannelList(r.Lineup)

			if err != nil {
				return err
			}

			logInfo("SD", fmt.Sprintf("Lineup: %s", name))

			// Sort channels by name
			var tmpChannelMap = make(map[string]interface{})
			var channelNames = make([]string, 0)

			for _, c := range channelList.Stations {

				var tmp = make(map[string]interface{})
				tmp["stationID"] = c.StationID
				tmp["channelName"] = c.Name
				tmp["broadcastLanguage"] = c.BroadcastLanguage

				tmpChannelMap[c.Name] = tmp
				channelNames = append(channelNames, c.Name)

			}

			sort.Strings(channelNames)

			// ---

			for _, name := range channelNames {

				var stationID = tmpChannelMap[name].(map[string]interface{})["stationID"].(string)
				var channelName = tmpChannelMap[name].(map[string]interface{})["channelName"].(string)
				var broadcastLanguage = tmpChannelMap[name].(map[string]interface{})["broadcastLanguage"].([]string)

				if all == false && none == false {

					fmt.Println()

					if _, ok := myChannels[stationID]; ok {
						fmt.Printf("[+] [%s] Channel: %s %s\n", stationID, channelName, broadcastLanguage)
					} else {
						fmt.Printf("[-] [%s] Channel: %s %s\n", stationID, channelName, broadcastLanguage)
					}

				}

				var channel = make(map[string]interface{})
				channel["name"] = channelName
				channel["stationID"] = stationID

				if none == false && all == false {

					var input = screenInput(020)

					switch strings.ToLower(input) {

					case "y":
						addChannel(channel)
					case "n":
						delete(myChannels, stationID)
					case "all":
						all = true
					case "none":
						none = true
					case "skip":
						skip = true

					}

				}

				if all == true {
					addChannel(channel)
				}

				if none == true {
					removeChannel(channel)
					//break
				}

				if skip == true {
					break
				}

			}

		}

		Config["channels"] = myChannels

		err = saveMapToJsonFile(configFile, Config)
		if err != nil {
			return
		}

		return
	}

	// Edit configuration file
	var editConfiguration = func(response SD_Status, configFile string) (err error) {
		showHeadline(101)
		var menu = make(map[int]string)

		menu[0] = "Exit"
		menu[1] = "Change credentials from Schedules Direct"
		menu[2] = "Add lineup"
		menu[3] = "Remove lineup"
		menu[4] = "Manage channels"

		createScreenMenu(menu)

	start:
		var s = screenInput(003)
		input, err := strconv.Atoi(s)
		if err != nil {
			err = errors.New("Invalid Input")
			checkErr(err, false)
			goto start
		}

		if _, ok := menu[input]; ok {

			switch input {

			case 0:
				return

			case 1:
				showHeadline(102)
				setCredentials(configFile)

			case 2:
				err = addLineup(configFile)

			case 3:
				err = removeLineup(response)
				CreateModifyConfigurationFile(configFile)
				return

			case 4:
				err = manageChannels(response)
				if err == nil {
					CreateModifyConfigurationFile(configFile)
					return
				}

			}

			if err != nil {
				checkErr(err, false)
				goto start
			}

		} else {

			err = errors.New("Invalid Input")
			checkErr(err, false)
			goto start

		}

		return
	}

	configFile = removeFilenameExtension(configFile) + ".json"

	var err = checkFile(configFile)
	if err != nil {

		err = newConfiguration(configFile)
		if err != nil {
			checkErr(err, true)
		}

	}

	Config, err = loadJsonFileToMap(configFile)
	if err != nil {

		checkErr(err, true)

	} else {

		if v, ok := Config["channels"].(map[string]interface{}); ok {
			myChannels = v
		}

	}

login:
	content, err := readByteFromFile(configFile)
	if err != nil {
		checkErr(err, true)
	}

	var sdgo_account SDGO_account
	err = json.Unmarshal(content, &sdgo_account)
	if err != nil {
		checkErr(err, true)
	}

	// Login to Schedules Direct
	showHeadline(105)

	err = sdLogin(sdgo_account.SdUsername, sdgo_account.SdPassword)
	if err != nil {
		checkErr(err, false)
		setCredentials(configFile)
		goto login
	}

	// Get status from Schedules Direct
	response, err := sdStatus()
	if err != nil {
		checkErr(err, true)
	}

	logInfo(AppName, "Channels: "+strconv.Itoa(len(myChannels)))

	err = editConfiguration(response, configFile)
	if err != nil {
		checkErr(err, true)
	}

	return
}

func GetData(configFile string) {

	var addSchedulesToCache = func(response SD_Schedules) (err error) {

		var schedules = make(map[string]interface{})
		if v, ok := Cache["schedules"].(map[string]interface{}); ok {
			schedules = v
		}

		for _, s := range response {

			var stationID = s.StationID

			var oldSchedules = make([]interface{}, 0)
			if v, ok := schedules[stationID].([]interface{}); ok {
				oldSchedules = v
			}

			for _, i := range s.Programs {
				oldSchedules = append(oldSchedules, i)
			}

			schedules[stationID] = oldSchedules
			Cache["schedules"] = schedules
		}

		err = saveCacheFile()

		return
	}

	var addProgramsToCache = func(response SD_Programs) (err error) {

		err = loadCachFile()
		if err != nil {
			return
		}

		for _, p := range response {

			var programs = Cache["programs"].(map[string]interface{})
			programs[p.ProgramID] = p
			Cache["programs"] = programs

		}

		err = saveCacheFile()

		return
	}

	var addMetadataToCache = func(response SD_Metadata) (err error) {

		err = loadCachFile()
		if err != nil {
			return
		}

		for _, p := range response {

			var metadata = Cache["metadata"].(map[string]interface{})
			metadata[p.ProgramID] = p.Data
			Cache["metadata"] = metadata

		}

		err = saveCacheFile()

		return
	}

	var getChannels = func(response SD_Status) (err error) {
		var myChannels = Config["channels"].(map[string]interface{})
		var channels = make(map[string]interface{})

		for _, r := range response.Lineups {

			channelList, err := sdChannelList(r.Lineup)
			if err != nil {
				return err
			}

			for _, c := range channelList.Stations {

				if _, ok := myChannels[c.StationID]; ok {

					var tmp = make(map[string]interface{})
					tmp["name"] = c.Name
					tmp["stationID"] = c.StationID
					tmp["callsign"] = c.Callsign
					tmp["logoURL"] = c.Logo.URL
					tmp["logoHeight"] = c.Logo.Height
					tmp["logoWidth"] = c.Logo.Width

					channels[c.StationID] = tmp
					Cache["channels"] = channels

				}

			}

		}

		err = saveCacheFile()

		return
	}

	var getSchedulesWithStationIDs = func() (err error) {

		Cache["schedules"] = make(map[string]interface{})

		var currentDay = time.Now()
		var date = make([]string, 0)
		var days = Config["schedule.days"].(float64)

		for i := 0; i < int(days); i++ {
			var nextDay = currentDay.Add(time.Hour * time.Duration(24*i))
			date = append(date, nextDay.Format("2006-01-02"))
		}

		var myChannels = make([]string, 0)
		var sdChannels = make([]interface{}, 0)
		var downloadLimit = 5000
		var count = 0

		// Get all stationsIDs from config file
		if v, ok := Config["channels"].(map[string]interface{}); ok {

			for key, _ := range v {
				myChannels = append(myChannels, key)
			}

		}

		logInfo(AppName, fmt.Sprintf("Schedules Day(s): %0.0f", days))
		logInfo(AppName, fmt.Sprintf("Channels: %d", len(myChannels)))

		// Download schedules from stationIDs
		for _, id := range myChannels {

			var thisCannel = make(map[string]interface{})

			thisCannel["stationID"] = id
			thisCannel["date"] = date

			sdChannels = append(sdChannels, thisCannel)

			count++
			if count == downloadLimit {

				response, err := sdGetSchedules(mapToJson(sdChannels))
				if err != nil {
					return err
				}

				addSchedulesToCache(response)
				sdChannels = make([]interface{}, 0)
				count = 0

			}

		}

		if count < downloadLimit && count > 0 {

			response, err := sdGetSchedules(mapToJson(sdChannels))

			if err != nil {
				return err
			}

			err = addSchedulesToCache(response)
			if err != nil {
				checkErr(err, false)
			}

		}

		return
	}

	// Programs, Metadata
	var getDataWitchProgramIDs = func(dataType string) (err error) {

		var sdPrograms = make([]interface{}, 0)

		err = loadCachFile()
		if err != nil {
			return
		}

		var downloadLimit = 499
		var count = 0
		var newData = 0

		var schedules = Cache["schedules"].(map[string]interface{})

		for key, _ := range schedules {

			jsonString := mapToJson(schedules[key])
			var g2g_Programs G2G_Programs
			json.Unmarshal([]byte(jsonString), &g2g_Programs)

			for _, program := range g2g_Programs {

				var programID string

				switch dataType {
				case "metadata":
					programID = program.ProgramID[0:10]
				case "programs":
					programID = program.ProgramID
				}

				if checkIfDataHasBeenCached(dataType, programID) == false {
					sdPrograms = append(sdPrograms, programID)
					count++
					newData++
				}

				if count == downloadLimit {

					switch dataType {

					case "metadata":

						response, err := sdGetMetadata(mapToJson(sdPrograms))
						if err != nil {
							return err
						}

						err = addMetadataToCache(response)
						if err != nil {
							return err
						}

					case "programs":

						response, err := sdGetPrograms(mapToJson(sdPrograms))
						if err != nil {
							return err
						}

						err = addProgramsToCache(response)
						if err != nil {
							return err
						}

					}

					count = 0
					sdPrograms = make([]interface{}, 0)

				}

			}

		}

		if count > 0 {

			switch dataType {

			case "metadata":

				response, err := sdGetMetadata(mapToJson(sdPrograms))
				if err != nil {
					return err
				}

				err = addMetadataToCache(response)
				if err != nil {
					return err
				}

			case "programs":

				response, err := sdGetPrograms(mapToJson(sdPrograms))
				if err != nil {
					return err
				}

				err = addProgramsToCache(response)
				if err != nil {
					return err
				}

			}

		}

		switch dataType {

		case "programs":
			logInfo(AppName, fmt.Sprintf("New program data: %d", newData))
		case "metadata":
			logInfo(AppName, fmt.Sprintf("New metadata: %d", newData))

		}

		return
	}

	showHeadline(150)
	logInfo(AppName, "Config file: "+configFile)

	content, err := readByteFromFile(configFile)
	if err != nil {
		checkErr(err, true)
	}

	var sdgo_account SDGO_account
	err = json.Unmarshal(content, &sdgo_account)
	if err != nil {
		checkErr(err, true)
	}

	config, err := loadJsonFileToMap(configFile)
	if err != nil {
		checkErr(err, true)
	}

	Config = config
	err = loadCachFile()
	if err != nil {
		checkErr(err, true)
	}

	// Login to Schedules Direct
	showHeadline(105)

	err = sdLogin(sdgo_account.SdUsername, sdgo_account.SdPassword)
	if err != nil {
		checkErr(err, true)
	}

	// Get status from Schedules Direct
	response, err := sdStatus()
	if err != nil {
		checkErr(err, true)
	}

	err = getChannels(response)
	if err != nil {
		checkErr(err, true)
	}

	err = getSchedulesWithStationIDs()
	if err != nil {
		checkErr(err, true)
	}

	err = getDataWitchProgramIDs("programs")
	if err != nil {
		checkErr(err, true)
	}

	err = getDataWitchProgramIDs("metadata")
	if err != nil {
		checkErr(err, true)
	}

	return
}

func CleanUpTheCache(configFile string) {

	var oldData int64

	var removeProgramFromCache = func(dataType, programID string) {

		if data, ok := Cache[dataType].(map[string]interface{}); ok {

			if _, ok := data[programID]; ok {

				delete(data, programID)
				oldData++

			}

		}

	}

	showHeadline(111)
	config, err := loadJsonFileToMap(configFile)
	if err != nil {
		checkErr(err, true)
	}

	Config = config
	err = loadCachFile()
	if err != nil {
		checkErr(err, true)
	}

	var usedIDs = make(map[string]bool)

	if schedules, ok := Cache["schedules"].(map[string]interface{}); ok {

		for _, i := range schedules {

			var tmp = i.([]interface{})
			var jsonString = mapToJson(tmp)
			var g2g_schedules G2G_Schedules
			json.Unmarshal([]byte(jsonString), &g2g_schedules)

			for _, p := range g2g_schedules {

				if len(p.ProgramID) > 0 {
					usedIDs[p.ProgramID] = true
				}

			}

		}

	}

	logInfo(AppName, fmt.Sprintf("Current program data: %d", len(usedIDs)))

	if program, ok := Cache["programs"].(map[string]interface{}); ok {

		for programID, _ := range program {
			if _, ok := usedIDs[programID]; ok {
				// Key already available
			} else {

				removeProgramFromCache("programs", programID)
				removeProgramFromCache("md5", programID)
				removeProgramFromCache("metadata", programID[0:10])

			}

		}

	}

	err = saveCacheFile()
	if err != nil {
		checkErr(err, true)
	}

	logInfo(AppName, fmt.Sprintf("Deleted data: %d", oldData))

	return
}

func loadCachFile() (err error) {

	if cacheFile, ok := Config["file.cache"].(string); ok {
		cache, err := loadJsonFileToMap(cacheFile)
		if err != nil {
			Cache = make(map[string]interface{})
		}

		Cache = cache

		var needKeys = []string{"schedules", "md5", "programs", "metadata", "channels"}
		for _, key := range needKeys {

			if _, ok := cache[key]; ok {
				// Key already available
			} else {
				Cache[key] = make(map[string]interface{})
			}

		}

		err = saveCacheFile()

		if err != nil {
			return err
		}

		return err
	}

	os.Exit(0)

	return
}

func saveCacheFile() (err error) {
	if cacheFile, ok := Config["file.cache"].(string); ok {
		err = saveMapToJsonFile(cacheFile, Cache)
	}

	return
}

func checkIfDataHasBeenCached(dataType, programmID string) (status bool) {
	status = false

	if data, ok := Cache[dataType]; ok {

		var tmp = data.(map[string]interface{})

		if _, ok := tmp[programmID]; ok {
			status = true
		}

	}

	return
}

func checkErr(err error, exit bool) {

	var logMsg = fmt.Sprintf("[ERROR] - %s", err)

	switch runtime.GOOS {

	case "windows":
		log.Println(logMsg)

	default:

		fmt.Print("\033[31m")
		log.Println(logMsg)
		fmt.Print("\033[0m")

	}

	if exit == true {
		os.Exit(0)
	}

	return
}

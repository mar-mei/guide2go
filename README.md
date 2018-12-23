# Short instructions - Version 1.0.2

CLI parameter:
guide2go -h

1. Create a config file

```guide2go -configure filename```

Follow the instructions in the terminal

## Structure of the configuration file:

### Path to cache file (can be changed). In this file, the data is stored by SD, so that the data does not have to be downloaded again next time.

```
"file.cache": "./cache_filename.json" 
```

### Path of the xml file (can be changed)
```
"file.output": "filename.xml"
```

### Change the image format (can be changed). Only the pictures in the highest resolution available are used. Posters are only created if they are also available in the specified format

```
"poster.aspect": "all"
```

- all:  All available image formats are written to the XML file. This feature may be supported by Emby soon.
- 2x3:  Poster is written in 2x3 format in the XML file (Plex image format), 
- 4x3:  Poster is written in 4x3 format in the XML file
- 16x9: Poster is written in 16x9 format in the XML file

### Number of days for the program data (can be changed). Maximum 14 days

```
"schedule.days": 7
```

## Do not change access data in the configuration file
### Changes to the lineup and access data via:

```guide2go -configure filename```

2. Create the XML file:

```guide2go -config filename```
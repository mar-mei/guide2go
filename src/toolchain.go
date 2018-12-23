package src

import (
  "bytes"
  "path/filepath"
  "strings"
  "os"
  "crypto/sha1"
  "io"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "compress/gzip"
)

func stringToSHA1(str string)(strSHA1 string) {

  h := sha1.New()
  io.WriteString(h, str)
  strSHA1 = strings.ToLower(fmt.Sprintf("% x", h.Sum(nil)))
  strSHA1 = strings.Replace(strSHA1, " ", "", -1)

  return
}

func getFilenameFromPath(path string)(string) {

  file := filepath.Base(path)
  
  return file
}

func getPlatformPath(path string)(string) {

  var newPath = filepath.Dir(path) + string(os.PathSeparator)

  return newPath
}

func removeFilenameExtension(basename string)(filename string) {

  filename = strings.TrimSuffix(basename, filepath.Ext(basename))
  
  return 
}

func getPlatformFile(filename string)(newFileName string) {

  path, file := filepath.Split(filename)
  var newPath = filepath.Dir(path)
  newFileName = newPath + string(os.PathSeparator) + file
  
  return
}

func checkFile(filename string)(err error) {

  var file = getPlatformFile(filename)
  _, err = os.Stat(file); os.IsNotExist(err)
  
  return
}

func mapToJson(tmpMap interface{})(string) {

  jsonString, err := json.MarshalIndent(tmpMap, "", "  ")
  if err != nil {
    return "{}"
  }

  return string(jsonString)
}

func saveMapToJsonFile(file string, tmpMap interface{})(error) {

  var filename = getPlatformFile(file)
  jsonString, err := json.MarshalIndent(tmpMap, "", "  ")

  if err != nil {
    return err
  }

  err = ioutil.WriteFile(filename, []byte(jsonString), 0644)
  if err != nil {
    return err
  }

  return nil
}

func loadJsonFileToMap(file string)(error, map[string]interface{})  {

  var tmpMap = make(map[string]interface{})
  var filename = getPlatformFile(file)
  content, err := ioutil.ReadFile(filename)
  if err != nil {

    return err, tmpMap
  
  } else {
  
    err = json.Unmarshal([]byte(content), &tmpMap)
    if err != nil {
      return err, tmpMap
    }
  
  }

  return nil, tmpMap
}

func readByteFromFile(file string)(error, []byte) {

  var err error
  var content []byte

  var filename = getPlatformFile(file)
  content, err = ioutil.ReadFile(filename)
  if err != nil {
    return err, []byte("")
  }

  return err, content
}

func writeByteToFile(file string, data []byte)(error) {
  
  var filename = getPlatformFile(file)  
  var err = ioutil.WriteFile(filename, data, 0644)
  
  return err
}

func gUnzipData(data []byte) (err error, res []byte) {
  
  b := bytes.NewBuffer(data)

  var r io.Reader
  r, err = gzip.NewReader(b)
  if err != nil {
    return
  }

  var resB bytes.Buffer
  _, err = resB.ReadFrom(r)
  if err != nil {
    return
  }

  res = resB.Bytes()
  
  return
}

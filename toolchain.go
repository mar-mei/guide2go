package main

import (
  "bytes"
  "compress/gzip"
  "crypto/sha1"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "strings"
)

// SHA1 : SHA1
func SHA1(str string) (strSHA1 string) {

  h := sha1.New()
  io.WriteString(h, str)
  strSHA1 = strings.ToLower(fmt.Sprintf("% x", h.Sum(nil)))
  strSHA1 = strings.Replace(strSHA1, " ", "", -1)

  return
}

// ContainsString : Get string position in slice
func ContainsString(slice []string, e string) int {
  for i, a := range slice {
    if a == e {
      return i
    }
  }
  return -1
}

func gUnzip(data []byte) (res []byte, err error) {

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
func writeByteToFile(filename string, data []byte) error {

  var err = ioutil.WriteFile(filename, data, 0644)

  return err
}

func mapToJson(tmpMap interface{}) (jsonString string) {

  jsonString = "{}"
  jsonByte, err := json.MarshalIndent(tmpMap, "", "  ")
  if err != nil {
    return
  }

  jsonString = string(jsonByte)
  return
}

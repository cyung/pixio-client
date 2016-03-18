package main

import (
  "io/ioutil"
  "log"
  "encoding/json"
)

type Configuration struct {
  Key string `json:"key"`
  Volume float64 `json:"volume"`
  Directory string `json:"directory"`
}

var _key string
var _volume float64
var _directory string
const _base_url string = "http://pixio.space"

func init() {
  file, err := ioutil.ReadFile("./config.json")
  if err != nil {
    log.Fatal(err)
  }

  var config Configuration
  err = json.Unmarshal(file, &config)
  if err != nil {
    log.Fatal(err)
  }

  _key = config.Key
  _volume = config.Volume
  _directory = config.Directory
}

func GetKey() string {
  return _key
}

func GetBaseUrl() string {
  return _base_url
}

func GetVolume() float64 {
  return _volume
}

func GetFolder() string {
  return _directory
}
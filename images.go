package main

import (
  "time"
  "net/http"
  "os"
  "fmt"
  "io"
  "errors"
  "bytes"
  "mime/multipart"
  "encoding/json"
  "github.com/atotto/clipboard"
)

type Image struct {
  Url string `json:"url"`
}

const folder string = "/Users/chris/Desktop"

func UploadImages() {
  for {
    time.Sleep(1 * time.Second)

    // check if there are files
    filename, err := getFile()
    if err != nil {
      fmt.Println(err)
      continue
    }

    req, err := prepareFile(filename)
    if err != nil {
      fmt.Println(err)
      continue
    }

    url, err := sendFile(req)
    if err != nil {
      fmt.Println(err)
      continue
    }
    clipboard.WriteAll(url)

    err = deleteFile(filename)
    if err != nil {
      fmt.Println(err)
      continue
    }
  }
}

func getFile() (string, error) {
  d, err := os.Open(folder)
  if err != nil {
    return "", err
  }
  defer d.Close()

  filenames, err := d.Readdirnames(-1)
  if err != nil {
    return "", err
  }

  var filtered []string

  for _, filename := range filenames {
    if filename != ".DS_Store" {
      filtered = append(filtered, filename)
    }
  }

  if len(filtered) == 0 {
    return "", errors.New("No images to upload")
  }

  return filtered[0], nil
}

func prepareFile(filename string) (*http.Request, error) {
  file, err := os.Open(folder + "/" + filename)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  // defer writer.Close()

  out, err := writer.CreateFormFile("image", filename)
  if err != nil {
    return nil, err
  }

  _, err = io.Copy(out, file)
  if err != nil {
    return nil, err
  }

  err = writer.Close()
  if err != nil {
    return nil, err
  }

  url := GetBaseUrl() + "/img"
  req, err := http.NewRequest("POST", url, body)
  if err != nil {
    panic(err)
  }
  req.Header.Add("Authorization", GetKey())
  req.Header.Add("Content-Type", writer.FormDataContentType())

  return req, nil
}

func sendFile(req *http.Request) (string, error) {
  fmt.Println("sending file")
  client := &http.Client{}

  res, err := client.Do(req)
  if err != nil {
    fmt.Println("Error sending request")
    return "", err
  }
  defer res.Body.Close()

  if res.StatusCode != 201 {
    return "", errors.New("Error uploading file")
  }

  var image Image
  err = json.NewDecoder(res.Body).Decode(&image)
  if err != nil {
    return "", err
  }

  return image.Url, nil
}

func deleteFile(filename string) error {
  err := os.Remove(folder + "/" + filename)
  if err != nil {
    return err
  }

  return nil
}

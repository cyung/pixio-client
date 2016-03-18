package main

import (
  "io/ioutil"
  "fmt"
  "time"
  "github.com/timshannon/go-openal/openal"
)

func PlaySound() {
  device := openal.OpenDevice("")
  defer device.CloseDevice()

  context := device.CreateContext()
  defer context.Destroy()
  context.Activate()

  source := openal.NewSource()
  defer source.Pause()
  source.SetLooping(false)

  buffer := openal.NewBuffer()

  if err := openal.Err(); err != nil {
    fmt.Println(err)
    return
  }

  data, err := ioutil.ReadFile("assets/ping.wav")
  if err != nil {
    fmt.Println(err)
    return
  }

  buffer.SetData(openal.FormatMono16, data, 44100)

  source.SetBuffer(buffer)
  source.SetGain(GetVolume())
  source.Play()

  // don't delete source until done playing
  for source.State() == openal.Playing {
    time.Sleep(time.Millisecond * 100)
  }
  source.Delete()
}
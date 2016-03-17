package main

import (
  "golang.org/x/mobile/exp/audio"
  "golang.org/x/mobile/asset"
  "log"
)

var _player *audio.Player

func init() {
  rc, err := asset.Open("ping.wav")
  if err != nil {
    log.Fatal(err)
  }

  _player, err = audio.NewPlayer(rc, 0, 0)
  if err != nil {
    log.Fatal(err)
  }

  _player.SetVolume(GetVolume())
}

func PlaySound() {
  _player.Play()
}
##Client for http://pixio.space

###Configuring the client
Modify `config.json`
  * `key` key used to authorize the client
  * `volume` value between 0.0 and 1.0
  * `directory` folder where images will be stored

###Dependencies
* openal-soft (OSX) `brew install openal-soft`
* libopenal (Linux) `sudo apt-get install libopenal`
* OpenAL and OpenAL SDK (Windows) https://www.openal.org/downloads/

###Compiling on OSX/Linux
```
go get
go build
```

###Compiling on Windows
```
go get
go build -ldflags "-H windowsgui"
```
##Client for http://pixio.space

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
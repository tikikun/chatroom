### A quick and dirty chat room on the spot

----------

How to install 

```
$ go get https://github.com/gin-gonic/gin
$ go get https://gopkg.in/olahol/melody.v1
$ go run chat.go 5000 (remeber to choose your port, here is 5000)
```
Nothing much just establish a websocket and has some making room function, work best with heroku free dyno:
* Auto delete the cache fire if no one using the app
* Quick to use

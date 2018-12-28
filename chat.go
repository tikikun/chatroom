package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.Redirect(c.Writer, c.Request, "/room/global", http.StatusFound)
	})

	r.GET("/room/", func(c *gin.Context) {
		http.Redirect(c.Writer, c.Request, "/room/global", http.StatusFound)
	})

	r.GET("/room/:name", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "room.html")
	})

	r.GET("/room/:name/log", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, strings.Split(c.Request.URL.Path, "/")[2]+".log")
	})

	r.GET("/room/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			if q.Request.URL.Path == s.Request.URL.Path {
				logChat(strings.Split(q.Request.URL.Path, "/")[2]+".log", []byte(string(msg)+"\n"))
			}
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	m.HandleDisconnect(func(s *melody.Session) {
		m.BroadcastFilter([]byte("Oops someone has quit"), func(q *melody.Session) bool {
			if q.Request.URL.Path == s.Request.URL.Path {
				logChat(strings.Split(q.Request.URL.Path, "/")[2]+".log", []byte("Oops someone has quit"))
			}
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	r.Run(":" + os.Args[1])
}

func logChat(filename string, bytelog []byte) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(bytelog); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

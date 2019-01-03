package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/olahol/melody"
)

func main() {
	m := http.NewServeMux()
	me := melody.New()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val := strings.Split(r.URL.Path, "/")
		name := val[len(val)-1]
		switch name {
		case "chat.css":
			http.ServeFile(w, r, "assets/css/chat.css")
		case "chat.js":
			http.ServeFile(w, r, "assets/js/chat.js")
		case "room":
			http.Redirect(w, r, "/room/global", http.StatusFound)
		default:
			if name == "" {
				http.Redirect(w, r, "/room/global", http.StatusFound)
			} else {
				fmt.Fprintf(w, "Status not found code %v", http.StatusNotFound)
			}
		}
	})

	m.HandleFunc("/room/", func(w http.ResponseWriter, r *http.Request) {
		val := strings.Split(r.URL.Path, "/")
		name := val[len(val)-1]
		switch name {
		case "":
			fmt.Println("room")
			http.Redirect(w, r, "/room/global", http.StatusFound)
		case "ws":
			me.HandleRequest(w, r)
		case "log":
			http.ServeFile(w, r, val[len(val)-2]+".log")
		default:
			fmt.Println(name)
			http.ServeFile(w, r, "room.html")
		}
	})

	me.HandleMessage(func(s *melody.Session, msg []byte) {
		me.BroadcastFilter(msg, func(q *melody.Session) bool {
			if q.Request.URL.Path == s.Request.URL.Path {
				logChat(strings.Split(q.Request.URL.Path, "/")[2]+".log", []byte(string(msg)+"\n"))
			}
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	me.HandleDisconnect(func(s *melody.Session) {
		me.BroadcastFilter([]byte("Oops someone has quit"), func(q *melody.Session) bool {
			if q.Request.URL.Path == s.Request.URL.Path {
				logChat(strings.Split(q.Request.URL.Path, "/")[2]+".log", []byte("Oops someone has quit"))
			}
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	http.ListenAndServe(":"+os.Args[1], m)
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

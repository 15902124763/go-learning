package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/conn", websocket.Handler(Conn))
	http.ListenAndServe("：9999", nil)
}

func Conn(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err)
			continue
		}

		if err = websocket.Message.Send(ws, strings.ToUpper(reply)); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	w.Header()
}

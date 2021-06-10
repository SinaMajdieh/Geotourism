package commenting

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/database"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var l_upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var conns []*websocket.Conn
func Comment_handler(w http.ResponseWriter , r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	conns = append(conns , conn)
	if err != nil {
		fmt.Println(err)
	}
	for {
		// Read message from browser
		data := database.MSG{}
		err := conn.ReadJSON(&data)
		if err != nil {
			return
		}

		//Print the message to the console


		db.Insert(&data)
		// Write message back to browser

		for _ , v := range conns{
			if err = v.WriteJSON(&data); err != nil {
				fmt.Println(err)
			}
		}

	}
}
func Comment_load_handler(w http.ResponseWriter , r *http.Request){
	l_upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := l_upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		fmt.Println(err)
	}
	for {
		// Read message from browser

		_ , data , err := conn.ReadMessage()
		if err != nil {
			return
		}

		//Print the message to the console

		url := string(data)
		//msg:= MSG{}
		fmt.Println(url)
		sts := db.Load_page_comments(url)
		for _ , v := range sts{
			fmt.Println(v)
		}
		conn.WriteJSON(sts)


	}
}

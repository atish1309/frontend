package main

import (
	"fmt"
	"net/http"

	"./websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		// fmt.Fprintf("websocket not upgraded")
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	//read is a struct client method
	client.Read()
}
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}
func main() {
	fmt.Println("app running")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}

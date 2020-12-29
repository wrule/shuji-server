package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("新的客户端连接:", s.ID())
		return nil
	})

	server.OnEvent("/", "num", func(s socketio.Conn, num int) {
		fmt.Println("num:", num)
		// s.Emit("reply", "have "+msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("连接关闭:", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/", server)
	// http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:9079...")
	log.Fatal(http.ListenAndServe(":9079", nil))
}

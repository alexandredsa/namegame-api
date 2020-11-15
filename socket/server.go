package socket

import (
	"fmt"
	"log"
	"net/http"

	"api.namegame.com/socket/data"

	"api.namegame.com/socket/events/consume"
	socketio "github.com/googollee/go-socket.io"
)

type Server struct {
	Sessions  map[string]data.GameContext
	Listeners []consume.EventConsumer
}

func (s *Server) Start() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		panic(err)
	}

	s.bindEvents(server)

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (s *Server) bindEvents(server *socketio.Server) {
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})

	for _, listener := range s.Listeners {
		name, cb := listener.Bind(data.GameContext{})
		server.OnEvent("/", name, cb)
	}

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
}

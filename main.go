package main

import (
	"api.namegame.com/socket"
	"api.namegame.com/socket/data"
	"api.namegame.com/socket/events/consume"
)

func main() {
	initSocketServer()
}

func initSocketServer() {
	sessions := make(map[string]data.GameContext, 0)
	listeners := make([]consume.EventConsumer, 0)
	listeners = append(listeners, consume.RoomCreate{})
	listeners = append(listeners, consume.RoomJoin{})
	listeners = append(listeners, consume.PlayerStateUpdate{})
	listeners = append(listeners, consume.HunchCreate{})
	server := socket.Server{Listeners: listeners, Sessions: sessions}
	defer server.Start()
}

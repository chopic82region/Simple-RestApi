package main

import (
	"fmt"
	"rest_api/http"
	"rest_api/music"
)

func main() {

	playlist := music.NewPlaylist()
	httpHandler := http.NewHttpHandler(playlist)
	server := http.NewServer(httpHandler)

	if err := server.StartServer(); err != nil {
		fmt.Println("Error of start Server, err")
	}

}

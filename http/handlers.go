package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api/music"
)

type HttpHandler struct {
	musicPlaylist music.Playlist
}

func NewHttpHandler(p music.Playlist) *HttpHandler {
	return &HttpHandler{
		musicPlaylist: p,
	}
}

// Pattern: /music
// Method: /POST
// Response status: 201, 400, 500
func HandleAddMusic(w http.ResponseWriter, r *http.Request) {

	var song music.Music

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		er := music.NewErrMessage(err)
		http.Error(w, er.ErrToString(), http.StatusBadRequest)
		return
	}

	if err := song.IsValidate(); err != nil {
		er := music.NewErrMessage(err)
		http.Error(w, er.ErrToString(), http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(song, "", "	")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("Error of write http response")
	}

}

// Pattern: /music
// Method: /GET
// Response status: 200, 400, 500
func HandleShowPlaylist(w http.ResponseWriter, r *http.Request) {

}

// Pattern: /music/{title}
// Method: /PATCH
// Response status: 200, 404, 500
func HandleDownload(w http.ResponseWriter, r *http.Request) {

}

// Pattern: /music?downloaded=true
// Method: /GET
// Response status: 200, 400, 500
func HandleDownloadedMusic(w http.ResponseWriter, r *http.Request) {

}

// Pattern: /music/{title}
// Method: /GET
// Response status: 200, 404, 500
func HandlePlayMusic(w http.ResponseWriter, r *http.Request) {

}

// Pattern: /music/{title}
// Method: /DELETE
// Response status: 200, 404, 500
func HandleDeleteMusic(w http.ResponseWriter, r *http.Request) {

}

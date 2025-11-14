package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rest_api/music"

	"github.com/gorilla/mux"
)

type HttpHandler struct {
	musicPlaylist *music.Playlist
}

func NewHttpHandler(playlist *music.Playlist) *HttpHandler {
	return &HttpHandler{
		musicPlaylist: playlist,
	}
}

// Pattern: /music
// Method: /POST
// Response status: 201, 400, 500
func (h *HttpHandler) HandleAddMusic(w http.ResponseWriter, r *http.Request) {

	var song MusicDTO

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		errDto := NewErrMessage(err)
		http.Error(w, errDto.ErrToString(), http.StatusBadRequest)
		return
	}

	if err := song.IsValidate(); err != nil {

		errDto := NewErrMessage(err)

		if errors.Is(err, music.ErrMusicNotFound) {
			http.Error(w, errDto.ErrToString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDto.ErrToString(), http.StatusInternalServerError)
		}

		return
	}

	music := music.NewMusic(song.Name, song.Author)

	h.musicPlaylist.AddMusic(music)

	b, err := json.MarshalIndent(song, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write(b); err != nil {
		fmt.Println("Error of write http response")
		return
	}

}

// Pattern: /music
// Method: /GET
// Response status: 200, 400, 500
func (h *HttpHandler) HandleShowPlaylist(w http.ResponseWriter, r *http.Request) {

	b, err := json.MarshalIndent(h.musicPlaylist.ShowPlaylist(), "", "	")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("Error of write http response")
		return
	}

	w.WriteHeader(http.StatusOK)

}

// Pattern: /music/{title}
// Method: /PATCH
// Response status: 200, 404, 500
func (h *HttpHandler) HandleDownload(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		er := NewErrMessage(err)
		http.Error(w, er.ErrToString(), http.StatusBadRequest)
		return
	}

	song, err := h.musicPlaylist.DownloadMusic(name)

	if err != nil {

		errDto := NewErrMessage(err)

		if errors.Is(err, music.ErrMusicNotFound) {
			http.Error(w, errDto.ErrToString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDto.ErrToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(song, " ", "	")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("error of http response", err)
	}

}

// Pattern: /music?downloaded=true
// Method: /GET
// Response status: 200, 400, 500
func (h *HttpHandler) HandleDownloadedMusic(w http.ResponseWriter, r *http.Request) {

	playlist := h.musicPlaylist.ShowDownloadedMusic()

	b, err := json.MarshalIndent(playlist, " ", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("error of http response", err)
	}
}

// Pattern: /music/{title}
// Method: /GET
// Response status: 200, 404, 500
func (h *HttpHandler) HandlePlayMusic(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		er := NewErrMessage(err)
		http.Error(w, er.ErrToString(), http.StatusBadRequest)
		return
	}

	song, err := h.musicPlaylist.PlayMusic(name)

	if err != nil {

		errDto := NewErrMessage(err)

		if errors.Is(err, music.ErrMusicNotFound) {
			http.Error(w, errDto.ErrToString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDto.ErrToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(song, " ", "	")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("error of http response", err)
	}

}

// Pattern: /music/{title}
// Method: /DELETE
// Response status: 200, 404, 500
func (h *HttpHandler) HandleDeleteMusic(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		er := NewErrMessage(err)
		http.Error(w, er.ErrToString(), http.StatusBadRequest)
		return
	}

	err := h.musicPlaylist.DeleteMusic(name)

	if err != nil {

		errDto := NewErrMessage(err)

		if errors.Is(err, music.ErrMusicNotFound) {
			http.Error(w, errDto.ErrToString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDto.ErrToString(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusOK)

}

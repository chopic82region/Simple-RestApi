package http

import (
	"net/http"
	"rest_api/music"
)

type HttpHandler struct {
	musicPlaylist music.Pla
}

// Pattern: /music
// Method: /POST
// Response status: 201, 400, 500
func HandleAddMusic(w http.ResponseWriter, r *http.Request) {

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

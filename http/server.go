package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	httpHandler *HttpHandler
}

func NewServer(httpHandler *HttpHandler) *Server {
	return &Server{
		httpHandler: httpHandler,
	}
}

func (s *Server) StartServer() error {

	router := mux.NewRouter()

	router.Path("/music").Methods("GET").Queries("downloaded", "true").HandlerFunc(s.httpHandler.HandleDownloadedMusic)
	router.Path("/music").Methods("POST").HandlerFunc(s.httpHandler.HandleAddMusic)
	router.Path("/music").Methods("GET").HandlerFunc(s.httpHandler.HandleShowPlaylist)
	router.Path("/music/{name}").Methods("PATCH").HandlerFunc(s.httpHandler.HandleDownload)
	router.Path("/music/{name}").Methods("GET").HandlerFunc(s.httpHandler.HandlePlayMusic)
	router.Path("/music/{name}").Methods("DELETE").HandlerFunc(s.httpHandler.HandleDeleteMusic)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}

package gateway

import (
	"net/http"
)

type httpServer struct {
	server *http.ServeMux
}

func NewHttpServer(server *http.ServeMux) *httpServer {
	return &httpServer{
		server: server,
	}
}

func (h *httpServer) registerRoutes() {
	h.server.HandleFunc("/user", userHandler)
	h.server.HandleFunc("/post", postHandler)
	h.server.HandleFunc("/comment", commentHandler)
	h.server.HandleFunc("/like", likeHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		return
	case http.MethodPut:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		return
	case http.MethodPut:
		return
	case http.MethodDelete:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		return
	case http.MethodPut:
		return
	case http.MethodDelete:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func likeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		return
	case http.MethodPut:
		return
	case http.MethodDelete:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

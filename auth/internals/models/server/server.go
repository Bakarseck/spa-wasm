package server

import (
	"auth/internals/handlers/login"
	"auth/internals/handlers/register"
	"auth/internals/models"
	"auth/internals/utils"
	"fmt"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

var (
	ENDPOINTS = []models.EndPoint{
		{Path: "/login", Handler: login.HandleLogin, Method: http.MethodPost},
		{Path: "/register", Handler: register.HandleRegister, Method: http.MethodPost},
	}
)

func NewServer() Server {
	return Server{
		Router: http.NewServeMux(),
	}
}

func (s *Server) ConfigureRoutes() {
	for _, endpoint := range ENDPOINTS {
		s.Router.HandleFunc(endpoint.Path, s.handleRequest(endpoint.Path, endpoint.Handler, endpoint.Method))
	}
}

func (s *Server) handleRequest(path string, handler http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if path == r.URL.Path {
			if r.Method == method {
				handler(w, r)
				return
			}

			data := map[string]interface{}{"error": "Method Not allowed"}
			utils.RespondWithJSON(w, data, http.StatusMethodNotAllowed)
			return
		}

		data := map[string]interface{}{"error": "Resources not found"}
		utils.RespondWithJSON(w, data, http.StatusNotFound)
	}
}

func (s *Server) StartServer(port string) error {
	s.ConfigureRoutes()
	database, err := utils.OpenDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}
	app := models.App{}
	app.Database = database
	app.SessionHandler = models.CreateSessionManager()
	fmt.Printf("http://localhost:%v", port)
	return http.ListenAndServe(":"+port, s.Router)
}

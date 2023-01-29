package controller

import "github.com/Joel-K-Muraguri/Go-Crud/api/middleware"


func (s *Server) initializeRoutes(){

	s.Router.HandleFunc("/api/games/home", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/api/games/", middleware.SetMiddlewareJSON(s.CreateGame)).Methods("POST")
	s.Router.HandleFunc("/api/games/all", middleware.SetMiddlewareJSON(s.GetGames)).Methods("GET")
	s.Router.HandleFunc("/api/games/{id}", middleware.SetMiddlewareJSON(s.GetGame)).Methods("GET")
	s.Router.HandleFunc("/api/games/delete/{id}", middleware.SetMiddlewareJSON(s.DeleteGame)).Methods("DELETE")
	s.Router.HandleFunc("/api/games/update/{id}", middleware.SetMiddlewareJSON(s.UpdateGame)).Methods("PUT")

}


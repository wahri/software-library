package controllers

import "software_library/backend/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
}

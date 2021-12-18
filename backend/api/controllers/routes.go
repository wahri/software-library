package controllers

import "software_library/backend/api/middlewares"

func (s *Server) initializeRoutes() {
	// // Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

	// Software routes
	s.Router.HandleFunc("/softwares", middlewares.SetMiddlewareJSON(s.CreateSoftware)).Methods("POST")
	s.Router.HandleFunc("/softwares", middlewares.SetMiddlewareJSON(s.GetSoftwares)).Methods("GET")
	s.Router.HandleFunc("/softwares/{id}", middlewares.SetMiddlewareJSON(s.GetSoftware)).Methods("GET")
	s.Router.HandleFunc("/softwares/{id}", middlewares.SetMiddlewareJSON(s.UpdateSoftware)).Methods("PUT")
	s.Router.HandleFunc("/softwares/{id}", middlewares.SetMiddlewareJSON(s.DeleteSoftware)).Methods("DELETE")

	// Software routes
	s.Router.HandleFunc("/video", middlewares.SetMiddlewareJSON(s.CreateVideoTutorial)).Methods("POST")
	s.Router.HandleFunc("/video", middlewares.SetMiddlewareJSON(s.GetVideoTutorials)).Methods("GET")
	s.Router.HandleFunc("/video/{id}", middlewares.SetMiddlewareJSON(s.GetVideoTutorial)).Methods("GET")
	s.Router.HandleFunc("/video/{id}", middlewares.SetMiddlewareJSON(s.UpdateVideoTutorial)).Methods("PUT")
	s.Router.HandleFunc("/video/{id}", middlewares.SetMiddlewareJSON(s.DeleteVideoTutorial)).Methods("DELETE")
}

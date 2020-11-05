package controllers

import (
	"github.com/loeken/rest-blog/api/middlewares"
	_ "github.com/loeken/rest-blog/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) initializeRoutes() {


	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/v1/auth", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")



	//Users routes
	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST","OPTIONS")
	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/user/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/v1/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/user/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/api/v1/post", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/api/v1/post", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/api/v1/post/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/api/v1/post/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/post/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

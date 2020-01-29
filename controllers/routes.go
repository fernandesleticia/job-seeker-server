package controllers

import "job-seeker-server/middlewares"

func (s *Server) initializeRoutes() {

	//Jobs routes
	s.Router.HandleFunc("/jobs", middlewares.SetMiddlewareJSON(s.GetJobs)).Methods("GET")

}

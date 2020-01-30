package controllers

import (
	"job-seeker-server/api/middlewares"
)

func (s *Server) initializeRoutes() {

	//Jobs routes
	s.Router.HandleFunc("/jobs", middlewares.SetMiddlewareJSON(s.GetJobs)).Methods("GET")

	//uploading file
	s.Router.HandleFunc("/upload", s.uploadFile)

	//ask for upload
	s.Router.HandleFunc("/indexUpload", s.indexUpload)

}

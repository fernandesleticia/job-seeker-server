package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"job-seeker-server/models"
	"job-seeker-server/responses"
	"job-seeker-server/utils/formaterror"
	"net/http"
)

func (server *Server) CreateJob(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	job := models.Job{}
	err = json.Unmarshal(body, &job)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	job.Prepare()
	err = job.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	jobCreated, err := job.SaveJob(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, jobCreated.ID))
	responses.JSON(w, http.StatusCreated, jobCreated)
}

func (server *Server) GetJobs(w http.ResponseWriter, r *http.Request) {

	job := models.Job{}

	jobs, err := job.FindAllJobs(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, jobs)
}

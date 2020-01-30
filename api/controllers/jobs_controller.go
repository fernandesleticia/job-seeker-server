package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"job-seeker-server/api/models"
	"job-seeker-server/api/responses"
	"job-seeker-server/api/utils/formaterror"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func (server *Server) SaveJob(job models.Job) {

	_, err := job.SaveJob(server.DB)

	if err != nil {

		log.Fatal(err)
		return
	}
}

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

func (server *Server) uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-txt", "upload-*.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	/****************Reading file***************/

	s := bufio.NewScanner(file)
	s.Scan()
	for s.Scan() {
		var job models.Job

		line := s.Text()
		lineSplited := strings.Split(line, "|")

		partnerId := lineSplited[0]
		partnerIdInt, err := strconv.Atoi(partnerId)
		if err == nil {
			job.PartnerId = partnerIdInt
		}

		title := lineSplited[1]
		job.Title = title

		categoryId := lineSplited[2]
		categoryIdInt, err := strconv.Atoi(categoryId)
		if err == nil {
			job.CategoryId = categoryIdInt
		}

		// date

		// getting location
		location, err := time.LoadLocation("America/Fortaleza")
		if err != nil {
			fmt.Printf("ERROR : %s", err)
		}

		expiresAt := lineSplited[3]
		expiresAtPieces := strings.Split(expiresAt, "/")
		day, month, year := expiresAtPieces[0], expiresAtPieces[1], expiresAtPieces[2]

		dayInt, err := strconv.Atoi(day)
		monthInt, err := strconv.Atoi(month)
		yearInt, err := strconv.Atoi(year)

		expiresAtDate := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, location)
		job.ExpiresAt = expiresAtDate

		_, err = job.SaveJob(server.DB)
		if err != nil {
			fmt.Printf("ERROR : %s", err)
		}

	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func (server *Server) indexUpload(w http.ResponseWriter, r *http.Request) {

	fp := path.Join("static", "form.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

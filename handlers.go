package main

import (
	"bufio"
	"flag"
	"net/http"
	"strconv"

	//"io/ioutil"
	"log"
	"os"
	"strings"

	// "encoding/json"
	"time"
	// "strings"
	// "strconv"
	"fmt"
)

// Recebe, aplica mudanças e retorna nova data
func ServeJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// read file
	fptr := flag.String("fpath", "jobs.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//
	var job Job
	s := bufio.NewScanner(f)
	s.Scan()
	for s.Scan() {
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

		fmt.Println(job)
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Hi there")
}

package main

import "job-seeker-server/api"

func main() {
	// router := mux.NewRouter()

	// router.HandleFunc("/", ServeJobs)

	// log.Fatal(http.ListenAndServe(":8000", router))
	api.Run()
}

package main

import "time"

type Job struct {
	PartnerId  int       `json:"partnerId"`
	Title      string    `json:"title"`
	CategoryId int       `json:"categoryId"`
	ExpiresAt  time.Time `json:"expiresAt"`
	Status     string    `json:"status"`
}

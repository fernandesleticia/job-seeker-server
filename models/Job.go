package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Job struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	PartnerId  int       `gorm:"not null;unique" json:"partnerId"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	CategoryId int       `gorm:"not null" json:"categoryId"`
	ExpiresAt  time.Time `gorm:"not null" json:"expiresAt"`
	Status     string    `gorm:"size:255;not null" json:"status"`
}

func (job *Job) SaveJob(db *gorm.DB) (*Job, error) {

	var err error
	err = db.Debug().Create(&job).Error
	if err != nil {
		return &Job{}, err
	}
	return job, nil
}

func (job *Job) FindAllJobs(db *gorm.DB) (*[]Job, error) {
	var err error
	jobs := []Job{}
	err = db.Debug().Model(&Job{}).Limit(100).Find(&jobs).Error
	if err != nil {
		return &[]Job{}, err
	}
	return &jobs, err
}

package models

import (
	"errors"
	"time"
)

type Task struct {
	ID        uint       `gorm:"primarykey"`
	Title     string     `json:"title"`
	Priority  string     `gorm:"type:ENUM('0','1','2','3');default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `gorm:"default:false" json:"done"`
	ProjectID uint       `json:"project_id"`
}

type Tasks []Task

func (t *Task) Create() error {
	if t.ProjectID == 0 {
		return errors.New("Project id should be specified")
	}
	return db.Create(&t).Error
}

func (t *Task) Update() error {
	return db.Save(&t).Error
}
func (t *Task) Delete() error {
	return db.Delete(&t).Error
}
func (t *Task) MarkAsDone() error {
	done := true
	if t.Done {
		done = false
	}

	t.Done = done
	return db.Model(&t).Select("Done").Updates(&t).Error
}

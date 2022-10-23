package models

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Task struct {
	ID        string     `gorm:"type:varchar(30);primarykey:true;not null;unique" json:"id"`
	Title     string     `gorm:"type:varchar(150)" json:"title"`
	Priority  int        `gorm:"type:decimal;default:0" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline" time_format:"sql_date"`
	Done      bool       `gorm:"default:false" json:"done"`
	ProjectId string     `json:"project_id"`
}

type Tasks []Task

func (t *Task) Create(projectId, userId string) error {
	var project Project
	if err := project.Get(projectId, userId); err != nil {
		return err
	}

	t.ID = ksuid.New().String()
	t.ProjectId = project.ID
	return db.Create(&t).Error
}

func (t *Task) Get(id, projectId, userId string) error {
	var project Project
	if err := project.Get(projectId, userId); err != nil {
		return err
	}

	return db.First(&t, "id = ?", id).Error
}
func (t *Task) Update(id, projectId, userId string) error {
	task := &Task{}
	*task = *t // copy value directly

	if err := t.Get(id, projectId, userId); err != nil {
		return err
	}
	t.ID = id
	return db.Model(&t).Updates(&task).Error
}
func (t *Task) Delete(id, projectId, userId string) error {
	if err := t.Get(id, projectId, userId); err != nil {
		return err
	}
	return db.Delete(&t, "id = ?", id).Error
}

func (t *Task) MarkAsDone(id, projectId, userId string) error {
	if err := t.Get(id, projectId, userId); err != nil {
		return err
	}

	if t.Done {
		t.Done = false
	} else {
		t.Done = true
	}
	return db.Model(&t).Select("Done").Updates(&t).Error
}

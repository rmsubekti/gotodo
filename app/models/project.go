package models

import (
	"errors"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm/clause"
)

type Project struct {
	ID       string `gorm:"type:varchar(30);primarykey:true;not null;unique" json:"id"`
	Title    string `gorm:"type:varchar(150);not null;unique" json:"title"`
	Archived bool   `gorm:"default:false" json:"archived"`
	Tasks    Tasks  `gorm:"auto_preload;foreignkey:ProjectId" json:"tasks"`
	User     User   `gorm:"foreignkey:UserId" json:"-"`
	UserId   string `json:"user_id"`
}

type Projects []Project

func (p *Project) Create() error {
	if len(p.Title) < 1 {
		return errors.New("project title is required")
	}
	p.ID = ksuid.New().String()
	return db.Create(&p).Error
}

func (p *Project) Update(id string, userID string) error {

	title := p.Title

	if err := p.Get(id, userID); err != nil {
		return err
	}

	return db.Model(&p).Updates(Project{Title: title}).Error
}

func (p *Project) Delete(id string, userId string) error {
	if err := p.Get(id, userId); err != nil {
		return err
	}

	// Delete All task on this project
	db.Delete(&Task{}, "ProjectID = ?", id)
	//delete the project
	return db.Delete(&p, "id = ?", id).Error
}

func (p *Project) Get(id string, userId string) error {
	return db.Preload(clause.Associations).First(&p, "id = ? and user_id = ?", id, userId).Error
}

func (p *Projects) List(userId string) error {
	return db.Preload(clause.Associations).Find(&p, "user_id = ?", userId).Error
}

func (p *Project) Archive(id string, userId string) error {
	if err := p.Get(id, userId); err != nil {
		return err
	}

	if p.Archived {
		p.Archived = false
	} else {
		p.Archived = true
	}
	return db.Model(&p).Select("Archived").Updates(&p).Error
}

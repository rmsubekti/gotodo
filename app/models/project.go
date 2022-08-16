package models

import (
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
	p.ID = ksuid.New().String()
	return db.Create(&p).Error
}

func (p *Project) Update(id string) error {

	title := p.Title

	if err := p.Get(id); err != nil {
		return nil
	}

	return db.Model(&p).Updates(Project{Title: title}).Error
}

func (p *Project) Delete(id string) error {
	// Delete All task on this project
	db.Delete(&Task{}, "ProjectID = ?", id)
	//delete the project
	return db.Delete(&p, "id = ?", id).Error
}

func (p *Project) Get(id string) error {
	return db.Preload(clause.Associations).First(&p, "id = ?", id).Error
}

func (p *Projects) List() error {
	return db.Preload(clause.Associations).Find(&p).Error
}

func (p *Project) Archive(id string) error {
	if err := p.Get(id); err != nil {
		return err
	}

	if p.Archived {
		p.Archived = false
	} else {
		p.Archived = true
	}
	return db.Model(&p).Select("Archived").Updates(&p).Error
}

package models

import "gorm.io/gorm/clause"

type Project struct {
	ID       uint   `gorm:"primarykey"`
	Title    string `json:"title"`
	Archived bool   `gorm:"default:false" json:"archived"`
	Tasks    []Task `gorm:"foreignkey:ProjectID" json:"tasks"`
}

type Projects []Project

func (p *Project) Create() error {
	return db.Create(&p).Error
}

func (p *Project) Update() error {
	return db.Save(&p).Error
}

func (p *Project) Delete() error {
	t := &Task{}
	db.Delete(t, "ProjectID = ?", p.ID)
	return db.Delete(&p).Error
}

func (p *Project) Get() error {
	return db.First(&p).Error
}

func (p *Projects) List() error {
	return db.Preload(clause.Associations).Find(&p).Error
}

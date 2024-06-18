package model

import (
	"time"
)

type Article struct {
	ArticleID      int    `gorm:"primaryKey"`
	Title          string `gorm:"not null;unique"`
	Content        string `gorm:"not null"`
	ImageURL       string
	SymptomSummary string
	DiseaseID      int
	Disease        Disease `gorm:"foreignKey:DiseaseID"`
	LabelID        int
	Label          Label        `gorm:"foreignKey:LabelID"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoCreateTime;autoUpdateTime"`
	Symptoms       []Symptom    `gorm:"foreignKey:ArticleID"`
	Preventions    []Prevention `gorm:"foreignKey:ArticleID"`
	Treatments     []Treatment  `gorm:"foreignKey:ArticleID"`
}

func (a Article) TableName() string {
	return "article"
}

type Disease struct {
	DiseaseID   int     `gorm:"primaryKey"`
	DiseaseName string  `gorm:"unique;not null"`
	Cause       string  `gorm:"not null"`
	Plants      []Plant `gorm:"many2many:plant_disease;foreignKey:DiseaseID;joinForeignKey:DiseaseID;References:PlantID;joinReferences:PlantID"`
}

func (d Disease) TableName() string {
	return "disease"
}

type Plant struct {
	PlantID   int       `gorm:"primaryKey"`
	PlantName string    `gorm:"unique;not null"`
	Diseases  []Disease `gorm:"many2many:plant_disease;foreignKey:PlantID;joinForeignKey:PlantID;References:DiseaseID;joinReferences:DiseaseID"`
}

func (p Plant) TableName() string {
	return "plant"
}

type Symptom struct {
	SymptomID          int `gorm:"primaryKey"`
	ArticleID          int
	SymptomDescription string `gorm:"not null"`
}

func (s Symptom) TableName() string {
	return "symptom"
}

type Prevention struct {
	PreventionID          int `gorm:"primaryKey"`
	ArticleID             int
	PreventionDescription string `gorm:"not null"`
}

func (p Prevention) TableName() string {
	return "prevention"
}

type Treatment struct {
	TreatmentID          int `gorm:"primaryKey"`
	ArticleID            int
	TreatmentDescription string `gorm:"not null"`
	TreatmentType        string `gorm:"not null;type:ENUM('organic', 'chemical')"`
}

func (t Treatment) TableName() string {
	return "treatment"
}

type Label struct {
	LabelID   int    `gorm:"primaryKey"`
	LabelName string `gorm:"not null;unique"`
}

func (l Label) TableName() string {
	return "label"
}

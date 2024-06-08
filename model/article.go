package model

import (
	"time"
)

type Article struct {
	ArticleID   int    `gorm:"primaryKey"`
	Title       string `gorm:"not null;unique"`
	Content     string `gorm:"not null"`
	Cause       string `gorm:"not null"`
	ImageURL    string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoCreateTime;autoUpdateTime"`
	CategoryID  int
	Category    Category
	Symptoms    []Symptom    `gorm:"foreignKey:ArticleID"`
	Preventions []Prevention `gorm:"foreignKey:ArticleID"`
	Treatments  []Treatment  `gorm:"foreignKey:ArticleID"`
}

func (a Article) TableName() string {
	return "article"
}

type Category struct {
	CategoryID   int    `gorm:"primaryKey"`
	CategoryName string `gorm:"unique;not null"`
	Articles     []Article
}

func (c Category) TableName() string {
	return "category"
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

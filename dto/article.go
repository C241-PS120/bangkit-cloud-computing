package dto

import "time"

type ArticleResponse struct {
	ArticleID   int               `json:"article_id"`
	Title       string            `json:"title"`
	ImageURL    string            `json:"image_url"`
	Content     string            `json:"content"`
	Category    string            `json:"category"`
	Cause       string            `json:"cause,omitempty"`
	Symptoms    []string          `json:"symptoms,omitempty"`
	Preventions []string          `json:"preventions,omitempty"`
	Treatments  map[string]string `json:"treatments,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type ArticleRequest struct {
	Title       string            `json:"title"`
	Content     string            `json:"content"`
	ImageURL    string            `json:"image_url"`
	Category    string            `json:"category"`
	Cause       string            `json:"cause"`
	Symptoms    []string          `json:"symptoms"`
	Preventions []string          `json:"preventions"`
	Treatments  map[string]string `json:"treatments"`
}

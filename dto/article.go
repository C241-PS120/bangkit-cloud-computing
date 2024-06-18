package dto

type ArticleResponse struct {
	ArticleID      int               `json:"article_id"`
	Title          string            `json:"title"`
	ImageURL       string            `json:"image_url"`
	Disease        string            `json:"disease"`
	Content        string            `json:"content"`
	Cause          string            `json:"cause,omitempty"`
	SymptomSummary string            `json:"symptom_summary,omitempty"`
	Label          string            `json:"label,omitempty"`
	Symptoms       []string          `json:"symptoms,omitempty"`
	Preventions    []string          `json:"preventions,omitempty"`
	Treatments     map[string]string `json:"treatments,omitempty"`
	Plants         []string          `json:"plants,omitempty"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
}

type ArticleRequest struct {
	Title          string            `json:"title" validate:"required"`
	Content        string            `json:"content" validate:"required"`
	Disease        Disease           `json:"disease" validate:"required"`
	Label          string            `json:"label"`
	SymptomSummary string            `json:"symptom_summary"`
	Symptoms       []string          `json:"symptoms"`
	Preventions    []string          `json:"preventions"`
	Treatments     map[string]string `json:"treatments"`
}

type Disease struct {
	DiseaseName string   `json:"disease_name" validate:"required"`
	Cause       string   `json:"cause"`
	Plants      []string `json:"plants"`
}

package converter

import (
	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/C241-PS120/bangkit-cloud-computing/model"
)

func ArticleToResponse(article *model.Article) *dto.ArticleResponse {
	return &dto.ArticleResponse{
		ArticleID:   article.ArticleID,
		Title:       article.Title,
		Content:     article.Content,
		ImageURL:    article.ImageURL,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		Cause:       article.Cause,
		Category:    article.Category.CategoryName,
		Symptoms:    listSymptomToString(article.Symptoms),
		Preventions: listPreventionToString(article.Preventions),
		Treatments:  ListTreatmentToObjectResponse(article.Treatments),
	}
}

func listSymptomToString(symptoms []model.Symptom) []string {
	var response []string
	for _, symptom := range symptoms {
		response = append(response, symptom.SymptomDescription)
	}
	return response
}

func listPreventionToString(preventions []model.Prevention) []string {
	var response []string
	for _, prevention := range preventions {
		response = append(response, prevention.PreventionDescription)
	}
	return response
}

func ListTreatmentToObjectResponse(treatments []model.Treatment) map[string]string {
	response := make(map[string]string)
	for _, treatment := range treatments {
		response[treatment.TreatmentType] = treatment.TreatmentDescription
	}
	return response
}

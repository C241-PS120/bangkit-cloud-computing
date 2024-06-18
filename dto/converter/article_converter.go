package converter

import (
	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/C241-PS120/bangkit-cloud-computing/model"
)

func ArticleToResponse(article *model.Article) *dto.ArticleResponse {
	return &dto.ArticleResponse{
		ArticleID:      article.ArticleID,
		Title:          article.Title,
		Label:          article.Label.LabelName,
		ImageURL:       article.ImageURL,
		Disease:        article.Disease.DiseaseName,
		Content:        article.Content,
		Cause:          article.Disease.Cause,
		SymptomSummary: article.SymptomSummary,
		Symptoms:       listSymptomToString(article.Symptoms),
		Preventions:    listPreventionToString(article.Preventions),
		Treatments:     ListTreatmentToObjectResponse(article.Treatments),
		Plants:         listPlantsToString(article.Disease.Plants),
		CreatedAt:      article.CreatedAt.Format("2 Jan 2006"),
		UpdatedAt:      article.UpdatedAt.Format("2 Jan 2006"),
	}
}

func RequestToArticle(request *dto.ArticleRequest) *model.Article {
	return &model.Article{
		Title:          request.Title,
		Content:        request.Content,
		SymptomSummary: request.SymptomSummary,
		Disease: model.Disease{
			DiseaseName: request.Disease.DiseaseName,
			Cause:       request.Disease.Cause,
			Plants:      listStringToPlant(request.Disease.Plants),
		},
		Label: model.Label{
			LabelName: request.Label,
		},
		Symptoms:    listStringToSymptom(request.Symptoms),
		Preventions: listStringToPrevention(request.Preventions),
		Treatments:  listStringToTreatment(request.Treatments),
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

func listPlantsToString(plants []model.Plant) []string {
	var response []string
	for _, plant := range plants {
		response = append(response, plant.PlantName)
	}
	return response
}

func listStringToSymptom(symptoms []string) []model.Symptom {
	var response []model.Symptom
	for _, symptom := range symptoms {
		response = append(response, model.Symptom{
			SymptomDescription: symptom,
		})
	}
	return response
}

func listStringToPrevention(preventions []string) []model.Prevention {
	var response []model.Prevention
	for _, prevention := range preventions {
		response = append(response, model.Prevention{
			PreventionDescription: prevention,
		})
	}
	return response
}

func listStringToTreatment(treatments map[string]string) []model.Treatment {
	var response []model.Treatment
	for key, value := range treatments {
		response = append(response, model.Treatment{
			TreatmentDescription: value,
			TreatmentType:        key,
		})
	}
	return response
}

func listStringToPlant(plants []string) []model.Plant {
	var response []model.Plant
	for _, plant := range plants {
		response = append(response, model.Plant{
			PlantName: plant,
		})
	}
	return response
}

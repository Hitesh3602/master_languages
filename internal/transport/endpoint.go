package http

import (
	"context"
	"time"

	"github.com/Hitesh3602/master_languages/internal/model"
	"github.com/Hitesh3602/master_languages/internal/service"
	"github.com/go-kit/kit/endpoint"
)

// type createLanguageRequest struct {
//     Language model.Language
// }

// type getLanguagesResponse struct {
// 	Languages []*model.Language `json:"languages"`
// }

type getLanguagesResponse struct {
	Languages []*model.Language `json:"languages,omitempty"`
	Language  *model.Language   `json:"language,omitempty"`
}

//	type getLanguageByIDRequest struct {
//		ID int `json:"id"`
//	}
type getLanguagesRequest struct {
	ID *int `json:"id,omitempty"`
}

type updateLanguageRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	// Languages []*model.Language `json:"languages,omitempty"`
	// Language  *model.Language   `json:"language,omitempty"`
	ShortCode string            `json:"short_code"`
	IsActive  bool              `json:"is_active"`
}

type deleteLanguageRequest struct {
	ID int `json:"id"`
}

type createLanguageRequest struct {
	Name      string `json:"name"`
	ShortCode string `json:"short_code"`
	IsActive  bool   `json:"is_active"`
}

func makeCreateLanguageEndpoint(svc service.LanguageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createLanguageRequest)
		language := &model.Language{
			Name:      req.Name,
			ShortCode: req.ShortCode,
			IsActive:  req.IsActive,
		}
		err := svc.CreateLanguage(language)
		if err != nil {
			return nil, err
		}
		return language, nil
	}
}

//	func makeGetLanguagesEndpoint(svc service.LanguageService) endpoint.Endpoint {
//		return func(ctx context.Context, request interface{}) (interface{}, error) {
//			languages, err := svc.GetLanguages()
//			if err != nil {
//				return nil, err
//			}
//			return getLanguagesResponse{Languages: languages}, nil
//		}
//	}
func makeGetLanguagesEndpoint(svc service.LanguageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getLanguagesRequest)
		if req.ID != nil {
			language, err := svc.GetLanguageByID(*req.ID)
			if err != nil {
				return nil, err
			}
			return getLanguagesResponse{Language: language}, nil
		} else {
			languages, err := svc.GetLanguages()
			if err != nil {
				return nil, err
			}
			return getLanguagesResponse{Languages: languages}, nil
		}
	}
}

// func makeGetLanguageByIDEndpoint(svc service.LanguageService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(getLanguageByIDRequest)
// 		language, err := svc.GetLanguageByID(req.ID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return language, nil
// 	}
// }

func makeUpdateLanguageEndpoint(svc service.LanguageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateLanguageRequest)
		language := &model.Language{
			ID:        req.ID,
			Name:      req.Name,
			ShortCode: req.ShortCode,
			IsActive:  req.IsActive,
			CreatedAt: time.Time{},
			UpdatedAt: time.Now(),
		}
		err := svc.UpdateLanguage(language)
		if err != nil {
			return nil, err
		}
		return language, nil
	}
}

func makeDeleteLanguageEndpoint(svc service.LanguageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteLanguageRequest)
		err := svc.DeleteLanguage(req.ID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

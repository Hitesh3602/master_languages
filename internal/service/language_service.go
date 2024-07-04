package service

import (
	"time"

	"github.com/Hitesh3602/master_languages/internal/db"
	"github.com/Hitesh3602/master_languages/internal/model"
)

type LanguageService interface {
	CreateLanguage(language *model.Language) error
	GetLanguages() ([]*model.Language, error)
	GetLanguageByID(id int) (*model.Language, error)
	UpdateLanguage(language *model.Language) error
	DeleteLanguage(id int) error
}

type languageService struct {
	repo db.LanguageRepository
}

func NewLanguageService(repo db.LanguageRepository) LanguageService {
	return &languageService{repo: repo}
}

func (s *languageService) CreateLanguage(language *model.Language) error {
	language.CreatedAt = time.Now()
	language.UpdatedAt = time.Now()
	return s.repo.Create(language)
}

func (s *languageService) GetLanguages() ([]*model.Language, error) {
	return s.repo.GetAll()
}

func (s *languageService) GetLanguageByID(id int) (*model.Language, error) {
	return s.repo.GetByID(id)
}

func (s *languageService) UpdateLanguage(language *model.Language) error {
	return s.repo.Update(language)
}

func (s *languageService) DeleteLanguage(id int) error {
	return s.repo.Delete(id)
}

package db

import "github.com/Hitesh3602/master_languages/internal/model"

type LanguageRepository interface {
    Create(language *model.Language) error
    GetAll() ([]*model.Language, error)
    GetByID(id int) (*model.Language, error)
    Update(language *model.Language) error
    Delete(id int) error
}

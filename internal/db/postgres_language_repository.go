package db

import (
	"database/sql"

	"github.com/Hitesh3602/master_languages/internal/model"
)

type PostgresLanguageRepository struct {
	DB *sql.DB
}

func NewPostgresLanguageRepository(db *sql.DB) *PostgresLanguageRepository {
	return &PostgresLanguageRepository{DB: db}
}

func (r *PostgresLanguageRepository) Create(language *model.Language) error {
	query := `INSERT INTO master_languages (name, short_code, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.DB.QueryRow(query, language.Name, language.ShortCode, language.IsActive, language.CreatedAt, language.UpdatedAt).Scan(&language.ID)
}

func (r *PostgresLanguageRepository) GetAll() ([]*model.Language, error) {
    query := `SELECT id, name, short_code, is_active, created_at, updated_at FROM master_languages ORDER BY id`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    languages := []*model.Language{}
    for rows.Next() {
        language := new(model.Language)
        err := rows.Scan(&language.ID, &language.Name, &language.ShortCode, &language.IsActive, &language.CreatedAt, &language.UpdatedAt)
        if err != nil {
            return nil, err
        }
        languages = append(languages, language)
    }
    return languages, nil
}

	func (r *PostgresLanguageRepository) GetByID(id int) (*model.Language, error) {
	    query := `SELECT id, name, short_code, is_active, created_at, updated_at FROM master_languages WHERE id = $1`
	    language := new(model.Language)
	    err := r.DB.QueryRow(query, id).Scan(&language.ID, &language.Name, &language.ShortCode, &language.IsActive, &language.CreatedAt, &language.UpdatedAt)
	    if err != nil {
	        return nil, err
	    }
	    return language, nil
	}


func (r *PostgresLanguageRepository) Update(language *model.Language) error {
	query := `UPDATE master_languages SET name = $1, short_code = $2, is_active = $3, updated_at = $4 WHERE id = $5`
	_, err := r.DB.Exec(query, language.Name, language.ShortCode, language.IsActive, language.UpdatedAt, language.ID)
	return err
}

func (r *PostgresLanguageRepository) Delete(id int) error {
	query := `DELETE FROM master_languages WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

package model

import "time"

type Language struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    ShortCode string    `json:"short_code"`
    IsActive  bool      `json:"is_active"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

package models

import "time"

type Banner struct {
	ID        int       `pg:"id" json:"id"`
	FeatureID int       `pg:"feature" json:"feature_id"`
	TagID     int       `pg:"tag" json:"tag_id"`
	Content   string    `pg:"contents" json:"content"`
	CreatedAt time.Time `pg:"created_at" json:"created_at"`
	UpdatedAt time.Time `pg:"updated_at" json:"updated_at"`
	IsActive  bool      `pg:"is_active" json:"is_active"`
}

type BannerOut struct {
	Banners []Banner `json:"banners"`
}

package models

import "time"

type Banner struct {
	ID        int       ``
	FeatureID int       ``
	TagID     int       ``
	Title     string    ``
	Text      string    ``
	URL       string    ``
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	IsActive  bool      ``
}

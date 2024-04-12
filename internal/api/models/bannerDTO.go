package models

type Content struct {
	Title string
	Text  string
	URL   string
}

type BannerDTO struct {
	TagIDs    []int `json:"tag_ids"`
	FeatureID int   `json:"feature_id"`
	IsActive  bool  `json:"is_active"`
	Content   `json:"content"`
}

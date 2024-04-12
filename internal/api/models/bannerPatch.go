package models

type BannerPatch struct {
	ID        int
	TagIDs    *[]int  `json:"tag_ids"`
	FeatureID *int    `json:"feature_id"`
	IsActive  *bool   `json:"is_active"`
	Content   *string `json:"content"`
}

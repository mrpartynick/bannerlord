package storage

import "bannerlord/internal/api/models"

type banners interface {
	GetBanners(feature int, tag int, isActive bool) ([]models.Banner, error)
	GetAdminBanners(feature int, tag int) ([]models.Banner, error)
	GetByFeature(feature int) ([]models.Banner, error)
	GetByTag(tag int) ([]models.Banner, error)
	GetAll() ([]models.Banner, error)
	UpdateBanner(banner models.BannerPatch) error
	DeleteBanner(id int) error
}
